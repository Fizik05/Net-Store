package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type Storage interface {
	GetUser(ctx context.Context, email string) (string, error)
	SaveUser(ctx context.Context, name, email, password string) error
}

type User struct {
	ctx     context.Context
	logger  *slog.Logger
	storage Storage
}

func NewUser(ctx context.Context, log *slog.Logger, storage Storage) *User {
	return &User{
		ctx:     ctx,
		logger:  log,
		storage: storage,
	}
}

func (u *User) Login(ctx context.Context, email, password string) (string, error) {
	const fn = "Login"

	hash, err := u.storage.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("%s : There's not user with this email", fn)
		}
		return "", fmt.Errorf("%s : %w", fn, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("%s : The password is wrong", fn)
	}

	u.logger.Info("user logged in", "email", email)

	// TODO: generate token

	return "", nil
}

func (u *User) Register(ctx context.Context, name, email, password string) error {
	const fn = "Login"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}

	err = u.storage.SaveUser(ctx, name, email, string(hash))
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}

	return nil
}
