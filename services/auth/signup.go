package auth

import "github.com/srv-cashpay/auth/entity"

func (s *authService) Signup(user *entity.User) error {
	// Implement hashing password logic here if needed
	// return s.repo.Create(user)
	return s.Repo.Signup(user)
}
