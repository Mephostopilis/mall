package biz

import (
	"context"
	"crypto/rsa"
	"time"

	pb "edu/api/sso/v1"
	"edu/pkg/meta"
	"edu/pkg/net/ip"
	"edu/pkg/tools"
	"edu/service/sso/internal/conf"
	"edu/service/sso/internal/dao"
	"edu/service/sso/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
)

var (
	captchaStore = base64Captcha.DefaultMemStore
)

// GinJWTMiddleware provides a Json-Web-Token authentication implementation. On failure, a 401 HTTP response
// is returned. On success, the wrapped middleware is called, and the userID is made available as
// c.Get("userID").(string).
// Users can get a token by posting a json request to LoginHandler. The token then needs to be passed in
// the Authentication header. Example: Authorization:Bearer XXX_TOKEN_XXX
type JWTUsecase struct {
	// Realm name to display to the user. Required.
	Realm string

	// signing algorithm - possible values are HS256, HS384, HS512
	// Optional, default is HS256.
	SigningAlgorithm string

	// Secret key used for signing. Required.
	Key []byte

	// Duration that a jwt token is valid. Optional, defaults to one hour.
	Timeout time.Duration

	// This field allows clients to refresh their token until MaxRefresh has passed.
	// Note that clients can refresh their token in the last moment of MaxRefresh.
	// This means that the maximum validity timespan for a token is TokenTime + MaxRefresh.
	// Optional, defaults to 0 meaning not refreshable.
	MaxRefresh time.Duration

	// Callback function that will be called during login.
	// Using this function it is possible to add additional payload data to the webtoken.
	// The data is then made available during requests via c.Get("JWT_PAYLOAD").
	// Note that the payload is not encrypted.
	// The attributes mentioned on jwt.io can't be used as keys for the map.
	// Optional, by default no additional data will be set.
	PayloadFunc PayloadHandler

	// Set the identity handler function
	IdentityFunc IdentityHandler

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc func() time.Time

	// Private key file for asymmetric algorithms
	PrivKeyFile string

	// Public key file for asymmetric algorithms
	PubKeyFile string

	// Private key
	privKey *rsa.PrivateKey

	// Public key
	pubKey *rsa.PublicKey

	// log
	log *log.Helper

	// dao
	dao dao.Dao
}

// New 分配
func NewJWTUsecase(c *conf.Jwt, logger log.Logger, d dao.Dao) (*JWTUsecase, error) {
	log := log.NewHelper(log.With(logger, "module", "usecase/jwtauth"))
	du := c.Timeout.AsDuration()
	// mw, err := jwtauth.New(
	// 	&jwtauth.Jwt{Secret: c.Secret, Timeout: du},
	// 	logger,
	// 	jwtauth.DataPermissionToClaimsFunc,
	// 	jwtauth.ClaimsToDataPermissionFunc)
	// if err != nil {
	// 	return nil, err
	// }
	uc := &JWTUsecase{
		Realm:            "test zone",
		SigningAlgorithm: "HS256",
		Key:              []byte(c.Secret),
		Timeout:          du,
		MaxRefresh:       time.Hour,
		PayloadFunc:      DataPermissionToClaimsFunc,
		IdentityFunc:     ClaimsToDataPermissionFunc,
		TimeFunc:         time.Now,

		log: log,
		dao: d,
	}
	return uc, nil
}

// MiddlewareFunc makes GinJWTMiddleware implement the Middleware interface.
func (mw *JWTUsecase) ValidationMidToken(ctx context.Context, token string) (pd *meta.DataPermission, err error) {
	out, err := mw.ValidationToken(token)
	if err != nil {
		return
	}
	pd = out.(*meta.DataPermission)
	return
}

