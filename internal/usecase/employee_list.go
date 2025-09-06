package usecase

import (
	"context"
	"employee-service/internal/domain"
	"errors"
)

var ErrForbidden = errors.New("forbidden")

type ListEmployeeInput struct {
	AccessLevel domain.AccessLevel
	Offset      int
	Limit       int
}

type EmployeeLister interface {
	List(ctx context.Context, in ListEmployeeInput) ([]*domain.Employee, error)
}

type EmployeeListUsecase struct {
	repo domain.EmployeeRepository
}

func NewEmployeeListUsecase(r domain.EmployeeRepository) *EmployeeListUsecase {
	return &EmployeeListUsecase{repo: r}
}

func (uc *EmployeeListUsecase) List(ctx context.Context, in ListEmployeeInput) ([]*domain.Employee, error) {
	if !domain.CanListEmployees(in.AccessLevel) {
		return nil, ErrForbidden
	}
	if in.Limit <= 0 || in.Limit > 100 {
		in.Limit = 50
	}
	if in.Offset < 0 {
		in.Offset = 0
	}

	return uc.repo.List(ctx, in.Offset, in.Limit)
}
