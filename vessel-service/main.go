package main

import (
	"errors"
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/maurofran/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type vesselRepository struct {
	vessels []*pb.Vessel
}

func (r *vesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range r.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by spec")
}

type service struct {
	repo *vesselRepository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	repo := &vesselRepository{vessels: vessels}
	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo: repo})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
