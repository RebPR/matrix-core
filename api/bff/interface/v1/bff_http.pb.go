// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type BffHTTPServer interface {
	GetAccount(context.Context, *emptypb.Empty) (*GetAccountReply, error)
	GetCosSessionKey(context.Context, *emptypb.Empty) (*GetCosSessionKeyReply, error)
	GetProfile(context.Context, *emptypb.Empty) (*GetProfileReply, error)
	GetProfileUpdate(context.Context, *emptypb.Empty) (*GetProfileUpdateReply, error)
	LoginByCode(context.Context, *LoginByCodeReq) (*LoginReply, error)
	LoginByGithub(context.Context, *LoginByGithubReq) (*LoginReply, error)
	LoginByPassword(context.Context, *LoginByPasswordReq) (*LoginReply, error)
	LoginByWeChat(context.Context, *LoginByWeChatReq) (*LoginReply, error)
	LoginPasswordReset(context.Context, *LoginPasswordResetReq) (*emptypb.Empty, error)
	SendEmailCode(context.Context, *SendEmailCodeReq) (*emptypb.Empty, error)
	SendPhoneCode(context.Context, *SendPhoneCodeReq) (*emptypb.Empty, error)
	SetProfileUpdate(context.Context, *SetProfileUpdateReq) (*emptypb.Empty, error)
	SetUserEmail(context.Context, *SetUserEmailReq) (*emptypb.Empty, error)
	SetUserPhone(context.Context, *SetUserPhoneReq) (*emptypb.Empty, error)
	UserRegister(context.Context, *UserRegisterReq) (*emptypb.Empty, error)
}

func RegisterBffHTTPServer(s *http.Server, srv BffHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user/register", _Bff_UserRegister0_HTTP_Handler(srv))
	r.POST("/v1/user/login/password", _Bff_LoginByPassword0_HTTP_Handler(srv))
	r.POST("/v1/user/login/code", _Bff_LoginByCode0_HTTP_Handler(srv))
	r.POST("/v1/user/login/wechat", _Bff_LoginByWeChat0_HTTP_Handler(srv))
	r.POST("/v1/user/login/github", _Bff_LoginByGithub0_HTTP_Handler(srv))
	r.POST("/v1/user/login/password/reset", _Bff_LoginPasswordReset0_HTTP_Handler(srv))
	r.POST("/v1/user/code/phone", _Bff_SendPhoneCode0_HTTP_Handler(srv))
	r.POST("/v1/user/code/email", _Bff_SendEmailCode0_HTTP_Handler(srv))
	r.GET("/v1/get/cos/session/key", _Bff_GetCosSessionKey0_HTTP_Handler(srv))
	r.GET("/v1/get/user/account", _Bff_GetAccount0_HTTP_Handler(srv))
	r.GET("/v1/get/user/profile", _Bff_GetProfile0_HTTP_Handler(srv))
	r.GET("/v1/get/user/profile/update", _Bff_GetProfileUpdate0_HTTP_Handler(srv))
	r.POST("/v1/set/user/profile/update", _Bff_SetProfileUpdate0_HTTP_Handler(srv))
	r.POST("/v1/set/user/phone", _Bff_SetUserPhone0_HTTP_Handler(srv))
	r.POST("/v1/set/user/email", _Bff_SetUserEmail0_HTTP_Handler(srv))
}

