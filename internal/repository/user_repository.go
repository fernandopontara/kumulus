package repository

import (
	"database/sql"
	"kumulus/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `SELECT id, organization_id, name, email, password_hash, role FROM users WHERE email = $1`
	row := r.DB.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.OrganizationID, &user.Name, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
