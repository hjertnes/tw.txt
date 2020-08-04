package mocks

import (
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"github.com/stretchr/testify/mock"
)

type LoadFeedsMock struct {
	mock.Mock
}

func (l *LoadFeedsMock) Execute() []models.Feed{
	args := l.Called()

	return args.Get(constants.FirstArgument).([]models.Feed)
}


