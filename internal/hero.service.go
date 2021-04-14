package internal

import (
	"context"

	"google.golang.org/grpc"

	pb "hk01/go-rpc-service/gen"
)

type Hero struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type HeroServer struct {
	pb.UnimplementedHeroesServiceServer
}

func NewServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	pb.RegisterHeroesServiceServer(s, &HeroServer{})
	return s
}

func (s *HeroServer) FindOne(ctx context.Context, in *pb.HeroById) (*pb.Hero, error) {
	Heros := [2]Hero{{1, "John"}, {2, "Doe"}}
	var found Hero
	for _, h := range Heros {
		if h.ID == int32(in.Id) {
			found = h
			break
		}
	}
	return &pb.Hero{
		Id:   found.ID,
		Name: found.Name,
	}, nil
}
