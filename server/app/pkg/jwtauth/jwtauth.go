package jwtauth

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"edu/pkg/ecode"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/log"
)

type MapClaims map[string]interface{}

type PayloadHandler func(data interface{}) MapClaims

type IdentityHandler func(MapClaims) (interface{}, error)

// JWT config for jwt
type Jwt struct {
	Secret  string
	Timeout time.Duration
}

// GinJWTMiddleware provides a Json-Web-Token authentication implementation. On failure, a 401 HTTP response
// is returned. On success, the wrapped middleware is called, and the userID is made available as
// c.Get("userID").(string).
// Users can get a token by posting a json request to LoginHandler. The token then needs to be passed in
// the Authentication header. Example: Authorization:Bearer XXX_TOKEN_XXX
type GinJWTMiddleware struct {
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
}

var (
	// IdentityKey default identity key
	IdentityKey  = "identity"
	NiceKey      = "nice"
	DataScopeKey = "datascope"
	RKey         = "r"
	RoleIdKey    = "roleid"
	RoleKey      = "rolekey"
	RoleNameKey  = "rolename"
)

// New 分配
func New(jwtCfg *Jwt, logger log.Logger, payload PayloadHandler, identity IdentityHandler) (*GinJWTMiddleware, error) {
	log := log.NewHelper("middleware/jwtauth", logger)
	du := time.Duration(jwtCfg.Timeout) * time.Second
	mw := &GinJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: "HS256",
		Key:              []byte(jwtCfg.Secret),
		Timeout:          du,
		MaxRefresh:       time.Hour,
		PayloadFunc:      payload,
		IdentityFunc:     identity,
		TimeFunc:         time.Now,
		log:              log,
	}
	return mw, nil
}

func (mw *GinJWTMiddleware) readKeys() (err error) {
	if mw.privateKey(); err != nil {
		return
	}
	if mw.publicKey(); err != nil {
		return err
	}
	return
}

func (mw *GinJWTMiddleware) privateKey() (err error) {
	keyData, err := ioutil.ReadFile(mw.PrivKeyFile)
	if err != nil {
		err = ecode.ErrNoPrivKeyFile
		return
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		err = ecode.ErrInvalidPrivKey
		return
	}
	mw.privKey = key
	return
}

func (mw *GinJWTMiddleware) publicKey() (err error) {
	keyData, err := ioutil.ReadFile(mw.PubKeyFile)
	if err != nil {
		return ecode.ErrNoPubKeyFile
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return ecode.ErrInvalidPubKey
	}
	mw.pubKey = key
	return
}

func (mw *GinJWTMiddleware) usingPublicKeyAlgo() bool {
	switch mw.SigningAlgorithm {
	case "RS256", "RS512", "RS384":
		return true
	}
	return false
}

// GetClaimsFromJWT get claims from JWT token
func (mw *GinJWTMiddleware) getClaimsFromJWT(tokenstr string) (MapClaims, error) {
	token, err := mw.parseTokenString(tokenstr)
	if err != nil {
		return nil, err
	}
	claims := MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}
	return claims, nil
}

func (mw *GinJWTMiddleware) signedString(token *jwt.Token) (tokenString string, err error) {
	if mw.usingPublicKeyAlgo() {
		tokenString, err = token.SignedString(mw.privKey)
	} else {
		tokenString, err = token.SignedString(mw.Key)
	}
	return
}

// ParseTokenString parse jwt token string
func (mw *GinJWTMiddleware) parseTokenString(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(mw.SigningAlgorithm) != t.Method {
			return nil, ecode.ErrInvalidSigningAlgorithm
		}
		if mw.usingPublicKeyAlgo() {
			return mw.pubKey, nil
		}

		return mw.Key, nil
	})
}

// CheckIfTokenExpire check if token expire
func (mw *GinJWTMiddleware) checkIfTokenExpire(tokenstr string) (jwt.MapClaims, error) {
	token, err := mw.parseTokenString(tokenstr)
	if err != nil {
		// If we receive an error, and the error is anything other than a single
		// ValidationErrorExpired, we want to return the error.
		// If the error is just ValidationErrorExpired, we want to continue, as we can still
		// refresh the token if it's within the MaxRefresh time.
		// (see https://github.com/appleboy/gin-jwt/issues/176)
		validationErr, ok := err.(*jwt.ValidationError)
		if !ok || validationErr.Errors != jwt.ValidationErrorExpired {
			return nil, err
		}
	}

	claims := token.Claims.(jwt.MapClaims)
	origIat := int64(claims["orig_iat"].(float64))
	if origIat < mw.TimeFunc().Add(-mw.MaxRefresh).Unix() {
		return nil, ecode.ErrExpiredToken
	}
	return claims, nil
}

func (mw *GinJWTMiddleware) ValidationToken(tokenstr string) (out interface{}, err error) {
	claims, err := mw.getClaimsFromJWT(tokenstr)
	if err != nil {
		return
	}
	// 验证时效
	if claims["exp"] == nil {
		mw.log.Errorf("claims exp is null")
		err = ecode.ErrMissingExpField
		return
	}
	if _, ok := claims["exp"].(float64); !ok {
		mw.log.Errorf("claims exp err wrong formate")
		err = ecode.ErrWrongFormatOfExp
		return
	}
	if int64(claims["exp"].(float64)) < mw.TimeFunc().Unix() {
		err = ecode.ErrExpiredToken
		return
	}
	return mw.IdentityFunc(claims)
}

// TokenGenerator method that clients can use to get a jwt token.
func (mw *GinJWTMiddleware) TokenGenerator(data interface{}) (tokenString string, expire time.Time, err error) {
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if mw.PayloadFunc != nil {
		for key, value := range mw.PayloadFunc(data) {
			claims[key] = value
		}
	}

	expire = mw.TimeFunc().UTC().Add(mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()
	tokenString, err = mw.signedString(token)
	if err != nil {
		return
	}
	return
}

func (mw *GinJWTMiddleware) RefreshToken(tokenstr string) (tokenString string, expire time.Time, err error) {
	claims, err := mw.checkIfTokenExpire(tokenstr)
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
