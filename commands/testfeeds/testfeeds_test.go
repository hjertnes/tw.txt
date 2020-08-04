package testfeeds

import (
	"git.sr.ht/~hjertnes/tw.txt/mocks"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"testing"
)

func TestTest(t *testing.T) {
	m := &mocks.LoadFeedsMock{}

	m.On("Execute").Return([]models.Feed{
		models.Feed{},
	})

	New(m).Execute()
}