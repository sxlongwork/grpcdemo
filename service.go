package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "gorpc.com/grpcdemo2/services"
)

type ProdService struct {
}

func (prods *ProdService) GetProdStock(ctx context.Context, req *pb.ProdRequest) (*pb.ProdResponse, error) {

	var res *pb.ProdResponse = &pb.ProdResponse{Stock: 32}
	// id := req.GetProdId()
	// fmt.Printf("product %d has total %d\n", id, res.Stock)
	return res, nil
}

func main() {
	var (
		lis net.Listener
		err error
	)
	server := grpc.NewServer()
	pb.RegisterProdServiceServer(server, new(ProdService))

	if lis, err = net.Listen("tcp", "localhost:8010"); err != nil {
		fmt.Println("listen error.", err)
	}
	fmt.Println("server is listen in localhost:8010")
	if err = server.Serve(lis); err != nil {
		fmt.Println("server running error.", err)
	}
}
