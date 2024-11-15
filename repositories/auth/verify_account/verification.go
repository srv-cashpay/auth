package repositories

import (
	"fmt"

	"github.com/srv-cashpay/auth/entity"
)

func (u *verifyRepository) UpdateUserVerificationStatus(user *entity.UserVerified) error {
	// Replace the following code with your actual logic to update user verification status in the database
	// Example: err := userRepository.UpdateVerificationStatus(user.ID)

	// Simulate updating user verification status (replace with your actual logic)
	user.Verified = true

	// In a real-world scenario, you would use a database query to update the user's verification status
	// Example using GORM:
	err := u.DB.Model(&entity.UserVerified{}).Where("id = ?", user.ID).Update("verified", true).Error
	if err != nil {
		// Handle the error appropriately (e.g., log it, return it, etc.)
		return err
	}
	// Check for errors and return them
	if user.ID == "invalid_user_id" {
		return fmt.Errorf("user not found")
	}

	// Check for errors and return them
	return nil
}
