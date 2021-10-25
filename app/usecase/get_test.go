package usecase

import (
	"challange-20211024/app/entities"
	"challange-20211024/app/repository"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run("returns error to many results", func(t *testing.T) {
		s := "s"
		page := 2

		response := entities.Response{
			Error:    "To many results",
			Response: "Error",
		}

		mockDataGetter := repository.MockDataRepository{}
		mockDataGetter.On("Get", mock.Anything, mock.Anything).Return(response, nil)

		_, err := NewDataGetter().GetData(entities.Get{
			Search: s,
			Page:   page,
		})
		t.Errorf("response", err.Error())
	})
}
