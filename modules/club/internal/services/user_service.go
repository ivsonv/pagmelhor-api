package services

import (
	responsesUsers "app/modules/club/domain/dto/responses/users"
	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/utils"
	"context"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) services.IUserServices {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) Get(ctx context.Context) ([]responsesUsers.GetUserResponseDto, error) {
	users, err := s.userRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	response := utils.Map(users, func(i int) responsesUsers.GetUserResponseDto {
		return responsesUsers.GetUserResponseDto{
			Email: users[i].Email,
			Name:  users[i].Name,
			ID:    users[i].ID,
		}
	})

	return response, nil
}
