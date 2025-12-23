package services

import (
	"server/internal/dto"
	"sync"

	"github.com/google/uuid"
)

func BulkCreateUsers(users []dto.CreateUserRequest, createdBy uuid.UUID, accessToken string) *dto.BulkCreateUsersResponse {
	response := &dto.BulkCreateUsersResponse{
		Results: make([]dto.BulkCreateUserResult, 0, len(users)),
	}

	// Usar goroutines con límite para procesar en paralelo
	maxConcurrent := 5
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, userReq := range users {
		wg.Add(1)
		sem <- struct{}{} // Adquirir semáforo

		go func(req dto.CreateUserRequest) {
			defer wg.Done()
			defer func() { <-sem }() // Liberar semáforo

			result := dto.BulkCreateUserResult{
				DNI:   req.DNI,
				Email: req.Email,
			}

			userResp, err := CreateUser(req, createdBy, accessToken)
			if err != nil {
				result.Success = false
				result.Error = err.Error()
			} else {
				result.Success = true
				result.User = userResp
			}

			mu.Lock()
			response.Results = append(response.Results, result)
			if result.Success {
				response.SuccessCount++
			} else {
				response.FailureCount++
			}
			mu.Unlock()
		}(userReq)
	}

	wg.Wait()
	return response
}