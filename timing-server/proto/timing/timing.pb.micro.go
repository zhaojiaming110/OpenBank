// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: timing.proto

package micro_open_bank_service_timing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Timing service

type TimingService interface {
	CheckAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type timingService struct {
	c    client.Client
	name string
}

func NewTimingService(name string, c client.Client) TimingService {
	return &timingService{
		c:    c,
		name: name,
	}
}

func (c *timingService) CheckAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Timing.CheckAccounts", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Timing service

type TimingHandler interface {
	CheckAccounts(context.Context, *Request, *Response) error
}

func RegisterTimingHandler(s server.Server, hdlr TimingHandler, opts ...server.HandlerOption) error {
	type timing interface {
		CheckAccounts(ctx context.Context, in *Request, out *Response) error
	}
	type Timing struct {
		timing
	}
	h := &timingHandler{hdlr}
	return s.Handle(s.NewHandler(&Timing{h}, opts...))
}

type timingHandler struct {
	TimingHandler
}

func (h *timingHandler) CheckAccounts(ctx context.Context, in *Request, out *Response) error {
	return h.TimingHandler.CheckAccounts(ctx, in, out)
}