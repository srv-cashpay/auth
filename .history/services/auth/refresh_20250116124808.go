package auth

import (
	"errors"

	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"
	"gorm.io/gorm"
)

func (u *authService) RefreshAccessToken(req dto.RefreshTokenRequest) (string, error) {
	// Validasi refresh token dan dapatkan user ID
	user, err := u.Repo.RefreshToken(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", res.ErrorBuilder(&res.ErrorConstant.RecordNotFound, err)
		}
		return "", res.ErrorResponse(err)
	}

	// Generate token baru
	accessToken, err := u.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Merchant.ID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
