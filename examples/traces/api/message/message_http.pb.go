// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	middleware "github.com/go-kratos/kratos/v2/middleware"
	transport "github.com/go-kratos/kratos/v2/transport"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = new(middleware.Middleware)
var _ = new(transport.Transporter)
var _ = binding.BindVars
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type MessageServiceHandler interface {
	GetUserMessage(context.Context, *GetUserMessageRequest) (*GetUserMessageReply, error)
}

func NewMessageServiceHandler(srv MessageServiceHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/message/user/{id}/{count}", func(w http.ResponseWriter, r *http.Request) {
		var in GetUserMessageRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserMessage(ctx, req.(*GetUserMessageRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		ctx := r.Context()
		transport.SetMethod(ctx, "/api.message.v1.MessageService/GetUserMessage")
		out, err := next(ctx, &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetUserMessageReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	return r
}

type MessageServiceHTTPClient interface {
	GetUserMessage(ctx context.Context, req *GetUserMessageRequest, opts ...http1.CallOption) (rsp *GetUserMessageReply, err error)
}

type MessageServiceHTTPClientImpl struct {
	cc *http1.Client
}

func NewMessageServiceHTTPClient(client *http1.Client) MessageServiceHTTPClient {
	return &MessageServiceHTTPClientImpl{client}
}

func (c *MessageServiceHTTPClientImpl) GetUserMessage(ctx context.Context, in *GetUserMessageRequest, opts ...http1.CallOption) (*GetUserMessageReply, error) {
	var out GetUserMessageReply
	path := binding.EncodePath("GET", "/v1/message/user/{id}/{count}", in)
	opts = append(opts, http1.Method("/api.message.v1.MessageService/GetUserMessage"))

	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)

	return &out, err
}