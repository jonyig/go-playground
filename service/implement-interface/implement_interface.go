package implement_interface

type MyInterface interface {
	Test() string
	Test1() int
}

type MyStructPoint struct{}

var _ MyInterface = (*MyStructPoint)(nil)

func (h *MyStructPoint) Test() string {
	return "123"
}
func (h MyStructPoint) Test1() int {
	return 123
}

type MyStructValue struct{}

var _ MyInterface = (*MyStructValue)(nil)
var _ MyInterface = MyStructValue{}

func (h MyStructValue) Test() string {
	return "123"
}
func (h MyStructValue) Test1() int {
	return 123
}
