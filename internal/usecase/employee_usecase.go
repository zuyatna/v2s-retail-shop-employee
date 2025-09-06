package usecase

import (
	"context"
	"employee-service/internal/domain"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeUsecase struct {
	repo    domain.EmployeeRepository
	timeout time.Duration
}

func NewEmployeeUsecase(r domain.EmployeeRepository, timeout time.Duration) *EmployeeUsecase {
	return &EmployeeUsecase{
		repo:    r,
		timeout: timeout,
	}
}

type CreateEmployeeInput struct {
	Name         string
	Email        string
	Password     string
	AccessLevel  domain.AccessLevel
	Position     string
	WorkLocation string
	PersonalID   string
	Address      string
	ZipCode      string
	Province     string
	City         string
	District     string
	PhoneNumber  string
	PhotoURL     string
	NPWP         string
	BankName     string
	BankAccount  string
	Salary       float64
	Status       string
}

func (uc *EmployeeUsecase) Create(ctx context.Context, in CreateEmployeeInput) (uuid.UUID, error) {
	if in.Name == "" || in.Email == "" || in.Password == "" {
		return uuid.Nil, errors.New("name/email/password are required")
	}
	if in.AccessLevel.Valid() {
		return uuid.Nil, errors.New("invalid access level")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	id := uuid.Must(uuid.NewV7())
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return uuid.Nil, err
	}
	now := time.Now().In(loc)

	e := &domain.Employee{
		ID:           id,
		Name:         in.Name,
		Email:        in.Email,
		PasswordHash: string(hash),
		AccessLevel:  in.AccessLevel,
		Position:     in.Position,
		WorkLocation: in.WorkLocation,
		PersonalID:   in.PersonalID,
		Address:      in.Address,
		ZipCode:      in.ZipCode,
		Province:     in.Province,
		City:         in.City,
		District:     in.District,
		PhoneNumber:  in.PhoneNumber,
		PhotoURL:     in.PhotoURL,
		NPWP:         in.NPWP,
		BankName:     in.BankName,
		BankAccount:  in.BankAccount,
		Salary:       in.Salary,
		Status:       in.Status,
		JoinDate:     now,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	cctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()
	if _, err := uc.repo.Create(cctx, e); err != nil {
		log.Printf("error creating employee: %v", err)
		return uuid.Nil, err
	}

	return uuid.Nil, nil
}

func (uc *EmployeeUsecase) GetByID(ctx context.Context, id string) (*domain.EmployeeResponse, error) {
	cctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()
	e, err := uc.repo.GetByID(cctx, id)
	if err != nil {
		return nil, err
	}

	res := &domain.EmployeeResponse{
		ID:          e.ID,
		Name:        e.Name,
		Email:       e.Email,
		AccessLevel: e.AccessLevel,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
	return res, nil
}

func (uc *EmployeeUsecase) GetByEmail(ctx context.Context, email string) (*domain.EmployeeResponse, error) {
	cctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()
	e, err := uc.repo.GetByEmail(cctx, email)
	if err != nil {
		return nil, err
	}

	res := &domain.EmployeeResponse{
		ID:          e.ID,
		Name:        e.Name,
		Email:       e.Email,
		AccessLevel: e.AccessLevel,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
	return res, nil
}

func (uc *EmployeeUsecase) Update(ctx context.Context, e *domain.Employee) error {
	if e.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if e.Name == "" || e.Email == "" || e.PasswordHash == "" {
		return errors.New("name/email/password are required")
	}
	if !e.AccessLevel.Valid() {
		return errors.New("invalid access level")
	}
	e.UpdatedAt = time.Now().In(time.FixedZone("Asia/Jakarta", 7*3600))

	cctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()
	return uc.repo.Update(cctx, e)
}
