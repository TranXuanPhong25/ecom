// services/product_service.go
package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/TranXuanPhong25/ecom/services/carts/configs"
	"github.com/TranXuanPhong25/ecom/services/carts/dtos"
	pb "github.com/TranXuanPhong25/ecom/services/shops/proto"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IShopsService interface {
	GetShopsByIds(shopIDs []string) (*dtos.GetShopsResponse, error)
	CloseConnection()
	// CheckStock(ctx context.Context, productID string) (int, error)
}
type ShopsService struct {
	rpcClient pb.ShopsServiceClient
	baseAddr  string
	conn      *grpc.ClientConn
	timeout   time.Duration
}

func NewShopsService(config *configs.ServiceConfig, shopsOnce *sync.Once) IShopsService {
	var client pb.ShopsServiceClient
	var conn *grpc.ClientConn
	shopsOnce.Do(func() {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		var err error
		conn, err := grpc.NewClient(config.ShopsServiceAddr, opts...)
		if err != nil {
			log.Errorf(fmt.Sprintf("failed to dial: %v", err))
		}

		client = pb.NewShopsServiceClient(conn)
		log.Infof("Successfully connected to users service at %s", config.ShopsServiceAddr)

	})
	return &ShopsService{
		baseAddr:  config.ShopsServiceAddr,
		rpcClient: client,
		conn:      conn,
		timeout:   config.Timeout,
	}
}

func (s *ShopsService) GetShopsByIds(shopIDs []string) (*dtos.GetShopsResponse, error) {
	if s.rpcClient == nil {
		return nil, fmt.Errorf("shops service client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	r, err := s.rpcClient.GetShopsByIDs(ctx, &pb.GetShopsByIDsRequest{Ids: shopIDs})
	if err != nil {
		return nil, err
	}
	shops := make([]dtos.Shop, len(r.GetShops()))
	for i, shop := range r.GetShops() {
		shops[i] = dtos.Shop{
			ID:   shop.GetId(),
			Name: shop.GetName(),
		}
	}
	return &dtos.GetShopsResponse{
		Shops:       shops,
		NotFoundIDs: r.GetNotFoundIds(),
	}, nil
}

func (s *ShopsService) CloseConnection() {
	log.Info("Closing shops service connection...")
	if s.conn != nil {
		if err := s.conn.Close(); err != nil {
			log.Errorf("Error closing shops service connection: %v", err)
		} else {
			log.Info("Shop service connection closed successfully")
		}
	}
}
