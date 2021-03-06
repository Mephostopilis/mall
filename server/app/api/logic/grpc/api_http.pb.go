// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0

package grpc

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type LogicHTTPServer interface {
	NodesInstances(context.Context, *NodesInstancesReq) (*NodesInstancesReply, error)
	NodesWeighted(context.Context, *NodesWeightedReq) (*NodesWeightedReply, error)
	OnlineRoom(context.Context, *OnlineRoomReq) (*OnlineRoomReply, error)
	OnlineTop(context.Context, *OnlineTopReq) (*OnlineTopReply, error)
	OnlineTotal(context.Context, *OnlineTotalReq) (*OnlineTotalReply, error)
	PushAll(context.Context, *PushAllReq) (*PushAllReply, error)
	PushKeys(context.Context, *PushKeysReq) (*PushKeysReply, error)
	PushMids(context.Context, *PushMidsReq) (*PushMidsReply, error)
	PushRoom(context.Context, *PushRoomReq) (*PushRoomReply, error)
}

func RegisterLogicHTTPServer(s *http.Server, srv LogicHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/goim/push/keys", _Logic_PushKeys0_HTTP_Handler(srv))
	r.POST("/api/v1/goim/push/mids", _Logic_PushMids0_HTTP_Handler(srv))
	r.POST("/api/v1/goim/push/room", _Logic_PushRoom0_HTTP_Handler(srv))
	r.POST("/api/v1/goim/push/all", _Logic_PushAll0_HTTP_Handler(srv))
	r.GET("/api/v1/goim/nodes/weighted", _Logic_NodesWeighted0_HTTP_Handler(srv))
	r.GET("/api/v1/goim/nodes/instances", _Logic_NodesInstances0_HTTP_Handler(srv))
	r.GET("/api/v1/goim/online/top", _Logic_OnlineTop0_HTTP_Handler(srv))
	r.GET("/api/v1/goim/online/room", _Logic_OnlineRoom0_HTTP_Handler(srv))
	r.GET("/api/v1/goim/online/total", _Logic_OnlineTotal0_HTTP_Handler(srv))
}

func _Logic_PushKeys0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PushKeysReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/PushKeys")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PushKeys(ctx, req.(*PushKeysReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PushKeysReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_PushMids0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PushMidsReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/PushMids")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PushMids(ctx, req.(*PushMidsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PushMidsReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_PushRoom0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PushRoomReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/PushRoom")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PushRoom(ctx, req.(*PushRoomReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PushRoomReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_PushAll0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PushAllReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/PushAll")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PushAll(ctx, req.(*PushAllReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PushAllReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_NodesWeighted0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NodesWeightedReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/NodesWeighted")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NodesWeighted(ctx, req.(*NodesWeightedReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NodesWeightedReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_NodesInstances0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NodesInstancesReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/NodesInstances")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NodesInstances(ctx, req.(*NodesInstancesReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NodesInstancesReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_OnlineTop0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OnlineTopReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/OnlineTop")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OnlineTop(ctx, req.(*OnlineTopReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OnlineTopReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_OnlineRoom0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OnlineRoomReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/OnlineRoom")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OnlineRoom(ctx, req.(*OnlineRoomReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OnlineRoomReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_OnlineTotal0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OnlineTotalReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.grpc.Logic/OnlineTotal")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OnlineTotal(ctx, req.(*OnlineTotalReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OnlineTotalReply)
		return ctx.Result(200, reply)
	}
}

type LogicHTTPClient interface {
	NodesInstances(ctx context.Context, req *NodesInstancesReq, opts ...http.CallOption) (rsp *NodesInstancesReply, err error)
	NodesWeighted(ctx context.Context, req *NodesWeightedReq, opts ...http.CallOption) (rsp *NodesWeightedReply, err error)
	OnlineRoom(ctx context.Context, req *OnlineRoomReq, opts ...http.CallOption) (rsp *OnlineRoomReply, err error)
	OnlineTop(ctx context.Context, req *OnlineTopReq, opts ...http.CallOption) (rsp *OnlineTopReply, err error)
	OnlineTotal(ctx context.Context, req *OnlineTotalReq, opts ...http.CallOption) (rsp *OnlineTotalReply, err error)
	PushAll(ctx context.Context, req *PushAllReq, opts ...http.CallOption) (rsp *PushAllReply, err error)
	PushKeys(ctx context.Context, req *PushKeysReq, opts ...http.CallOption) (rsp *PushKeysReply, err error)
	PushMids(ctx context.Context, req *PushMidsReq, opts ...http.CallOption) (rsp *PushMidsReply, err error)
	PushRoom(ctx context.Context, req *PushRoomReq, opts ...http.CallOption) (rsp *PushRoomReply, err error)
}

type LogicHTTPClientImpl struct {
	cc *http.Client
}

func NewLogicHTTPClient(client *http.Client) LogicHTTPClient {
	return &LogicHTTPClientImpl{client}
}

func (c *LogicHTTPClientImpl) NodesInstances(ctx context.Context, in *NodesInstancesReq, opts ...http.CallOption) (*NodesInstancesReply, error) {
	var out NodesInstancesReply
	pattern := "/api/v1/goim/nodes/instances"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/NodesInstances"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) NodesWeighted(ctx context.Context, in *NodesWeightedReq, opts ...http.CallOption) (*NodesWeightedReply, error) {
	var out NodesWeightedReply
	pattern := "/api/v1/goim/nodes/weighted"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/NodesWeighted"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) OnlineRoom(ctx context.Context, in *OnlineRoomReq, opts ...http.CallOption) (*OnlineRoomReply, error) {
	var out OnlineRoomReply
	pattern := "/api/v1/goim/online/room"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/OnlineRoom"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) OnlineTop(ctx context.Context, in *OnlineTopReq, opts ...http.CallOption) (*OnlineTopReply, error) {
	var out OnlineTopReply
	pattern := "/api/v1/goim/online/top"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/OnlineTop"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) OnlineTotal(ctx context.Context, in *OnlineTotalReq, opts ...http.CallOption) (*OnlineTotalReply, error) {
	var out OnlineTotalReply
	pattern := "/api/v1/goim/online/total"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/OnlineTotal"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) PushAll(ctx context.Context, in *PushAllReq, opts ...http.CallOption) (*PushAllReply, error) {
	var out PushAllReply
	pattern := "/api/v1/goim/push/all"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/PushAll"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) PushKeys(ctx context.Context, in *PushKeysReq, opts ...http.CallOption) (*PushKeysReply, error) {
	var out PushKeysReply
	pattern := "/api/v1/goim/push/keys"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/PushKeys"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) PushMids(ctx context.Context, in *PushMidsReq, opts ...http.CallOption) (*PushMidsReply, error) {
	var out PushMidsReply
	pattern := "/api/v1/goim/push/mids"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/PushMids"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) PushRoom(ctx context.Context, in *PushRoomReq, opts ...http.CallOption) (*PushRoomReply, error) {
	var out PushRoomReply
	pattern := "/api/v1/goim/push/room"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.logic.grpc.Logic/PushRoom"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
