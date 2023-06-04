package monkey_learn

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_instance(t *testing.T) {
	var te *test
	monkey.PatchInstanceMethod(reflect.TypeOf(te), "A", func(_ *test) string {
		return "test"
	})

	test := instance()
	expected := "test"

	assert.Equal(t, expected, test)
}
