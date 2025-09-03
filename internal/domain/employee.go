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
	AccessIntern     AccessLevel = "intern"
)

// TODO: Detailing employee fields such as address, phone number, position, etc.
type Employee struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Password    string      `json:"-"`
	AccessLevel AccessLevel `json:"access_level"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
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
