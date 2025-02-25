package usecases

import (
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/domain/repositories"

	"github.com/google/uuid"
)

type UserUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(name string, Point int) (*entities.User, error) {
	user := &entities.User{
		ID:     uuid.New(),
		Name:   name,
		Points: Point,
	}
	err := uc.repo.Create(user)
	return user, err
}

func (uc *UserUseCase) GetUserByID(id uuid.UUID) (*entities.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUseCase) GetAllUsers() ([]entities.User, error) {
	return uc.repo.GetAll()
}

func (uc *UserUseCase) DeleteVoucher(id uuid.UUID) error {
	return uc.repo.Delete(id)
}

func (uc *UserUseCase) UpdatePoints(id uuid.UUID, points int) error {
	return uc.repo.UpdatePoints(id, points)
}
