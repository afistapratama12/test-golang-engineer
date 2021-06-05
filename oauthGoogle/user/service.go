package user

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetUserByEmail(email string) (User, error)
	CreateNewUser(input GoogleLoginInput) (User, error)
}

var (
	err          = godotenv.Load()
	passGenerate = os.Getenv("PASS_USER_GENERATE")
)

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CreateNewUser(input GoogleLoginInput) (User, error) {
	var newUser User

	user, err := s.GetUserByEmail(input.Email)

	if user.Email == input.Email {
		return user, nil
	} else {
		hashPass, _ := bcrypt.GenerateFromPassword([]byte(passGenerate), bcrypt.DefaultCost)

		var userInput = User{
			Name:      input.Name,
			Email:     input.Email,
			Password:  string(hashPass),
			Avatar:    input.Picture,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		valUser := &newUser
		*valUser, err = s.repository.Create(userInput)

		if err != nil {
			return *valUser, err
		}
	}
	return newUser, nil
}
