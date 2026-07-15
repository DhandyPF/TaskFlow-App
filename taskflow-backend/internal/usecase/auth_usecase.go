package usecase

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"taskflow-backend/internal/domain"
	"taskflow-backend/internal/repository"
	"taskflow-backend/pkg/jwt"
	"taskflow-backend/pkg/password"
)

// ErrEmailTaken is returned when registering with an email already in use.
var ErrEmailTaken = errors.New("email already registered")

// ErrInvalidCredentials is returned when login credentials do not match.
var ErrInvalidCredentials = errors.New("invalid email or password")

// AuthUsecase implements registration and login flows.
type AuthUsecase struct {
	users     domain.UserRepository
	jwtSecret string
}

// NewAuthUsecase builds an AuthUsecase.
func NewAuthUsecase(users domain.UserRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{users: users, jwtSecret: jwtSecret}
}

// Register creates a new user account with a hashed password.
func (u *AuthUsecase) Register(req domain.RegisterRequest) (*domain.User, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	if _, err := u.users.FindByEmail(email); err == nil {
		return nil, ErrEmailTaken
	} else if !errors.Is(err, repository.ErrNotFound) {
		return nil, err
	}

	hash, err := password.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	user := &domain.User{
		ID:           uuid.NewString(),
		Name:         strings.TrimSpace(req.Name),
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := u.users.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// Login validates credentials and issues a JWT access token.
func (u *AuthUsecase) Login(req domain.LoginRequest) (string, *domain.User, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	user, err := u.users.FindByEmail(email)
	if errors.Is(err, repository.ErrNotFound) {
		return "", nil, ErrInvalidCredentials
	}
	if err != nil {
		return "", nil, err
	}

	if !password.Verify(user.PasswordHash, req.Password) {
		return "", nil, ErrInvalidCredentials
	}

	token, err := jwt.Generate(user.ID, u.jwtSecret)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}
