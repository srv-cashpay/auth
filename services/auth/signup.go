package auth

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"strings"
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"
	"golang.org/x/crypto/blake2b"
	"gopkg.in/gomail.v2"
)

func (u *authService) Signup(req dto.SignupRequest) (dto.SignupResponse, error) {
	// Validate email
	if !util.IsValidEmail(req.Email) {
		return dto.SignupResponse{}, res.ErrorBuilder(&res.ErrorConstant.RegisterMail, nil)
	}

	req.Whatsapp = FormatWhatsappNumber(req.Whatsapp)

	// Encrypt the email
	encryptedEmail, err := util.Encrypt(req.Email)
	if err != nil {
		return dto.SignupResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Proceed with the signup process
	encryp := util.EncryptPassword(&req)
	if encryp != nil {
		return dto.SignupResponse{}, encryp
	}

	// Generate unique ID and OTP
	secureID, err := generateSecureID()
	if err != nil {
		return dto.SignupResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	user := dto.SignupRequest{
		ID:           secureID,
		Otp:          GenerateRandomNumeric(4),
		Whatsapp:     req.Whatsapp,
		FullName:     req.FullName,
		Email:        encryptedEmail,
		Password:     req.Password,
		Token:        util.GenerateRandomString(),
		AccessRoleID: "e9Wl2JyVeBM_",
	}

	createdUser, err := u.Repo.Signup(user)
	if err != nil {
		return dto.SignupResponse{}, err
	}

	if err := util.Zoho(req.Email, user.Otp); err != nil {
		return dto.SignupResponse{}, err
	}

	response := dto.SignupResponse{
		ID:           createdUser.ID,
		Whatsapp:     createdUser.Whatsapp,
		Email:        req.Email, // Send back the plain email
		FullName:     createdUser.FullName,
		Password:     createdUser.Password,
		Token:        createdUser.Token,
		AccessRoleID: createdUser.AccessRoleID,
	}

	return response, nil
}

func (u *authService) sendMail(to, verificationToken string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "aseprayana95@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Test Email")
	mailer.SetBody("text/html", "Hello, this is a test email from MailHog.")

	dialer := gomail.NewDialer("localhost", 1025, "", "")

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

// sendVerificationEmail mengirim email verifikasi ke alamat email.
func (u *authService) sendVerificationEmail(to, verificationToken string) error {
	// Konfigurasi pengaturan email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "aseprayana95@gmail.com") // Ganti dengan alamat email Gmail pengirim
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Account Verification")
	mailer.SetBody("text/html", "Click the following link to verify your account: "+
		util.GetVerificationLink(verificationToken))

	// Konfigurasi pengaturan koneksi email untuk Gmail
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "aseprayana95@gmail.com", "tybm gndz imkq deev")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Hanya gunakan ini dalam pengembangan, tidak aman untuk produksi

	// Kirim email
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

func GenerateRandomNumeric(length int) string {
	const chars = "0123456789"

	var result strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		result.WriteRune(rune(chars[rand.Intn(len(chars))]))
	}

	return result.String()
}

// Fungsi untuk menghasilkan ID unik dengan Blake2 dan panjang 12 karakter
func generateSecureID() (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-="

	// Generate a salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Combine salt and current timestamp for uniqueness
	timestamp := time.Now().UnixNano()
	saltedID := fmt.Sprintf("%x%d", salt, timestamp)

	// Hash the combination using Blake2
	hash, err := blake2b.New512(nil)
	if err != nil {
		return "", err
	}
	hash.Write([]byte(saltedID))
	hashBytes := hash.Sum(nil)

	// Convert hash bytes into a valid string
	var secureID []byte
	for i := 0; i < 12; i++ {
		secureID = append(secureID, chars[hashBytes[i]%byte(len(chars))])
	}

	return string(secureID), nil
}
