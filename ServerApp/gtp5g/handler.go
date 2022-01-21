package main

import (
	"context"

	"github.com/muthuramanecs03g/nfcli/gen-go/UpfGtp5g"
)

func (g *Gtp5gThriftServer) GetPdr(ctx context.Context, iface string, id int32) (_r *UpfGtp5g.GetPdrResponse, _err error) {
	return &UpfGtp5g.GetPdrResponse{ErrCode: 0,
		ErrMsg: "Error",
		Pdr:    nil}, nil
}

func (g *Gtp5gThriftServer) GetFar(ctx context.Context, iface string, id int32) (_r *UpfGtp5g.GetFarResponse, _err error) {
	return &UpfGtp5g.GetFarResponse{ErrCode: 0,
		ErrMsg: "Error",
		Far:    nil}, nil
}

func (g *Gtp5gThriftServer) GetQer(ctx context.Context, iface string, id int32) (_r *UpfGtp5g.GetQerResponse, _err error) {
	return &UpfGtp5g.GetQerResponse{ErrCode: 0,
		ErrMsg: "Error",
		Qer:    nil}, nil
}