// @Summary 登陆
// @Description 获取token
// @Description LoginHandler can be used by clients to get a jwt token.
// @Description Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// @Description Reply will be of the form {"token": "TOKEN"}.
// @Description dev mode：It should be noted that all fields cannot be empty, and a value of 0 can be passed in addition to the account password
// @Description 注意：开发模式：需要注意全部字段不能为空，账号密码外可以传入0值
// @Accept  application/json
// @Product application/json
// @Param account body pb.LoginRequest  true "account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /login [post]
func (mw *JWTUsecase) authenticator(ctx context.Context, req *pb.LoginRequest) (*model.SysUser, *model.SysRole, error) {
	var loginVals model.Login
	var status = "0"
	var msg = "登录成功"
	var username = req.Username

	loginVals.Username = req.Username
	loginVals.Password = req.Password
	loginVals.Code = req.Code
	loginVals.UUID = req.Uuid
	// 验证参数
	if loginVals.Username == "" {
		// 数据解析失败
		mw.log.Errorf("username is null")
		return nil, nil, pb.ErrMissingLoginValues("")
	}
	// 验证验证吗
	if !captchaStore.Verify(loginVals.UUID, loginVals.Code, true) {
		msg = "验证码错误"
		status = "1"
		mw.saveLoginLogToDB(ctx, status, msg, username)
		return nil, nil, pb.ErrInvalidVerificationode("")
	}
	// 验证
	user, role, err := mw.dao.GetJwtAuthUser(ctx, &loginVals)
	if err != nil {
		msg = "登录失败"
		status = "1"
		mw.saveLoginLogToDB(ctx, status, msg, username)
		mw.log.Errorf("err = %v", err)
		return nil, nil, err
	}

	mw.saveLoginLogToDB(ctx, status, msg, username)
	return &user, &role, nil
}

// Write log to database
func (mw *JWTUsecase) saveLoginLogToDB(ctx context.Context, status string, msg string, username string) (err error) {
	uastr, _ := meta.GetUA(ctx)
	remoteAddr, _ := meta.GetRemoteAddr(ctx)

	var loginlog model.SysLoginLog
	ua := user_agent.New(uastr)
	loginlog.Ipaddr = remoteAddr
	loginlog.Username = username
	location := ip.GetLocation(remoteAddr)
	loginlog.LoginLocation = location
	loginlog.LoginTime = tools.GetCurrentTime()
	loginlog.Status = status
	loginlog.Remark = uastr
	browserName, browserVersion := ua.Browser()
	loginlog.Browser = browserName + " " + browserVersion
	loginlog.Os = ua.OS()
	loginlog.Msg = msg
	loginlog.Platform = ua.Platform()
	_, err = mw.dao.CreateJwtAuthLoginLog(ctx, &loginlog)
	return
}

// LoginHandler can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
// 登录
func (mw *JWTUsecase) LoginHandler(ctx context.Context, req *pb.LoginRequest) (tokenString string, expire time.Time, err error) {
	u, r, err := mw.authenticator(ctx, req)
	if err != nil {
		return
	}

	dp := &meta.DataPermission{
		UserId:  u.UserId,
		RoleId:  uint64(r.RoleId),
		RoleKey: r.RoleKey,
	}

	// Create the token
	tokenString, expire, err = mw.TokenGenerator(dp)
	return
}

// RefreshHandler can be used to refresh a token. The token still needs to be valid on refresh.
// Shall be put under an endpoint that is using the GinJWTMiddleware.
// Reply will be of the form {"token": "TOKEN"}.
// 刷新token
func (mw *JWTUsecase) RefreshHandler(ctx context.Context, token string) (tokenString string, expire time.Time, err error) {
	claims, err := mw.checkIfTokenExpire(token)
	if err != nil {
		return
	}

	// Create the token
	newToken := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	newClaims := newToken.Claims.(jwt.MapClaims)

	for key := range claims {
		newClaims[key] = claims[key]
	}

	expire = mw.TimeFunc().Add(mw.Timeout)
	newClaims["exp"] = expire.Unix()
	newClaims["orig_iat"] = mw.TimeFunc().Unix()
	tokenString, err = mw.signedString(newToken)
	if err != nil {
		return
	}
	return
}

// @Summary 退出登录
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "msg": "成功退出系统" }"
// @Router /logout [post]
// @Security Bearer
func (mw *JWTUsecase) LogOut(ctx context.Context) {
	uastr, _ := meta.GetUA(ctx)
	remoteAddr, _ := meta.GetRemoteAddr(ctx)
	var loginlog model.SysLoginLog
	ua := user_agent.New(uastr)
	loginlog.Ipaddr = remoteAddr
	location := ip.GetLocation(remoteAddr)
	loginlog.LoginLocation = location
	loginlog.LoginTime = tools.GetCurrentTime()
	loginlog.Status = "0"
	loginlog.Remark = uastr
	browserName, browserVersion := ua.Browser()
	loginlog.Browser = browserName + " " + browserVersion
	loginlog.Os = ua.OS()
	loginlog.Platform = ua.Platform()
	// loginlog.Username = GetUserName(c)
	loginlog.Msg = "退出成功"

	mw.dao.CreateJwtAuthLoginLog(ctx, &loginlog)
}
