package migrate

import (
	"database/sql"
	"edu/service/sso/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp(c *conf.Data, logger log.Logger) error {
	log := log.NewHelper(log.With(logger, "module", "migrate"))
	db, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Errorf("err = %v", err)
		return err
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Errorf("err = %v", err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"mysql", driver)
	if err != nil {
		log.Errorf("err = %v", err)
		return err
	}
	if err = m.Up(); err != nil {
		if err.Error() == "no change" {
			goto handleDone
		}
		log.Errorf("err = %v", err)
		return err
	}
handleDone:
	err1, err2 := m.Close()
	if err1 != nil {
		log.Errorf("err = %v", err1)
		return err1
	}
	if err2 != nil {
		log.Errorf("err = %v", err2)
		return err2
	}
	log.Debugf("数据库结构初始化成功！")
	return nil
}
