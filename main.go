package main

import (
	"context"
	gw "github.com/d1mpi/grpc-tron/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"tronGrpcGateway/bot"
)

type WalletServer struct {
	gw.WalletServer
	client gw.WalletClient
}

type WalletSolidityServer struct {
	gw.WalletSolidityServer
	client gw.WalletClient
}

type ExtensionServer struct {
	gw.WalletExtensionServer
	client gw.WalletClient
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s\n", err)
	}

	lis, err := net.Listen("tcp", ":"+viper.GetString("server.listen"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor), grpc.StreamInterceptor(StreamServerInterceptor))
	log.Printf("WalletServer listening at %v", lis.Addr())

	var client = InitGrpcClient()
	log.Println("Tron GRPC successful started!")

	bot.InitTelegramBot()

	gw.RegisterWalletServer(s, &WalletServer{
		client: client,
	})
	gw.RegisterWalletSolidityServer(s, &WalletSolidityServer{
		client: client,
	})
	gw.RegisterWalletExtensionServer(s, &ExtensionServer{
		client: client,
	})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	println(md.Get("TRON-PRO-API-KEY")[0])
	resp, err = handler(ctx, req)
	return resp, err
}

func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)
	return err
}
