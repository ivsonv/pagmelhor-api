package users

import (
	responses "app/modules/club/domain/dto/responses/users"
	"app/modules/club/domain/results"
	"app/modules/club/utils"
	"context"
	"log"
)

func (s *UserService) Get(ctx context.Context) results.Result[[]responses.GetUserResponseDto] {
	items, err := s.userRepository.Get(ctx)
	if err != nil {
		log.Printf("error getting service users: %v", err)
		return results.Failure[[]responses.GetUserResponseDto](ErrGetUsers)
	}

	return results.Success(
		utils.Map(items, func(i int) responses.GetUserResponseDto {
			return responses.GetUserResponseDto{
				Email: items[i].Email,
				Name:  items[i].Name,
				ID:    items[i].ID,
			}
		}),
	)
}
