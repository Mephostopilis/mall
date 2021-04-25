package biz

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	pb "edu/api/sso"
	"edu/pkg/ecode"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
)

// GetRedirectURI get redirect uri
func (s *GreeterUsecase) GetRedirectURI(redirectURI string, state string, responseType oauth2.ResponseType, data map[string]interface{}) (string, error) {
	u, err := url.Parse(redirectURI)
	if err != nil {
		return "", err
	}
	q := u.Query()
	if state != "" {
		q.Set("state", state)
	}
	for k, v := range data {
		q.Set(k, fmt.Sprint(v))
	}

	switch responseType {
	case oauth2.Code:
		u.RawQuery = q.Encode()
	case oauth2.Token:
		u.RawQuery = ""
		fragment, err := url.QueryUnescape(q.Encode())
		if err != nil {
			return "", err
		}
		u.Fragment = fragment
	}
	return u.String(), nil
}

// CheckResponseType check allows response type
func (s *GreeterUsecase) checkResponseType(rt oauth2.ResponseType) bool {
	for _, art := range s.oc.AllowedResponseTypes {
		if art == rt {
			return true
		}
	}
	return false
}

// ValidationAuthorizeRequest the authorization request validation
func (s *GreeterUsecase) ValidationAuthorizeRequest(ctx context.Context, req *pb.AuthorizeReq) (err error) {
	resType := oauth2.ResponseType(req.ResponseType)
	if resType.String() == "" {
		err = ecode.ErrUnsupportedResponseType
		return
	} else if allowed := s.checkResponseType(resType); !allowed {
		err = ecode.ErrUnauthorizedClient
		return
	}
	return
}

// GetAuthorizeToken get authorization token(code)
func (s *GreeterUsecase) GetAuthorizeToken(ctx context.Context, req *pb.AuthorizeReq, userID string, exp time.Duration) (oauth2.TokenInfo, error) {
	// check the client allows the grant type
	gt := oauth2.AuthorizationCode
	if oauth2.ResponseType(req.ResponseType) == oauth2.Token {
		gt = oauth2.Implicit
	}

	// 用户登陆
	allowed, err := s.onClientAuthorizedHandler(ctx, req.ClientID, gt)
	if err != nil {
		return nil, err
	} else if !allowed {
		return nil, ecode.ErrUnauthorizedClient
	}

	// check the client allows the authorized scope
	allowed, err = s.onClientScopeHandler(ctx, req.ClientID, req.Scope)
	if err != nil {
		return nil, err
	} else if !allowed {
		return nil, ecode.ErrInvalidScope
	}

	tgr := &oauth2.TokenGenerateRequest{
		ClientID:       req.ClientID,
		UserID:         userID,
		RedirectURI:    req.RedirectURI,
		Scope:          req.Scope,
		AccessTokenExp: exp,
	}
	return s.manager.GenerateAuthToken(ctx, oauth2.ResponseType(req.ResponseType), tgr)
}

// GetAuthorizeData get authorization response data
func (s *GreeterUsecase) GetAuthorizeData(ctx context.Context, state string, ti oauth2.TokenInfo) (resp *pb.AuthorizeResp, err error) {
	resp = &pb.AuthorizeResp{
		Location: ti.GetRedirectURI(),
		Code:     ti.GetCode(),
		State:    state,
	}
	return
}

// HandleAuthorizeRequest the authorization request handling
func (s *GreeterUsecase) HandleAuthorizeRequest(ctx context.Context, req *pb.AuthorizeReq) (resp *pb.AuthorizeResp, err error) {
	if err = s.ValidationAuthorizeRequest(ctx, req); err != nil {
		return
	}

	// user authorization
	userID, err := s.onUserAuthorizationHandler(ctx, req)
	if err != nil {
		err = ecode.ErrInvalidUserId
		return
	} else if userID == "" {
		err = ecode.ErrInvalidUserId
		return
	}

	// specify the scope of authorization
	scope, err := s.onAuthorizeScopeHandler(ctx, req)
	if err != nil {
		return
	} else if scope != "" {
		req.Scope = scope
	}

	// specify the expiration time of access token
	exp, err := s.onAccessTokenExpHandler(ctx, req)
	if err != nil {
		return
	}

	ti, err := s.GetAuthorizeToken(ctx, req, userID, exp)
	if err != nil {
		return
	}

	// If the redirect URI is empty, the default domain provided by the client is used.
	if req.RedirectURI == "" {
		client, e := s.manager.GetClient(ctx, req.ClientID)
		if e != nil {
			err = e
			return
		}
		req.RedirectURI = client.GetDomain()
	}
	resp, err = s.GetAuthorizeData(ctx, req.State, ti)
	return
}

