package dao

import (
	"context"

	"edu/service/sso/internal/model"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/google/uuid"
)

// Create create and store the new token information
func (d *dao) Create(info oauth2.TokenInfo) (err error) {
	// ct := time.Now()
	basicID := uuid.Must(uuid.NewRandom()).String()
	var m model.SsoToken
	m.Id = basicID
	m.ClientId = info.GetClientID()
	m.UserId = info.GetUserID()
	m.RedirectUri = info.GetRedirectURI()
	m.Scope = info.GetScope()
	m.Code = info.GetCode()
	m.CodeCreatedAt = info.GetCodeCreateAt()
	m.SetCodeExpiresIn(info.GetCodeExpiresIn())
	m.Access = info.GetAccess()
	m.AccessCreatedAt = info.GetAccessCreateAt()
	m.SetAccessExpiresIn(info.GetAccessExpiresIn())
	m.Refresh = info.GetRefresh()
	m.RefreshCreatedAt = info.GetRefreshCreateAt()
	m.SetRefreshExpiresIn(info.GetRefreshExpiresIn())
	_, err = d.CreateSsoToken(context.Background(), &m)
	if err != nil {
		return
	}
	// if code := info.GetCode(); code != "" {
	// 	err = ts.db.Put([]byte(code), []byte(jv), nil)
	// 	return
	// }

	// aexp := info.GetAccessExpiresIn()
	// rexp := aexp
	// // expires := true
	// if refresh := info.GetRefresh(); refresh != "" {
	// 	rexp = info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn()).Sub(ct)
	// 	if aexp.Seconds() > rexp.Seconds() {
	// 		aexp = rexp
	// 	}
	// 	// expires = info.GetRefreshExpiresIn() != 0
	// 	if err = ts.db.Put([]byte(refresh), []byte(basicID), nil); err != nil {
	// 		return
	// 	}
	// }
	// // opt.Options{}
	// // o := &opt.WriteOptions{}
	// if err = ts.db.Put([]byte(basicID), jv, nil); err != nil {
	// 	return
	// }
	// err = ts.db.Put([]byte(info.GetAccess()), []byte(basicID), nil)
	return
}

// RemoveByCode use the authorization code to delete the token information
func (d *dao) RemoveByCode(code string) error {
	_, err := d.DeleteSsoTokenX(context.TODO(), &model.SsoToken{Code: code})
	if err != nil {
		return err
	}
	return nil
}

// RemoveByAccess use the access token to delete the token information
func (d *dao) RemoveByAccess(access string) error {
	_, err := d.DeleteSsoTokenX(context.TODO(), &model.SsoToken{Access: access})
	if err != nil {
		return err
	}
	return nil
}

// RemoveByRefresh use the refresh token to delete the token information
func (d *dao) RemoveByRefresh(refresh string) error {
	_, err := d.DeleteSsoTokenX(context.TODO(), &model.SsoToken{Refresh: refresh})
	if err != nil {
		return err
	}
	return nil
}

// GetByCode use the authorization code for token information data
func (d *dao) GetByCode(code string) (oauth2.TokenInfo, error) {
	m, err := d.GetSsoToken(context.TODO(), &model.SsoToken{Code: code})
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetByAccess use the access token for token information data
func (d *dao) GetByAccess(access string) (oauth2.TokenInfo, error) {
	m, err := d.GetSsoToken(context.TODO(), &model.SsoToken{Access: access})
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetByRefresh use the refresh token for token information data
func (d *dao) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	m, err := d.GetSsoToken(context.TODO(), &model.SsoToken{Refresh: refresh})
	if err != nil {
		return nil, err
	}
	return &m, nil
}
