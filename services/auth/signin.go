package auth

import (
	"errors"
	"strings"
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"

	"gorm.io/gorm"
)

func (u *authService) SigninByPhoneNumber(req dto.SigninRequest) (*dto.SigninResponse, error) {
	req.Email = strings.ToLower(req.Email)

	user, err := u.Repo.SigninByPhoneNumber(req)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, res.ErrorBuilder(&res.ErrorConstant.RecordNotFound, err)
		}

		// Handle other errors
		return nil, res.ErrorResponse(err)
	}

	// if !user.Verified.Verified {
	// 	return nil, res.ErrorBuilder(&res.ErrorConstant.Unverified, err)
	// }

	if !user.Verified.Verified {

		otp := GenerateRandomNumeric(4)
		tokenVerified := util.GenerateRandomString()

		// Kirim ulang OTP ke email user (gunakan pustaka pengiriman email yang sesuai)
		if err := util.Mailtrap(user.Email, otp); err != nil {
			return nil, err
		}

		// Perbarui nilai TokenVerified di repository
		if _, err := u.Repo.UpdateTokenVerified(user.ID, otp, tokenVerified); err != nil {
			return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		// Contoh respons:
		response := &dto.SigninResponse{
			VerifiedResp: &dto.AuthUnverifiedResponse{
				TokenVerified: tokenVerified,
				Otp:           otp,
			},
		}
		// // return dto.SigninResponse{
		// // 	TokenVerified: user.Verified.Token,
		// // 	Otp:           user.Verified.Otp,
		// // }, res.ErrorBuilder(&res.ErrorConstant.Unverified, err)
		return response, nil
		// return nil, res.ErrorBuilder(&res.ErrorConstant.Unverified, err)

	}

	if user.SuspendedUntil != nil && user.SuspendedUntil.After(time.Now()) {
		return nil, res.ErrorBuilder(&res.ErrorConstant.Suspend, nil)
	}

	if err != nil || util.VerifyPassword(user.Password, req.Password) != nil {
		user.FailedSigninAttempts++

		// Check if the maximum allowed failed attempts is reached
		if user.FailedSigninAttempts >= 6 {
			// Suspend the account for 5 minutes
			suspendedUntil := time.Now().Add(5 * time.Minute)
			user.SuspendedUntil = &suspendedUntil

			// Reset failed signin attempts counter
			user.FailedSigninAttempts = 0
		}

		// Update the user in the repository
		if err := u.Repo.UpdateUser(user); err != nil {
			return nil, res.ErrorResponse(err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.VerifyPassword, err)
	}

	user.FailedSigninAttempts = 0

	token, err := u.jwt.GenerateToken(user.ID, user.UserDetail.FullName, user.UserDetail.ProfileID)
	return &dto.SigninResponse{
		ID:            user.ID,
		FullName:      user.UserDetail.FullName,
		ProfileID:     user.UserDetail.ProfileID,
		Email:         user.Email,
		Token:         token,
		TokenVerified: user.Verified.Token,
	}, nil
}
func (u *authService) Signin(req dto.SigninRequest) (*dto.SigninResponse, error) {
	req.Email = strings.ToLower(req.Email)

	// Encrypt the email
	encryptedEmail, err := util.Encrypt(req.Email)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Set the encrypted email in the request
	req.Email = encryptedEmail

	user, err := u.Repo.Signin(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, res.ErrorBuilder(&res.ErrorConstant.RecordNotFound, err)
		}
		return nil, res.ErrorResponse(err)
	}

	if !user.Verified.Verified {
		otp := GenerateRandomNumeric(4)
		tokenVerified := util.GenerateRandomString()

		if err := util.Mailtrap(req.Email, otp); err != nil {
			return nil, err
		}

		if _, err := u.Repo.UpdateTokenVerified(user.ID, otp, tokenVerified); err != nil {
			return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		response := &dto.SigninResponse{
			VerifiedResp: &dto.AuthUnverifiedResponse{
				TokenVerified: tokenVerified,
				Otp:           otp,
			},
		}
		return response, nil
	}

	if user.SuspendedUntil != nil && user.SuspendedUntil.After(time.Now()) {
		return nil, res.ErrorBuilder(&res.ErrorConstant.Suspend, nil)
	}

	if err != nil || util.VerifyPassword(user.Password, req.Password) != nil {
		user.FailedSigninAttempts++

		// Check if the maximum allowed failed attempts is reached
		if user.FailedSigninAttempts >= 6 {
			// Suspend the account for 5 minutes
			suspendedUntil := time.Now().Add(5 * time.Minute)
			user.SuspendedUntil = &suspendedUntil

			// Reset failed signin attempts counter
			user.FailedSigninAttempts = 0
		}

		// Update the user in the repository
		if err := u.Repo.UpdateUser(user); err != nil {
			return nil, res.ErrorResponse(err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.VerifyPassword, err)
	}

	user.FailedSigninAttempts = 0

	token, err := u.jwt.GenerateToken(user.ID, user.UserDetail.FullName, user.UserDetail.ProfileID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Decrypt the email for the response
	decryptedEmail, err := util.Decrypt(user.Email)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &dto.SigninResponse{
		ID:            user.ID,
		FullName:      user.UserDetail.FullName,
		ProfileID:     user.UserDetail.ProfileID,
		Email:         decryptedEmail, // Return decrypted email
		Token:         token,
		TokenVerified: user.Verified.Token,
	}, nil
}
