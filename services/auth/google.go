package auth

import (
	"context"
	"errors"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	util "github.com/srv-cashpay/util/s"
	"google.golang.org/api/idtoken"
)

func (s *authService) SignInWithGoogle(req dto.GoogleSignInRequest) (*dto.AuthResponse, error) {
	// 1. Validasi ID token dari Google
	payload, err := idtoken.Validate(context.Background(), req.IdToken, "")
	if err != nil {
		return nil, err
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return nil, errors.New("email not found in token")
	}
	name, _ := payload.Claims["name"].(string)

	// 2. Enkripsi email untuk query dan simpan
	encryptedEmail, err := util.Encrypt(email)
	if err != nil {
		return nil, err
	}

	// 3. Cek apakah user sudah ada
	user, err := s.Repo.FindByEncryptedEmail(encryptedEmail)
	if err != nil && err.Error() == "user not found" {
		// 4. Jika user belum ada, buat user baru
		if req.Whatsapp == "" {
			// Tidak bisa buat akun tanpa WhatsApp
			return &dto.AuthResponse{
				FullName: name,
				Email:    encryptedEmail,
				Whatsapp: "", // <- tanda frontend untuk arahkan ke input WhatsApp
			}, nil
		}

		secureID, err := generateSecureID()
		if err != nil {
			return nil, err
		}

		encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
		if err != nil {
			return nil, err
		}

		user = &entity.AccessDoor{
			ID:           secureID,
			Email:        encryptedEmail,
			FullName:     name,
			Provider:     "google",
			AccessRoleID: "e9Wl2JyVeBM_",
			Whatsapp:     encryptedWhatsapp,
		}

		if err := s.Repo.Create(user); err != nil {
			return nil, err
		}

		// Reload dengan preload relasi
		user, err = s.Repo.FindByEncryptedEmail(encryptedEmail)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		// Error lainnya
		return nil, err
	} else {
		// 5. Jika user sudah ada tapi WhatsApp masih kosong dan request bawa WA
		if user.Whatsapp == "/bvTmYgHVZjVt85fktdsXA==" && req.Whatsapp != "/bvTmYgHVZjVt85fktdsXA==" {
			if err := s.Repo.UpdateWhatsapp(user.ID, req.Whatsapp); err != nil {
				return nil, err
			}
			user.Whatsapp = req.Whatsapp
		}
	}

	// 6. Generate token
	token, err := s.jwt.GenerateToken(user.ID, user.FullName, user.Merchant.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Merchant.ID)
	if err != nil {
		return nil, err
	}

	// 7. Kembalikan response
	return &dto.AuthResponse{
		ID:            user.ID,
		MerchantID:    user.Merchant.ID,
		FullName:      user.FullName,
		Whatsapp:      user.Whatsapp,
		Email:         user.Email,
		Token:         token,
		RefreshToken:  refreshToken,
		TokenVerified: user.Verified.Token,
	}, nil
}
