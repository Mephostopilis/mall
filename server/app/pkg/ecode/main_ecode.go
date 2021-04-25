package ecode

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// main ecode interval is [0,990000]
func SignCheckErr(reason, format string, a ...interface{}) error {
	return errors.Errorf(10001, "签名错误", format, a...)
}

func RequestErr(reason, format string, a ...interface{}) error {
	return errors.Errorf(10002, "应用不存在错误", "")
}

func ServerErr(reason, format string, a ...interface{}) error {
	return errors.Errorf(10003, "token接口失效", "")
}

func KeyNotFound(reason, format string, a ...interface{}) error {
	return errors.Errorf(10004, "接口参数非法", "")
}

func ReadTimeout(reason, format string, a ...interface{}) error {
	return errors.Errorf(10100, "读取超时", "")
}

// admin jwt
// ErrMissingSecretKey indicates Secret key is required
func ErrMissingSecretKey(reason, format string, a ...interface{}) error {
	return errors.Errorf(11100, "secret key is required", "")
}

func ErrForbidden(reason, format string, a ...interface{}) error {
	return errors.Errorf(11101, "", "") // ErrForbidden when HTTP status 403 is give
}

var (
	// grom [8000, 9000)

	DialTimeout         = errors.Errorf(10101, "链接超时", "")
	WriteTimeout        = errors.Errorf(10102, "写入超时", "")
	EOF                 = errors.Errorf(10103, "", "")
	Reset               = errors.Errorf(10104, "", "")
	BrokenPipe          = errors.Errorf(10105, "", "")
	RpcConnectionClosed = errors.Errorf(10106, "", "")
	UnexpectedErr       = errors.Errorf(10107, "", "")

	ErrPassword     = errors.Errorf(10108, "", "")
	ErrUsername     = errors.Errorf(10109, "", "")
	ErrAesNewCipher = errors.Errorf(10110, "", "")

	// sso
	UserUpdateNameFailed       = errors.Errorf(10200, "", "")
	ErrUnsupportedGrantType    = errors.Errorf(10201, "", "")
	ErrInvalidRequest          = errors.Errorf(10202, "", "")
	ErrInvalidGrant            = errors.Errorf(10203, "", "")
	ErrUnsupportedResponseType = errors.Errorf(10204, "", "")
	ErrUnauthorizedClient      = errors.Errorf(10205, "", "")
	ErrToken                   = errors.Errorf(10206, "", "")
	ErrInvalidClient           = errors.Errorf(10207, "", "")
	ErrUnsupportedErrType      = errors.Errorf(10208, "", "")
	ErrInvalidScope            = errors.Errorf(10209, "", "")
	ErrInvalidAuthorizeCode    = errors.Errorf(10210, "", "")
	ErrExpiredRefreshToken     = errors.Errorf(10212, "", "")
	ErrInvalidUserId           = errors.Errorf(10213, "", "")
	ErrInvalidUsername         = errors.Errorf(10215, "", "")
	ErrInvalidUid              = errors.Errorf(10216, "", "")
	ErrAccountExist            = errors.Errorf(10217, "", "")
	ErrInvalidAccessToken      = errors.Errorf(10218, "", "")
	ErrInvalidCaptcha          = errors.Errorf(10220, "", "")
	ErrInvalidPassword         = errors.Errorf(10221, "", "")
	ErrInvalidRefreshToken     = errors.Errorf(10222, "", "")
	ErrInvalidParam            = errors.Errorf(10223, "", "")
	ErrKeyNotFound             = errors.Errorf(10224, "", "")
	UserNotExist               = errors.Errorf(10225, "", "")

	// ss
	SsVipNotExist = errors.Errorf(10301, "", "") // 不存在对应的vip
	SsKeyNotFound = errors.Errorf(10302, "", "") // 不存在对应的key
	SsInvalidMid  = errors.Errorf(10303, "", "") // 无效的mid
	SsLessThanVip = errors.Errorf(10304, "", "") // 无效的mid
	SsVipExist    = errors.Errorf(10305, "", "") // 存在对应的vip
	SsZoneExist   = errors.Errorf(10306, "", "") // 存在对应的vip

	// socks
	SocksPortExist   = errors.Errorf(10401, "", "") // 已经存在了
	SocksPortNoExist = errors.Errorf(10402, "", "") // 已经存在了
	SocksPFull       = errors.Errorf(10403, "", "") // 已经存在了
	SocksPExist      = errors.Errorf(10404, "", "") // 已经存在了
	SocksPClosed     = errors.Errorf(10405, "", "") // 已经存在了

	// account-free
	AFAccountExist = errors.Errorf(10501, "", "")

	// auth
	Unauthorized = errors.Errorf(10601, "", "")

	//
	ErrGetSysUser    = errors.Errorf(10701, "", "")
	ErrUpdateSysUser = errors.Errorf(10702, "", "")
	ErrDeptId        = errors.Errorf(10703, "", "")

	// mgo
	MgoNotFound = errors.Errorf(10800, "", "")

	// NCSC
	GTInvalidTx = errors.Errorf(10900, "", "")

	// Admin
	AdminTableEmpty      = errors.Errorf(11000, "", "")
	AdminNotSupportedJob = errors.Errorf(11001, "", "")

	ErrMissingAuthenticatorFunc = errors.Errorf(11102, "", "")                 // ErrMissingAuthenticatorFunc indicates Authenticator is required
	ErrMissingLoginValues       = errors.Errorf(11103, "", "")                 // ErrMissingLoginValues indicates a user tried to authenticate without username or password
	ErrFailedAuthentication     = errors.Errorf(11104, "", "")                 // ErrFailedAuthentication indicates authentication failed, could be faulty username or password
	ErrFailedTokenCreation      = errors.Errorf(11105, "", "")                 // ErrFailedTokenCreation indicates JWT Token failed to create, reason unknown
	ErrExpiredToken             = errors.Errorf(11106, "token is expired", "") // ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrEmptyAuthHeader          = errors.Errorf(11107, "", "")                 // ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrMissingExpField          = errors.Errorf(11108, "", "")                 // ErrMissingExpField missing exp field in token
	ErrWrongFormatOfExp         = errors.Errorf(11109, "", "")                 // ErrWrongFormatOfExp field must be float64 format
	ErrInvalidAuthHeader        = errors.Errorf(11110, "", "")                 // ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrEmptyQueryToken          = errors.Errorf(11111, "", "")                 // ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyCookieToken         = errors.Errorf(11112, "", "")                 // ErrEmptyCookieToken can be thrown if authing with a cookie, the token cokie is empty
	ErrEmptyParamToken          = errors.Errorf(11113, "", "")                 // ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrInvalidSigningAlgorithm  = errors.Errorf(11114, "", "")                 // ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidVerificationode   = errors.Errorf(11115, "", "")
	ErrNoPrivKeyFile            = errors.Errorf(11116, "", "") // ErrNoPrivKeyFile indicates that the given private key is unreadable
	ErrNoPubKeyFile             = errors.Errorf(11117, "", "") // ErrNoPubKeyFile indicates that the given public key is unreadable
	ErrInvalidPrivKey           = errors.Errorf(11118, "", "") // ErrInvalidPrivKey indicates that the given private key is invalid
	ErrInvalidPubKey            = errors.Errorf(11119, "", "") // ErrInvalidPubKey indicates the the given public key is invalid

	FileErr = errors.Errorf(111200, "", "")
)

