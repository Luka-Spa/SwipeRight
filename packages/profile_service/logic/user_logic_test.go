package logic

import (
	"testing"
	"time"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// func (mock *MockRepository) Save(UserProfile *model.UserProfile) (*model.UserProfile, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.(*model.UserProfile), args.Error(1)
// }

func (mock *MockRepository) All() ([]model.UserProfile, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.UserProfile), args.Error(1)
}

func TestGetAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var id = 1
	var firstName = "Go"
	var	time = time.Now()

	userProfile := model.UserProfile{Id:id, Firstname: firstName, CreatedAt: time }
	// Setup expectations
	mockRepo.On("All").Return([]model.UserProfile{userProfile}, nil)

	testService := NewUserlogic(mockRepo)

	result := testService.GetAll()

	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, id, result[0].Id)
	assert.Equal(t, time, result[0].CreatedAt)
	assert.Equal(t, firstName, result[0].Firstname)
}