// ValidationTokenRequest the token request validation
func (s *GreeterUsecase) ValidationTokenRequest(ctx context.Context, req *pb.TokenReq) (gt oauth2.GrantType, tgr *oauth2.TokenGenerateRequest, err error) {
	rt := oauth2.ResponseType(req.ResponseType)
	if rt != oauth2.Token {
		err = ecode.ErrUnsupportedResponseType
		return
	}
	gt = oauth2.GrantType(req.GrantType)
	if gt.String() == "" {
		err = ecode.ErrUnsupportedGrantType
		return
	}
	clientID := req.ClientID
	clientSecret := req.ClientSecret

	tgr = &oauth2.TokenGenerateRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	switch gt {
	case oauth2.AuthorizationCode:
		tgr.RedirectURI = req.RedirectURI
		tgr.Code = req.Code
		if tgr.RedirectURI == "" ||
			tgr.Code == "" {
			err = ecode.ErrInvalidRequest
			return
		}
	case oauth2.PasswordCredentials:
		tgr.Scope = req.Scope
		username, password := req.Username, req.Password
		if username == "" || password == "" {
			return "", nil, ecode.ErrInvalidUsername
		}

		// 针对password模式检测scope与grant_type
		// 用户登陆
		allowed, xerr := s.onClientAuthorizedHandler(ctx, req.ClientID, gt)
		if xerr != nil {
			err = xerr
			return
		} else if !allowed {
			return
		}

		// check the client allows the authorized scope
		allowed, err = s.onClientScopeHandler(ctx, req.ClientID, req.Scope)
		if err != nil {
			return
		} else if !allowed {
			err = ecode.ErrInvalidScope
			return
		}

		userID, xerr := s.onPasswordAuthorizationHandler(ctx, username, password)
		if err != nil {
			err = xerr
			return
		} else if userID == "" {
			err = ecode.ErrInvalidUid
			return
		}
		tgr.UserID = userID
	case oauth2.ClientCredentials:
		tgr.Scope = req.Scope
	case oauth2.Refreshing:
		if req.RefreshToken == "" {
			err = ecode.ErrInvalidRequest
			return
		}
		tgr.Refresh = req.RefreshToken
		tgr.Scope = req.Scope
	}
	return
}

// CheckGrantType check allows grant type
func (s *GreeterUsecase) checkGrantType(gt oauth2.GrantType) bool {
	for _, agt := range s.oc.AllowedGrantTypes {
		if agt == gt {
			return true
		}
	}
	return false
}

// GetAccessToken access token
func (s *GreeterUsecase) GetAccessToken(ctx context.Context, gt oauth2.GrantType, tgr *oauth2.TokenGenerateRequest) (ti oauth2.TokenInfo, err error) {
	if allowed := s.checkGrantType(gt); !allowed {
		s.log.Errorf("gt = %s", gt)
		err = ecode.ErrUnauthorizedClient
		return
	}
	allowed, err := s.onClientAuthorizedHandler(ctx, tgr.ClientID, gt)
	if err != nil {
		return
	} else if !allowed {
		err = ecode.ErrUnauthorizedClient
		return
	}

	switch gt {
	case oauth2.AuthorizationCode:
		ti, err = s.manager.GenerateAccessToken(ctx, gt, tgr)
		if err != nil {
			switch err {
			case errors.ErrInvalidAuthorizeCode:
				err = ecode.ErrInvalidAuthorizeCode
				return
			case errors.ErrInvalidClient:
				err = ecode.ErrInvalidClient
				return
			default:
				err = ecode.ErrUnsupportedErrType
				return
			}
		}
		return
	case oauth2.PasswordCredentials, oauth2.ClientCredentials:
		allowed, err = s.onClientScopeHandler(ctx, tgr.ClientID, tgr.Scope)
		if err != nil {
			return
		} else if !allowed {
			err = ecode.ErrInvalidScope
			return
		}
		return s.manager.GenerateAccessToken(ctx, gt, tgr)
	case oauth2.Refreshing:
		// check scope
		if scope := tgr.Scope; scope != "" {
			rti, e := s.manager.LoadRefreshToken(ctx, tgr.Refresh)
			if e != nil {
				if e == errors.ErrInvalidRefreshToken || e == errors.ErrExpiredRefreshToken {
					err = ecode.ErrInvalidGrant
					return
				}
				err = ecode.ErrUnsupportedErrType
				return
			}

			allowed, err := s.onRefreshingScopeHandler(scope, rti.GetScope())
			if err != nil {
				return nil, err
			} else if !allowed {
				return nil, ecode.ErrInvalidScope
			}
		}

		ti, err := s.manager.RefreshAccessToken(ctx, tgr)
		if err != nil {
			if err == errors.ErrInvalidRefreshToken || err == errors.ErrExpiredRefreshToken {
				return nil, errors.ErrInvalidGrant
			}
			return nil, err
		}
		return ti, nil
	}
	return nil, ecode.ErrUnsupportedGrantType
}

