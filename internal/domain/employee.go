package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AccessLevel string

const (
	AccessOrdinary   AccessLevel = "ordinary"
	AccessSupervisor AccessLevel = "supervisor"
	AccessManager    AccessLevel = "manager"
	AccessHR         AccessLevel = "hr"
	AccessIntern     AccessLevel = "intern"
)

type Employee struct {
	ID           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	Password     string      `json:"-"`
	AccessLevel  AccessLevel `json:"access_level"`
	Position     string      `json:"position"`
	WorkLocation string      `json:"work_location"`
	PersonalID   string      `json:"personal_id"`
	Address      string      `json:"address"`
	ZipCode      string      `json:"zip_code"`
	Province     string      `json:"province"`
	City         string      `json:"city"`
	District     string      `json:"district"`
	PhoneNumber  string      `json:"phone_number"`
	PhotoURL     string      `json:"photo_url"`
	NPWP         string      `json:"npwp"`
	BankName     string      `json:"bank_name"`
	BankAccount  string      `json:"bank_account"`
	Salary       float64     `json:"salary"`
	Status       string      `json:"status"`
	JoinDate     time.Time   `json:"join_date"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type EmployeeResponse struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	AccessLevel AccessLevel `json:"access_level"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type EmployeeRepository interface {
	Create(ctx context.Context, e *Employee) (string, error)
	GetByID(ctx context.Context, id string) (*Employee, error)
	GetByEmail(ctx context.Context, email string) (*Employee, error)
	Update(ctx context.Context, e *Employee) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Employee, error)
}
