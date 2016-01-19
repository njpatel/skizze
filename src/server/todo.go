package server

import (
	"golang.org/x/net/context"

	pb "datamodel"
)

func (s *serverStruct) SetDefaults(ctx context.Context, in *pb.Defaults) (*pb.Defaults, error) {
	return nil, nil
}
func (s *serverStruct) GetDefaults(ctx context.Context, in *pb.Empty) (*pb.Defaults, error) {
	return nil, nil
}
func (s *serverStruct) DeleteDomain(ctx context.Context, in *pb.Domain) (*pb.Empty, error) {
	return nil, nil
}
func (s *serverStruct) GetDomain(ctx context.Context, in *pb.Domain) (*pb.Domain, error) {
	return nil, nil
}
func (s *serverStruct) GetSketch(ctx context.Context, in *pb.Sketch) (*pb.Sketch, error) {
	return nil, nil
}
func (s *serverStruct) CreateSnapshot(ctx context.Context, in *pb.CreateSnapshotRequest) (*pb.CreateSnapshotReply, error) {
	return nil, nil
}
func (s *serverStruct) GetSnapshot(ctx context.Context, in *pb.GetSnapshotRequest) (*pb.GetSnapshotReply, error) {
	return nil, nil
}
func (s *serverStruct) List(ctx context.Context, in *pb.ListRequest) (*pb.ListReply, error) {
	return nil, nil
}
