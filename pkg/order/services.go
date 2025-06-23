package contact

import (
	"context"
	"time"

	"srv_order/main/api/presenter"
	"srv_order/main/pkg/entities"
	pb "srv_order/main/proto/contact"
)

type Service interface {
	InsertOrder(order *entities.Order) (*entities.Order, error)
	FetchOrders() (*[]presenter.Order, error)
	GetContactFromGRPC(id string) (*pb.ContactResponse, error)
}

type service struct {
	repository Repository
	grpcClient pb.ContactServiceClient
}

func NewService(r Repository, grpcClient pb.ContactServiceClient) Service {
	return &service{
		repository: r,
		grpcClient: grpcClient,
	}
}
func (s *service) InsertOrder(order *entities.Order) (*entities.Order, error) {
	return s.repository.CreateOrder(order)
}

func (s *service) FetchOrders() (*[]presenter.Order, error) {
	return s.repository.ReadOrder()
}

func (s *service) GetContactFromGRPC(id string) (*pb.ContactResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	resp, err := s.grpcClient.GetContactByID(ctx, &pb.GetContactRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
