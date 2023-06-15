package monkey_learn

import (
	"bou.ke/monkey"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	other_folder "go-playground/service/monkey-learn/other-folder"
	"testing"
)

func Test_getUserOrders(t *testing.T) {
	monkey.Patch(getAdmin, func() string {
		return "test"
	})
	defer monkey.Unpatch(getAdmin)
	orders := getUsers()

	expected := []string{"test"}
	assert.Equal(t, expected, orders)
}

func Test_getUserOrders_gomonkey(t *testing.T) {
	patches := gomonkey.ApplyFunc(getAdmin, func() string {
		return "test"
	})
	defer patches.Reset()

	orders := getUsers()
	expected := []string{"test"}
	assert.Equal(t, expected, orders)
}

func Test_getUserOrders1(t *testing.T) {
	orders := getUsers()
	expected := []string{"admin"}
	assert.Equal(t, expected, orders)
}

func Test_getOtherUserOrders(t *testing.T) {
	monkey.Patch(other_folder.GetAdmin1, func() string {
		return "test"
	})
	defer monkey.Unpatch(other_folder.GetAdmin1)
	orders := getOtherUsers()

	expected := []string{"test"}
	assert.Equal(t, expected, orders)
}

func Test_getOtherUserOrders_gomonkey(t *testing.T) {
	patches := gomonkey.ApplyFunc(other_folder.GetAdmin1, func() string {
		return "test"
	})
	defer patches.Reset()
	orders := getOtherUsers()

	expected := []string{"test"}
	assert.Equal(t, expected, orders)
}

func Test_getOtherUserOrders1(t *testing.T) {
	orders := getOtherUsers()
	expected := []string{"admin1"}
	assert.Equal(t, expected, orders)
}
