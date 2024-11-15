package product

import (
	entity "github.com/srv-cashpay/auth/entity"
)

func (s *merchantService) Get() ([]entity.MerchantDetail, error) {
	// Fetch comments from the repository layer based on post_id
	comments, err := s.Repo.Get()
	if err != nil {
		return nil, err
	}

	return comments, nil
}
