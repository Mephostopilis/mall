package biz

import (
	"context"
	"testing"

	"edu/service/im/internal/model"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/stretchr/testify/assert"
)

func TestNodes(t *testing.T) {
	var (
		c        = context.TODO()
		clientIP = "127.0.0.1"
	)
	lg.nodes = make([]*registry.ServiceInstance, 0)
	ins := lg.NodesInstances(c)
	assert.NotNil(t, ins)
	nodes := lg.NodesWeighted(c, model.PlatformWeb, clientIP)
	assert.NotNil(t, nodes)
	nodes = lg.NodesWeighted(c, "android", clientIP)
	assert.NotNil(t, nodes)
}
