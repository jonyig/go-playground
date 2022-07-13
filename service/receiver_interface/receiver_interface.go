package receiver_interface

type S struct {
	Data string
}

func (s S) Read() string {
	return s.Data
}

func (s *S) Write(str string) {
	s.Data = str
}

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}
