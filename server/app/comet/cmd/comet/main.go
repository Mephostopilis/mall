package main

import (
	"flag"

	"edu/comet/internal/conf"
	"edu/comet/internal/di"
	pkglog "edu/pkg/log"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/google/uuid"
	etcd "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// Branch is current branch name the code is built off.
	Branch string
	// Revision is the short commit hash of source tree.
	Revision string
	// BuildDate is the date when the binary was built.
	BuildDate string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf admin.yaml")
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}))
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	logger := pkglog.NewZapLogger()

	// uuid
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	// etcd
	cli, err := etcd.New(etcd.Config{
		Endpoints: bc.Reg.Etcd.Endpoints,
	})
	if err != nil {
		panic(err)
	}
	r := registry.New(cli)
	app, err := di.InitApp(id, bc.Service, bc.Server, bc.App, logger, r)
	if err != nil {
		panic(err)
	}
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
