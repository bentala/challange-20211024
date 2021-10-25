package repository

import (
	"github.com/stretchr/testify/mock"
)

type MockDataRepository struct {
	mock.Mock
}

func (m *MockDataRepository) Get(words string, page int) (interface{}, error) {
	args := m.Called(words, page)
	return args.Get(0).(interface{}), args.Error(1)
}
