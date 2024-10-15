package postgres

import (
	"context"
	"fmt"

	"letual/internal/models"
)

func (s *Storage) GetUser(ctx context.Context, email string) (string, error) {
	const fn = "GetUser"

	const query = `
		SELECT id, name, email, password
		FROM users
		WHERE email = $1`

	var user models.User

	err := s.client.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return "", fmt.Errorf("%s : %w", fn, err)
	}

	return user.Password, nil
}

func (s *Storage) SaveUser(ctx context.Context, name, email, password string) error {
	const fn = "SaveUser"

	const query = `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)`

	_, err := s.client.Exec(ctx, query, name, email, password)
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}

	return nil
}
