package auth

import (
	"errors"

	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"
	"gorm.io/gorm"
)

func (u *authService) RefreshAccessToken(req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	// Validasi refresh token dan dapatkan user ID
	request := dto.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
		UserID: req UserID,
	}

	user, err := u.Repo.RefreshToken(request)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.RefreshTokenResponse{}, res.ErrorBuilder(&res.ErrorConstant.RecordNotFound, err)
		}
		return dto.RefreshTokenResponse{}, res.ErrorResponse(err)
	}

	// Generate token baru
	accessToken, err := u.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Merchant.ID)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	resp := dto.RefreshTokenResponse{
		AccessToken: accessToken,
	}

	return resp, nil
}
