package monkey_learn

type test struct {
	aa string
	bb string
}

func (test *test) A() string {
	return test.aa
}

func (test *test) b() string {
	return test.bb
}

func instance() string {
	d := &test{
		aa: "aa",
	}

	return d.A()
}
