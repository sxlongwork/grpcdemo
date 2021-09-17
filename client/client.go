package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	pb "gorpc.com/grpcdemo2/services"
)

func main() {
	var (
		conn *grpc.ClientConn
		err  error
		cli  pb.ProdServiceClient
		req  *pb.ProdRequest
		res  *pb.ProdResponse
	)
	// 连接grpcserver
	if conn, err = grpc.Dial("localhost:8010", grpc.WithInsecure()); err != nil {
		fmt.Println("connect grpc server error.", err)
	}
	defer conn.Close()
	cli = pb.NewProdServiceClient(conn)

	// 调用接口
	req = &pb.ProdRequest{ProdId: 1}
	if res, err = cli.GetProdStock(context.TODO(), req); err != nil {
		fmt.Println("access grpc interface GetProdStock error.", err)
	}
	fmt.Printf("product %d has total %d stock.\n", req.GetProdId(), res.GetStock())

}