func _Bff_UserRegister0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/UserRegister")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserRegister(ctx, req.(*UserRegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Bff_LoginByPassword0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginByPasswordReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/LoginByPassword")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByPassword(ctx, req.(*LoginByPasswordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Bff_LoginByCode0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginByCodeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/LoginByCode")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByCode(ctx, req.(*LoginByCodeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Bff_LoginByWeChat0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginByWeChatReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/LoginByWeChat")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByWeChat(ctx, req.(*LoginByWeChatReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Bff_LoginByGithub0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginByGithubReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/LoginByGithub")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByGithub(ctx, req.(*LoginByGithubReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Bff_LoginPasswordReset0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginPasswordResetReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/LoginPasswordReset")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginPasswordReset(ctx, req.(*LoginPasswordResetReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Bff_SendPhoneCode0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendPhoneCodeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/SendPhoneCode")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendPhoneCode(ctx, req.(*SendPhoneCodeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Bff_SendEmailCode0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendEmailCodeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/SendEmailCode")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendEmailCode(ctx, req.(*SendEmailCodeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Bff_GetCosSessionKey0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/GetCosSessionKey")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCosSessionKey(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCosSessionKeyReply)
		return ctx.Result(200, reply)
	}
}

func _Bff_GetAccount0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/GetAccount")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAccount(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAccountReply)
		return ctx.Result(200, reply)
	}
}

func _Bff_GetProfile0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/GetProfile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Bff_GetProfileUpdate0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/GetProfileUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfileUpdate(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProfileUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _Bff_SetProfileUpdate0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetProfileUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/SetProfileUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetProfileUpdate(ctx, req.(*SetProfileUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Bff_SetUserPhone0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserPhoneReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/SetUserPhone")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserPhone(ctx, req.(*SetUserPhoneReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Bff_SetUserEmail0_HTTP_Handler(srv BffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserEmailReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bff.v1.Bff/SetUserEmail")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserEmail(ctx, req.(*SetUserEmailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type BffHTTPClient interface {
	GetAccount(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetAccountReply, err error)
	GetCosSessionKey(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetCosSessionKeyReply, err error)
	GetProfile(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetProfileReply, err error)
	GetProfileUpdate(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetProfileUpdateReply, err error)
	LoginByCode(ctx context.Context, req *LoginByCodeReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginByGithub(ctx context.Context, req *LoginByGithubReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginByPassword(ctx context.Context, req *LoginByPasswordReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginByWeChat(ctx context.Context, req *LoginByWeChatReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginPasswordReset(ctx context.Context, req *LoginPasswordResetReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	SendEmailCode(ctx context.Context, req *SendEmailCodeReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	SendPhoneCode(ctx context.Context, req *SendPhoneCodeReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	SetProfileUpdate(ctx context.Context, req *SetProfileUpdateReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	SetUserEmail(ctx context.Context, req *SetUserEmailReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	SetUserPhone(ctx context.Context, req *SetUserPhoneReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UserRegister(ctx context.Context, req *UserRegisterReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type BffHTTPClientImpl struct {
	cc *http.Client
}

func NewBffHTTPClient(client *http.Client) BffHTTPClient {
	return &BffHTTPClientImpl{client}
}

func (c *BffHTTPClientImpl) GetAccount(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetAccountReply, error) {
	var out GetAccountReply
	pattern := "/v1/get/user/account"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bff.v1.Bff/GetAccount"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) GetCosSessionKey(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetCosSessionKeyReply, error) {
	var out GetCosSessionKeyReply
	pattern := "/v1/get/cos/session/key"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bff.v1.Bff/GetCosSessionKey"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) GetProfile(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetProfileReply, error) {
	var out GetProfileReply
	pattern := "/v1/get/user/profile"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bff.v1.Bff/GetProfile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) GetProfileUpdate(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetProfileUpdateReply, error) {
	var out GetProfileUpdateReply
	pattern := "/v1/get/user/profile/update"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bff.v1.Bff/GetProfileUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) LoginByCode(ctx context.Context, in *LoginByCodeReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login/code"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/LoginByCode"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) LoginByGithub(ctx context.Context, in *LoginByGithubReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login/github"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/LoginByGithub"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) LoginByPassword(ctx context.Context, in *LoginByPasswordReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/LoginByPassword"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) LoginByWeChat(ctx context.Context, in *LoginByWeChatReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login/wechat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/LoginByWeChat"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) LoginPasswordReset(ctx context.Context, in *LoginPasswordResetReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/user/login/password/reset"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/LoginPasswordReset"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) SendEmailCode(ctx context.Context, in *SendEmailCodeReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/user/code/email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/SendEmailCode"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) SendPhoneCode(ctx context.Context, in *SendPhoneCodeReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/user/code/phone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/SendPhoneCode"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) SetProfileUpdate(ctx context.Context, in *SetProfileUpdateReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/set/user/profile/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/SetProfileUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) SetUserEmail(ctx context.Context, in *SetUserEmailReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/set/user/email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/SetUserEmail"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) SetUserPhone(ctx context.Context, in *SetUserPhoneReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/set/user/phone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/SetUserPhone"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BffHTTPClientImpl) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bff.v1.Bff/UserRegister"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
