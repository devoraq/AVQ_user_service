package service

import (
	"context"
	"log/slog"

	dtoV1 "github.com/DENFNC/awq_user_service/internal/adapters/dto/v1"
	"github.com/DENFNC/awq_user_service/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Save(
		ctx context.Context,
		agg *domain.UserAggregate,
	) (string, error)
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
	dto *dtoV1.UserCreateDTO,
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

	uid, err := svc.UserRepository.Save(ctx,
		&domain.UserAggregate{
			User: &domain.User{
				Nickname: dto.Nickname,
			},
			PrivateData: &domain.PrivateData{
				DateOfBirth: dto.Birthday,
			},
			SecurityData: &domain.SecurityData{
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

	return uid, nil
}

func (svc *UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func (svc *UserService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
