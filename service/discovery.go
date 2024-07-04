package service

import (
	pb "discovery_servcie/genproto"
	"discovery_servcie/storage/postgres"
	"golang.org/x/net/context"
)

type DiscoveryService struct {
	CompositionMeta *postgres.DiscoveryRepository
	pb.UnimplementedDiscoveryServiceServer
}

func NewDiscoveryService(com *postgres.DiscoveryRepository) *DiscoveryService {
	return &DiscoveryService{CompositionMeta: com}
}

func (service *DiscoveryService) GetCompositionTrending(ctx context.Context, in *pb.Void) (*pb.DiscoveriesResponse, error) {
	response, err := service.CompositionMeta.GetCompositionTrending(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *DiscoveryService) GetCompositionRecommend(ctx context.Context, in *pb.Void) (*pb.DiscoveriesResponse, error) {
	response, err := service.CompositionMeta.GetCompositionRecommend(in)
	if err != nil {
		return nil, err
	}
	return response, err
}
func (service *DiscoveryService) GetCompositionGenre(ctx context.Context, in *pb.GetGenre) (*pb.DiscoveriesResponse, error) {
	response, err := service.CompositionMeta.GetDiscoveriesByGenre(in)
	if err != nil {
		return nil, err
	}
	return response, err
}
func (service *DiscoveryService) GetDiscovery(ctx context.Context, in *pb.GetDiscoveryRequest) (*pb.DiscoveriesResponse, error) {
	response, err := service.CompositionMeta.GetCompositionMetadata(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *DiscoveryService) DeleteCompositionLike(ctx context.Context, in *pb.LikeRequest) (*pb.Void, error) {
	response, err := service.CompositionMeta.DeleteCompositionLike(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *DiscoveryService) CreateCompositionLike(ctx context.Context, in *pb.LikeRequest) (*pb.Void, error) {
	response, err := service.CompositionMeta.CreateCompositionLike(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *DiscoveryService) CreateCompositionListen(ctx context.Context, in *pb.LikeRequest) (*pb.Void, error) {
	response, err := service.CompositionMeta.CreateCompositionListen(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
