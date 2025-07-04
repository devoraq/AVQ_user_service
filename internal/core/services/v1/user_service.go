package service

import (
	"context"
	"log/slog"

	"github.com/DENFNC/awq_user_service/internal/adapters/dto"
	"github.com/DENFNC/awq_user_service/internal/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Save(
		ctx context.Context,
		agg *domain.UserAggregate,
	) error
}

type UserService struct {
	*slog.Logger
	UserRepository
}

func NewUserService(
	log *slog.Logger,
	repo UserRepository,
) *UserService {
	return &UserService{
		Logger:         log,
		UserRepository: repo,
	}
}

func (svc *UserService) Create(
	ctx context.Context,
	dto *dto.CreateUserDTO,
) (string, error) {
	const op = "service.UserService.Create"

	log := svc.Logger.With("op", op)

	passHash, err := svc.hashPassword(dto.Password)
	if err != nil {
		log.Error(
			"Password could not be hashed",
			slog.String("err", err.Error()),
		)
		return "", err
	}

	uid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	err = svc.UserRepository.Save(ctx,
		&domain.UserAggregate{
			User: &domain.User{
				UID:      uid.String(),
				Nickname: dto.Nickname,
			},
			PrivateData: &domain.PrivateData{
				UserID:      uid.String(),
				DateOfBirth: dto.Birthday,
			},
			SecurityData: &domain.SecurityData{
				UserID:       uid.String(),
				Login:        dto.Nickname,
				Email:        dto.Email,
				PasswordHash: passHash,
			},
		})
	if err != nil {
		log.Error(
			"Произошла ошибка",
			slog.String("err", err.Error()),
		)
		return "", err
	}

	return uid.String(), nil
}

func (svc *UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// func (svc *UserService) checkPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
