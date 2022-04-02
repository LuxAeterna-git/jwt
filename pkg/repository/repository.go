package repository

type Authorization interface {
}

type Repository struct {
	Authorization
}

func NewSRepository() *Repository {
	return &Repository{}
}
