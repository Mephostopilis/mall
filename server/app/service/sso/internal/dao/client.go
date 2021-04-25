package dao

import (
	"context"

	"edu/service/sso/internal/model"

	"github.com/go-oauth2/oauth2/v4"
)

// GetByID according to the ID for the client information
func (d *dao) GetByID(ctx context.Context, id string) (cli oauth2.ClientInfo, err error) {
	m, err := d.GetSsoApp(ctx, &model.SsoApp{Id: id})
	if err != nil {
		return
	}
	cli = &m
	return
}
