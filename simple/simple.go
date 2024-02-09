package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

// Constructor for SimpleRepository
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

// Constructor for SimpleService
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}

}
