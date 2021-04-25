// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package tools

import (
	"github.com/google/wire"
)

//go:generate kratos tool wire
func newTestDao() (*dao, func(), error) {
	panic(wire.Build(newDao))
}
