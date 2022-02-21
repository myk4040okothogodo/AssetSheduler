package server

import (
    "context"
    "github.com/hashicorp/go-hclog"
    protos "github.com/myk4040okothogodo/ResourceSheduler/protos/assetsdefinition/protos"

)

// AssetAvail is a gRPC server, it implments the methods by the AssetAvailServer interface
type AssetAvailability struct {
  log hclog.Logger
}


func NewAssetAvailability(l hclog.Logger) *AssetAvailability{
    return &AssetAvailability{l}

}

func (a *AssetAvailability) CheckForAssetAvailability(ctx context.Context,aa *protos.GetAssetAvailabilityRequest) (*protos.AssetAvailabilityResponse, error) {
    a.log.Info("Handle request for CheckAssetAvailability", "asset", aa.GetAssetchecked(),"owner",aa.GetOwner())
    return &protos.AssetAvailabilityResponse {
      Avalability : &protos.Asset {
      Availability : true, },
      Name : &protos.Asset {
      Name : "LandCruiser 300 ", },
      Owner : &protos.Asset {
      Owner : &protos.Person {
          FirstName: "kevin ",
          LastName: "mwendwa",
          Id      : 3056,
      },},
      Timeslots : &protos.Asset {
      Timeslots : &protos.Timeslots{
          Month: "June",
          Day  : "23rd",
          Hours : 6.30,
      },
    },
    }, nil
}  
