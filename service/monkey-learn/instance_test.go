package monkey_learn

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func Test_instance(t *testing.T) {
//	var te *test
//	monkey.PatchInstanceMethod(reflect.TypeOf(te), "A", func(_ *test) string {
//		return "test"
//	})
//
//	defer monkey.UnpatchInstanceMethod(reflect.TypeOf(te), "A")
//	test := instance()
//	expected := "test"
//
//	assert.Equal(t, expected, test)
//}

func Test_instance_gomonkey(t *testing.T) {
	var te *test
	patches := gomonkey.ApplyMethodFunc(te, "A", func() string {
		return "test"
	})

	defer patches.Reset()
	test := instance()
	expected := "test"

	assert.Equal(t, expected, test)
}
func Test_instance1(t *testing.T) {
	test := instance()
	expected := "aa"

	assert.Equal(t, expected, test)
}
