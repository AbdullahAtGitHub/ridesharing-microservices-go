package main

import (
	"context"

	"github.com/AbdullahAtGitHub/ridesharing-microservices-go/services/trip-service/internal/domain"
	"github.com/AbdullahAtGitHub/ridesharing-microservices-go/services/trip-service/internal/infrastructure/repository"
	"github.com/AbdullahAtGitHub/ridesharing-microservices-go/services/trip-service/internal/service"
)

func main() {

	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewTripService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID: "user123",
	}

	svc.CreateTrip(ctx, fare)
}