// func init() {
// 	cm := map[int]string{
// 		4001: "参数错误",
// 		4002: "应用不存在错误",
// 		4003: "token接口失效",
// 		4004: "接口参数非法",
// 		4005: "记录不存在",
// 		4006: "签名错误",
// 		4007: "币种不存在",
// 		4008: "公链不存在",
// 		4009: "提交失败",
// 		4010: "接口请求时间超时",
// 		4011: "余额不足",
// 		4012: "存在相同的交易hash",
// 		4013: "处理中的请求",
// 		4014: "查询失败",
// 		4015: "不支持矿工费充值",
// 		4016: "存在相同的订单",
// 		4017: "接受地址格式错误",
// 		4018: "不支持总帐号向总帐号提币",
// 		4019: "提币数量太大或太小",
// 		4020: "汇总异常",
// 		4021: "GT生成地址错误",
// 		4022: "不存在此订单",
// 		4023: "地址解析错误",
// 		4024: "汇总已经成功",
// 		4025: "矿工是空，转换错误",

// 		5001: "系统异常",
// 		5002: "不支持其他币种",
// 		5003: "网络异常",
// 		5004: "应用授权过期",
// 		5005: "redis异常",

// 		6001: "不存在此纪录",
// 		6002: "矿工加速apollo填写错误",
// 		6003: "不能有效获取余额",

// 		10001: "该类型不支持投稿",
// 		10002: "数据库版本存在",
// 		10012: "redis得到错误",

// 		10100: "不支持的错误",
// 		10101: "交易没有在pending中",
// 		10102: "汇总成功",
// 		10103: "JSON解析错误",
// 		10104: "浮点转换错误",

// 		10200: "没有这个文档,数据库找不到",
// 		10201: "不支持的授权类型",

// 		10500: "寻找路劲失败",
// 		10503: "读取key文件失败",
// 		10504: "读取当前key文件",
// 		10505: "没有这个网络接口",
// 		10506: "网络地址错误",

// 		10601: "没有授权",

// 		11100: "secret key is required",
// 		11101: "you don't have permission to access this resource",
// 		11102: "ginJWTMiddleware.Authenticator func is undefined",
// 		11103: "missing Username or Password or Code",
// 		11104: "incorrect Username or Password",
// 		11105: "failed to create JWT Token",
// 		11106: "token is expired",
// 		11107: "auth header is empty",
// 		11108: "missing exp field",
// 		11109: "exp must be float64 format",
// 		11110: "auth header is invalid",
// 		11111: "query token is empty",
// 		11112: "cookie token is empty",
// 		11113: "parameter token is empty",
// 		11114: "invalid signing algorithm",
// 		11115: "验证码错误",
// 		11116: "private key file unreadable",
// 		11117: "public key file unreadable",
// 		11118: "private key invalid",
// 		11119: "public key invalid",
// 	}
// xecode.Register(cm)
// }
