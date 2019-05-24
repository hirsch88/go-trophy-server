package services

import (
	"github.com/hirsch88/go-trophy-server/app/models"
	"github.com/hirsch88/go-trophy-server/app/providers"
	"github.com/hirsch88/go-trophy-server/config"
	"github.com/hirsch88/go-trophy-server/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func testSetup(t *testing.T) (UserService, models.User) {
	var user = models.User{
		Username: "bubu",
		Email:    "bubu@test.ch",
	}

	var mailProviderMock = new(mocks.MailProvider)
	var userRepositoryMock = new(mocks.UserRepository)
	var logProvider = providers.NewLoggerProvider(config.NewAppConfig())

	userRepositoryMock.On("Create", mock.Anything).Return(user)
	mailProviderMock.On("Send", mock.Anything, mock.Anything).Return(true)

	var service = NewUserService(userRepositoryMock, mailProviderMock, logProvider)
	return service, user
}

func TestCreate(t *testing.T) {
	service, user := testSetup(t)

	returnedUser := service.Create(user)

	assert.Equal(t, "bubu", returnedUser.Username)
	assert.Equal(t, "bubu@test.ch", returnedUser.Email)
}
