package services

import (
	"github.com/hirsch88/go-trophy-server/app/mail"
	"github.com/hirsch88/go-trophy-server/app/models"
	"github.com/hirsch88/go-trophy-server/app/providers"
	"github.com/hirsch88/go-trophy-server/app/repositories"
	"go.uber.org/zap"
)

type userService struct {
	userRepository repositories.UserRepository
	mailProvider   providers.MailProvider
	log *zap.SugaredLogger
}

func (s *userService) Create(user models.User) models.User {
	s.log.Info("STARTING UserService.create()")

	newUser := s.userRepository.Create(user)
	s.mailProvider.Send(mail.NewUserCreatedMail(user), user.Email)

	s.log.Info("FINISHED UserService.create()")
	return newUser
}

type UserService interface {
	Create(user models.User) models.User
}

func NewUserService(userRepository repositories.UserRepository, mailProvider providers.MailProvider, log *zap.SugaredLogger) UserService {
	return &userService{
		userRepository,
		mailProvider,
		log,
	}
}
