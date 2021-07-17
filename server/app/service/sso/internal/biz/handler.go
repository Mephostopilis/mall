package biz

import (
	"context"
	"fmt"
	"time"

	pb "edu/api/sso/v1"
	"edu/service/sso/internal/model"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
)

// func (s *Service) onClientInfoHandler(r *http.Request) (clientID, clientSecret string, err error) {
// 	clientID := r.Form.Get("client_id")
// 	clientSecret := r.Form.Get("client_secret")
// 	if clientID == "" || clientSecret == "" {
// 		return "", "", errors.ErrInvalidClient
// 	}
// 	return clientID, clientSecret, nil
// 	return
// }

// authorized 模式下用
// OnClientAuthorizedHandler check the client allows to use this authorization grant type
func (s *GreeterUsecase) onClientAuthorizedHandler(ctx context.Context, clientID string, gt oauth2.GrantType) (allowed bool, err error) {
	m, err := s.repo.GetSsoApp(ctx, &model.SsoApp{Id: clientID})
	if err != nil {
		return
	}
	x := oauth2.GrantType(m.GrantType)
	if gt == x {
		allowed = true
		return
	}
	allowed = false
	return
}

// ClientScopeHandler check the client allows to use scope
func (s *GreeterUsecase) onClientScopeHandler(ctx context.Context, clientID, scope string) (allowed bool, err error) {
	m, err := s.repo.GetSsoApp(ctx, &model.SsoApp{Id: clientID})
	if err != nil {
		return
	}
	if m.Scope == scope {
		allowed = true
	} else {
		allowed = false
	}
	return
}

// UserAuthorizationHandler get user id from request authorization
func (s *GreeterUsecase) onUserAuthorizationHandler(ctx context.Context, req *pb.AuthorizeReq) (userID string, err error) {
	// username := req.
	// password := r.FormValue("password")

	return
}

// AuthorizeScopeHandler set the authorized scope
func (s *GreeterUsecase) onAuthorizeScopeHandler(ctx context.Context, req *pb.AuthorizeReq) (scope string, err error) {
	scope = req.Scope
	return
}

// AccessTokenExpHandler set expiration date for the access token
func (s *GreeterUsecase) onAccessTokenExpHandler(ctx context.Context, req *pb.AuthorizeReq) (exp time.Duration, err error) {
	return
}

// password 模式下用
// PasswordAuthorizationHandler get user id from username and password
func (s *GreeterUsecase) onPasswordAuthorizationHandler(ctx context.Context, username, password string) (userID string, err error) {
	u, err := s.repo.GetSysUserInfo(&model.SysUser{
		LoginM: model.LoginM{
			UserName: model.UserName{Username: username},
			PassWord: model.PassWord{Password: password}}})
	if err != nil {
		return
	}
	return fmt.Sprint(u.UserId), nil
}

// RefreshingScopeHandler check the scope of the refreshing token
func (s *GreeterUsecase) onRefreshingScopeHandler(newScope, oldScope string) (allowed bool, err error) {
	return
}

// ResponseErrorHandler response error handing
func (s *GreeterUsecase) onResponseErrorHandler(re *errors.Response) {
	return
}

// InternalErrorHandler internal error handing
func (s *GreeterUsecase) onInternalErrorHandler(err error) (re *errors.Response) {
	return
}

// ExtensionFieldsHandler in response to the access token with the extension of the field
func (s *GreeterUsecase) onExtensionFieldsHandler(ti oauth2.TokenInfo) (fieldsValue map[string]interface{}) {
	return
}