// GetTokenData token data
func (s *GreeterUsecase) GetTokenData(ctx context.Context, ti oauth2.TokenInfo) (resp *pb.TokenResp) {
	d := &pb.TokenResp{
		AccessToken: ti.GetAccess(),
		TokenType:   s.oc.TokenType,
		ExpiresIn:   fmt.Sprintf("%d", int64(ti.GetAccessExpiresIn()/time.Second)),
	}
	if scope := ti.GetScope(); scope != "" {
		d.Scope = scope
	}
	if refresh := ti.GetRefresh(); refresh != "" {
		d.RefreshToken = refresh
	}
	if openid := ti.GetUserID(); openid != "" {
		d.OpenID = openid
	}
	return d
}

// HandleTokenRequest token request handling
// @Summary 生成ID
// @Description 健康状况
// @Accept text/html
// @Produce text/html
// @Success 200 {string} string "OK"
// @Router /sso/public/health [get]
// @BasePath
func (s *GreeterUsecase) HandleTokenRequest(ctx context.Context, req *pb.TokenReq) (resp *pb.TokenResp, err error) {
	gt, tgr, err := s.ValidationTokenRequest(ctx, req)
	if err != nil {
		return
	}
	ti, err := s.GetAccessToken(ctx, gt, tgr)
	if err != nil {
		return
	}
	resp = s.GetTokenData(ctx, ti)
	return
}

func (s *GreeterUsecase) ValidationMidToken(ctx context.Context, req *pb.IntrospectReq) (resp *pb.IntrospectResp, err error) {
	ti, err := s.manager.LoadAccessToken(ctx, req.AccessToken)
	if err != nil {

		err = ecode.ErrInvalidAccessToken
		return
	}
	id, err := strconv.Atoi(ti.GetUserID())
	if err != nil {
		err = ecode.UserNotExist
		return
	}
	mid := uint64(id)
	s.log.Infof("mid = %v", mid)
	resp = &pb.IntrospectResp{
		UserId: mid,
	}
	return
}

// GetErrorData get error response data
// func (s *Service) GetErrorData(ctx context.Context, err error) (map[string]interface{}, int) {
// 	var re errors.Response
// 	if v, ok := errors.Descriptions[err]; ok {
// 		re.Error = err
// 		re.Description = v
// 		re.StatusCode = errors.StatusCodes[err]
// 	} else {

// 		if v := s.onInternalErrorHandler(err); v != nil {
// 			re = *v
// 		}
// 		if re.Error == nil {
// 			re.Error = errors.ErrServerError
// 			re.Description = errors.Descriptions[errors.ErrServerError]
// 			re.StatusCode = errors.StatusCodes[errors.ErrServerError]
// 		}
// 	}

// 	s.onResponseErrorHandler(&re)

// 	data := make(map[string]interface{})
// 	if err := re.Error; err != nil {
// 		data["error"] = err.Error()
// 	}

// 	if v := re.ErrorCode; v != 0 {
// 		data["error_code"] = v
// 	}

// 	if v := re.Description; v != "" {
// 		data["error_description"] = v
// 	}

// 	if v := re.URI; v != "" {
// 		data["error_uri"] = v
// 	}

// 	// statusCode := http.StatusInternalServerError
// 	// if v := re.StatusCode; v > 0 {
// 	// 	statusCode = v
// 	// }

// 	return data, statusCode
// }
