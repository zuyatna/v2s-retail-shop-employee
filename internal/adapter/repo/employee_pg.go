package repo

import (
	"context"
	"database/sql"
	"employee-service/internal/domain"
)

type EmployeePG struct{ db *sql.DB }

func NewEmployeePG(db *sql.DB) *EmployeePG { return &EmployeePG{db: db} }

func (r *EmployeePG) Create(ctx context.Context, e *domain.Employee) (string, error) {
	q := `INSERT INTO employees (
		id, name, email, password_hash, access_level, position, work_location, personal_id, 
		address, zip_code, province, city, district, phone_number, photo_url, npwp, 
		bank_name, bank_account, salary, status, join_date, created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8,
		$9, $10, $11, $12, $13, $14, $15, $16,
		$17, $18, $19, $20, $21, $22, $23
	)`

	_, err := r.db.ExecContext(ctx, q,
		e.ID, e.Name, e.Email, e.PasswordHash, e.AccessLevel, e.Position, e.WorkLocation, e.PersonalID,
		e.Address, e.ZipCode, e.Province, e.City, e.District, e.PhoneNumber, e.PhotoURL, e.NPWP,
		e.BankName, e.BankAccount, e.Salary, e.Status, e.JoinDate, e.CreatedAt, e.UpdatedAt,
	)
	return e.ID.String(), err
}

func (r *EmployeePG) GetByID(ctx context.Context, id string) (*domain.Employee, error) {
	q := `SELECT id, name, email, password_hash, access_level, position, work_location, personal_id,
		address, zip_code, province, city, district, phone_number, photo_url, npwp,
		bank_name, bank_account, salary, status, join_date, created_at, updated_at
		FROM employees WHERE id = $1`
	row := r.db.QueryRowContext(ctx, q, id)

	var e domain.Employee
	if err := row.Scan(
		&e.ID, &e.Name, &e.Email, &e.PasswordHash, &e.AccessLevel, &e.Position, &e.WorkLocation, &e.PersonalID,
		&e.Address, &e.ZipCode, &e.Province, &e.City, &e.District, &e.PhoneNumber, &e.PhotoURL, &e.NPWP,
		&e.BankName, &e.BankAccount, &e.Salary, &e.Status, &e.JoinDate, &e.CreatedAt, &e.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *EmployeePG) GetByEmail(ctx context.Context, email string) (*domain.Employee, error) {
	q := `SELECT id, name, email, password_hash, access_level, position, work_location, personal_id,
		address, zip_code, province, city, district, phone_number, photo_url, npwp,
		bank_name, bank_account, salary, status, join_date, created_at, updated_at
		FROM employees WHERE email = $1`
	row := r.db.QueryRowContext(ctx, q, email)

	var e domain.Employee
	if err := row.Scan(
		&e.ID, &e.Name, &e.Email, &e.PasswordHash, &e.AccessLevel, &e.Position, &e.WorkLocation, &e.PersonalID,
		&e.Address, &e.ZipCode, &e.Province, &e.City, &e.District, &e.PhoneNumber, &e.PhotoURL, &e.NPWP,
		&e.BankName, &e.BankAccount, &e.Salary, &e.Status, &e.JoinDate, &e.CreatedAt, &e.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *EmployeePG) Update(ctx context.Context, e *domain.Employee) error {
	q := `UPDATE employees SET
		name = $1, email = $2, password_hash = $3, access_level = $4, position = $5, work_location = $6,
		personal_id = $7, address = $8, zip_code = $9, province = $10, city = $11, district = $12,
		phone_number = $13, photo_url = $14, npwp = $15, bank_name = $16, bank_account = $17,
		salary = $18, status = $19, join_date = $20, updated_at = $21
		WHERE id = $22`

	_, err := r.db.ExecContext(ctx, q,
		e.Name, e.Email, e.PasswordHash, e.AccessLevel, e.Position, e.WorkLocation,
		e.PersonalID, e.Address, e.ZipCode, e.Province, e.City, e.District,
		e.PhoneNumber, e.PhotoURL, e.NPWP, e.BankName, e.BankAccount,
		e.Salary, e.Status, e.JoinDate, e.UpdatedAt,
		e.ID,
	)
	return err
}

func (r *EmployeePG) List(ctx context.Context, offset, limit int) ([]*domain.Employee, error) {
	q := `SELECT id, name, email, password_hash, access_level, position, work_location, personal_id,
		address, zip_code, province, city, district, phone_number, photo_url, npwp,
		bank_name, bank_account, salary, status, join_date, created_at, updated_at
		FROM employees ORDER BY created_at DESC OFFSET $1 LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*domain.Employee
	for rows.Next() {
		var e domain.Employee
		if err := rows.Scan(
			&e.ID, &e.Name, &e.Email, &e.PasswordHash, &e.AccessLevel, &e.Position, &e.WorkLocation, &e.PersonalID,
			&e.Address, &e.ZipCode, &e.Province, &e.City, &e.District, &e.PhoneNumber, &e.PhotoURL, &e.NPWP,
			&e.BankName, &e.BankAccount, &e.Salary, &e.Status, &e.JoinDate, &e.CreatedAt, &e.UpdatedAt,
		); err != nil {
			return nil, err
		}
		employees = append(employees, &e)
	}
	return employees, nil
}
