// Code generated by protoc-gen-go-http. DO NOT EDIT.

package sms

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type AdminHandler interface {
	CreateCoupon(context.Context, *Coupon) (*ApiReply, error)

	CreateCouponHistory(context.Context, *CouponHistory) (*ApiReply, error)

	DeleteCoupon(context.Context, *DeleteCouponRequest) (*ApiReply, error)

	DeleteCouponHistory(context.Context, *DeleteCouponHistoryRequest) (*ApiReply, error)

	GetCoupon(context.Context, *GetCouponRequest) (*ApiReply, error)

	GetCouponHistory(context.Context, *GetCouponHistoryRequest) (*ApiReply, error)

	ListCoupon(context.Context, *ListCouponRequest) (*ApiReply, error)

	ListCouponHistory(context.Context, *ListCouponHistoryRequest) (*ApiReply, error)

	UpdateCoupon(context.Context, *Coupon) (*ApiReply, error)

	UpdateCouponHistory(context.Context, *CouponHistory) (*ApiReply, error)
}

func NewAdminHandler(srv AdminHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/admin/v1/smscouponList", func(w http.ResponseWriter, r *http.Request) {
		var in ListCouponRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCoupon(ctx, req.(*ListCouponRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/admin/v1/smscoupon/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in GetCouponRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCoupon(ctx, req.(*GetCouponRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/admin/v1/smscoupon", func(w http.ResponseWriter, r *http.Request) {
		var in Coupon
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCoupon(ctx, req.(*Coupon))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/admin/v1/smscoupon", func(w http.ResponseWriter, r *http.Request) {
		var in Coupon
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCoupon(ctx, req.(*Coupon))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("PUT")

	r.HandleFunc("/admin/v1/smscoupon/{ids}", func(w http.ResponseWriter, r *http.Request) {
		var in DeleteCouponRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCoupon(ctx, req.(*DeleteCouponRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("DELETE")

	r.HandleFunc("/admin/v1/smscouponhistoryList", func(w http.ResponseWriter, r *http.Request) {
		var in ListCouponHistoryRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCouponHistory(ctx, req.(*ListCouponHistoryRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/admin/v1/smscouponhistory/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in GetCouponHistoryRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCouponHistory(ctx, req.(*GetCouponHistoryRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/admin/v1/smscouponhistory", func(w http.ResponseWriter, r *http.Request) {
		var in CouponHistory
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCouponHistory(ctx, req.(*CouponHistory))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/admin/v1/smscouponhistory", func(w http.ResponseWriter, r *http.Request) {
		var in CouponHistory
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCouponHistory(ctx, req.(*CouponHistory))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("PUT")

	r.HandleFunc("/admin/v1/smscouponhistory/{ids}", func(w http.ResponseWriter, r *http.Request) {
		var in DeleteCouponHistoryRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCouponHistory(ctx, req.(*DeleteCouponHistoryRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ApiReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("DELETE")

	return r
}

type ApiHandler interface {
	GetCoupon(context.Context, *GetCouponRequest) (*GetCouponReply, error)

	ListCoupon(context.Context, *ListCouponRequest) (*ListCouponReply, error)

	SayHelloURL(context.Context, *HelloReq) (*HelloResp, error)
}

func NewApiHandler(srv ApiHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/sms/say_hello", func(w http.ResponseWriter, r *http.Request) {
		var in HelloReq
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHelloURL(ctx, req.(*HelloReq))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*HelloResp)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/api/v1/sms/couponList", func(w http.ResponseWriter, r *http.Request) {
		var in ListCouponRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCoupon(ctx, req.(*ListCouponRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ListCouponReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/api/v1/sms/coupon/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in GetCouponRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCoupon(ctx, req.(*GetCouponRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetCouponReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	return r
}
