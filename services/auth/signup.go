package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
)

func (s *authService) Signup(req dto.SignupRequest) (dto.SignupResponse, error) {
	// Generate random ID for the user
	user := dto.SignupRequest{
		ID:       util.GenerateRandomString(),
		Whatsapp: req.Whatsapp,
	}

	// Encrypt the WhatsApp number
	// encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
	// if err != nil {
	// 	return dto.SignupResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	// }

	// // Store encrypted WhatsApp number
	// user.Whatsapp = encryptedWhatsapp

	// Save the user to the database
	createdUser, err := s.Repo.Signup(user)
	if err != nil {
		return dto.SignupResponse{}, err
	}

	// Generate JWT token using encrypted WhatsApp
	// token, err := s.jwt.GenerateToken(user.ID, user.Whatsapp)
	// if err != nil {
	// 	return dto.SignupResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	// }

	// Prepare the signup response with encrypted WhatsApp
	response := dto.SignupResponse{
		ID:            createdUser.ID,
		Whatsapp:      createdUser.Whatsapp,
		TokenVerified: createdUser.TokenVerified,
	}

	return response, nil
}
