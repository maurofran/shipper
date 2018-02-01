package main

import (
	"fmt"
	"log"

	pb "github.com/maurofran/shipper/consignment-service/proto/consignment"
	vp "github.com/maurofran/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
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
	repository   Repository
	vesselClient vp.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vp.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s\n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}
	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := s.repository.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	consignments := s.repository.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &repository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repository: repo})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
