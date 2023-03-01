package repository

type HttpRepository struct {
}

func NewHttpRepository() *HttpRepository {
	return &HttpRepository{}
}

func (h *HttpRepository) Get() string {
	return "123"
}
