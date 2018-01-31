package main

import (
	"log"
	"net"

	pb "github.com/maurofran/shipper/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// Repository is the interface used to provide a consignment repository.
type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type repository struct {
	consignments []*pb.Consignment
}

func (r *repository) Create(c *pb.Consignment) (*pb.Consignment, error) {
	update := append(r.consignments, c)
	r.consignments = update
	return c, nil
}

func (r *repository) GetAll() []*pb.Consignment {
	return r.consignments
}

type service struct {
	repository Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repository.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	consignments := s.repository.GetAll()
	return &pb.GetResponse{Consignments: consignments}, nil
}

func main() {
	repo := &repository{}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repository: repo})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
