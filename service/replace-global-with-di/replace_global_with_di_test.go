package replace_global_with_di

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	testObj := Test{Data: "456"}
	testObj.On("Name").Return("123")

	expectation := "123"

	assert.Equal(t, expectation, testObj.Name())
}
