package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"
)

func (s *authService) Signup(req dto.SignupRequest) (dto.SignupResponse, error) {
	user := dto.SignupRequest{
		ID:       util.GenerateRandomString(),
		Whatsapp: req.Whatsapp,
	}

	createdUser, err := s.Repo.Signup(user)
	if err != nil {
		return dto.SignupResponse{}, err
	}

	token, err := s.jwt.GenerateToken(user.ID, user.Whatsapp)
	if err != nil {
		return dto.SignupResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	response := dto.SignupResponse{
		ID:       createdUser.ID,
		Whatsapp: createdUser.Whatsapp,
		Token:    token,
	}
	return response, nil
}
