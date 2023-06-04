package monkey_learn

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getUserOrders(t *testing.T) {
	monkey.Patch(getAdmin, func() string {
		return "test"
	})
	orders := getUsers()

	monkey.Unpatch(getAdmin)
	expected := []string{"test"}
	assert.Equal(t, expected, orders)
}

func Test_getUserOrders1(t *testing.T) {
	orders := getUsers()
	expected := []string{"admin"}
	assert.Equal(t, expected, orders)
}
