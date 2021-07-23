package biz

import (
	"io/ioutil"
	"time"

	pb "edu/api/sso/v1"
	"edu/pkg/ecode"

	"github.com/golang-jwt/jwt"
)

type MapClaims map[string]interface{}

type PayloadHandler func(data interface{}) MapClaims

type IdentityHandler func(MapClaims) (interface{}, error)

// JWT config for jwt
type Jwt struct {
	Secret  string
	Timeout time.Duration
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

func (mw *JWTUsecase) readKeys() (err error) {
	if mw.privateKey(); err != nil {
		return
	}
	if mw.publicKey(); err != nil {
		return err
	}
	return
}

func (mw *JWTUsecase) privateKey() (err error) {
	keyData, err := ioutil.ReadFile(mw.PrivKeyFile)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}
	mw.privKey = key
	return
}

func (mw *JWTUsecase) publicKey() (err error) {
	keyData, err := ioutil.ReadFile(mw.PubKeyFile)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}
	mw.pubKey = key
	return
}

func (mw *JWTUsecase) usingPublicKeyAlgo() bool {
	switch mw.SigningAlgorithm {
	case "RS256", "RS512", "RS384":
		return true
	}
	return false
}

// GetClaimsFromJWT get claims from JWT token
func (mw *JWTUsecase) getClaimsFromJWT(tokenstr string) (MapClaims, error) {
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

func (mw *JWTUsecase) signedString(token *jwt.Token) (tokenString string, err error) {
	if mw.usingPublicKeyAlgo() {
		tokenString, err = token.SignedString(mw.privKey)
	} else {
		tokenString, err = token.SignedString(mw.Key)
	}
	return
}

// ParseTokenString parse jwt token string
func (mw *JWTUsecase) parseTokenString(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(mw.SigningAlgorithm) != t.Method {
			return nil, pb.ErrInvalidSigningAlgorithm("")
		}
		if mw.usingPublicKeyAlgo() {
			return mw.pubKey, nil
		}

		return mw.Key, nil
	})
}

// CheckIfTokenExpire check if token expire
func (mw *JWTUsecase) checkIfTokenExpire(tokenstr string) (jwt.MapClaims, error) {
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
		return nil, pb.ErrExpiredToken("")
	}
	return claims, nil
}

func (mw *JWTUsecase) ValidationToken(tokenstr string) (out interface{}, err error) {
	claims, err := mw.getClaimsFromJWT(tokenstr)
	if err != nil {
		return
	}
	// 验证时效
	if claims["exp"] == nil {
		mw.log.Errorf("claims exp is null")
		err = pb.ErrMissingExpField("")
		return
	}
	if _, ok := claims["exp"].(float64); !ok {
		mw.log.Errorf("claims exp err wrong formate")
		err = pb.ErrWrongFormatOfExp("")
		return
	}
	if int64(claims["exp"].(float64)) < mw.TimeFunc().Unix() {
		err = pb.ErrExpiredToken("")
		return
	}
	return mw.IdentityFunc(claims)
}

// TokenGenerator method that clients can use to get a jwt token.
func (mw *JWTUsecase) TokenGenerator(data interface{}) (tokenString string, expire time.Time, err error) {
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
