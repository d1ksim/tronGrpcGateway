package main

import (
	"context"
	"errors"
	gw "github.com/d1mpi/grpc-tron/api"
	"github.com/d1mpi/tronGrpcGateway/database"
	"github.com/d1mpi/tronGrpcGateway/database/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

	database.InitPostgresDatabase()
	log.Println("Init DataBase successful")

	lis, err := net.Listen("tcp", ":"+viper.GetString("server.listen"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor), grpc.StreamInterceptor(StreamServerInterceptor))
	log.Printf("WalletServer listening at %v", lis.Addr())

	var client = InitGrpcClient()
	log.Println("Tron GRPC successful started!")

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

	token, err := jwt.Parse(md.Get("TRON-PRO-API-KEY")[0], func(token *jwt.Token) (interface{}, error) {
		return []byte("U3o41BzEnYuzwIGyg3Sg0kPz_VBk9NDVPlbojc4uEjU"), nil
	})

	var jwtFirst models.JwtData
	resultJwt := database.DataBase.First(&jwtFirst, "access_token = ?", md.Get("TRON-PRO-API-KEY")[0])
	if resultJwt.RowsAffected == 0 {
		return nil, errors.New("the jwt not exists")
	}

	var accountFind models.Accounts
	resultAccount := database.DataBase.First(&accountFind, "telegram_id = ?", jwtFirst.TelegramID)
	if resultAccount.RowsAffected == 0 {
		return nil, errors.New("the account not exists")
	}

	if !accountFind.Active {
		return nil, errors.New("the account not active")
	}

	switch {
	case token.Valid:
		resp, err = handler(ctx, req)
		return resp, err
	case errors.Is(err, jwt.ErrTokenMalformed):
		return nil, errors.New("that's not even a token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, errors.New("invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, errors.New("timing is everything")
	default:
		return nil, err
	}
}

func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)
	return err
}
