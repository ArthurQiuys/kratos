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

type UserHandler interface {
	GetMyMessages(context.Context, *GetMyMessagesRequest) (*GetMyMessagesReply, error)
}

func NewUserHandler(srv UserHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/user/get/message/{count}", func(w http.ResponseWriter, r *http.Request) {
		var in GetMyMessagesRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyMessages(ctx, req.(*GetMyMessagesRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		ctx := r.Context()
		transport.SetMethod(ctx, "/api.user.v1.User/GetMyMessages")
		out, err := next(ctx, &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetMyMessagesReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	return r
}

type UserHTTPClient interface {
	GetMyMessages(ctx context.Context, req *GetMyMessagesRequest, opts ...http1.CallOption) (rsp *GetMyMessagesReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http1.Client
}

func NewUserHTTPClient(client *http1.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) GetMyMessages(ctx context.Context, in *GetMyMessagesRequest, opts ...http1.CallOption) (*GetMyMessagesReply, error) {
	var out GetMyMessagesReply
	path := binding.EncodePath("GET", "/v1/user/get/message/{count}", in)
	opts = append(opts, http1.Method("/api.user.v1.User/GetMyMessages"))

	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)

	return &out, err
}