package replace_global_with_di

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"time"
)

type signer struct {
	now func() time.Time
}

func newSigner() *signer {
	return &signer{
		now: time.Now,
	}
}
func (s *signer) Sign(msg string) string {
	now := s.now()
	return signWithTime(msg, now)
}
func signWithTime(msg string, now time.Time) string {
	return fmt.Sprint(msg, now)
}

type Test struct {
	Data string
	mock.Mock
}

func (e *Test) Name() string {
	args := e.Called()
	return args.String(0)
}

func (e *Test) Write(text string) {
	e.Data = text
}
