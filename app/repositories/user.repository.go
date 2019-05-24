package repositories

import (
	"github.com/hirsch88/go-trophy-server/app/models"
	"github.com/hirsch88/go-trophy-server/app/providers"
	"go.uber.org/zap"
)

type userRepository struct {
	database providers.DatabaseProvider
	log      *zap.SugaredLogger
}

func (r *userRepository) Create(user models.User) models.User {
	r.log.Info("STARTING UserRepository.create()")
	db := r.database.Connect()
	defer db.Close()

	db.Create(&user)
	r.log.Info("FINISHED UserRepository.create()")
	return user
}

type UserRepository interface {
	Create(user models.User) models.User
}

func NewUserRepository(database providers.DatabaseProvider, log *zap.SugaredLogger) UserRepository {
	return &userRepository{
		database, log,
	}
}
