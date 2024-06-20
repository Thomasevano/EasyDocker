package models_test

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/test_helpers"
	"testing"
)

func TestStackCreateInput(t *testing.T) {
	t.Run("Name should be required", func(t *testing.T) {
		stack := models.StackCreateInput{Name: ""}
		test_helpers.AsserNotValidation(t, stack)
	})
	t.Run("Name should be at least 3 characters", func(t *testing.T) {
		stack := models.StackCreateInput{Name: "ab"}
		test_helpers.AsserNotValidation(t, stack)
	})
	t.Run("Stack should be valid when name at least 3 characters", func(t *testing.T) {
		stack := models.StackCreateInput{Name: "abc"}
		test_helpers.AssertValidation(t, stack)
	})
}
