//nolint
package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	servicename "github.com/NpoolPlatform/ledger-middleware/pkg/servicename"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

const (
	keyUsername  = "username"
	keyPassword  = "password"
	keyDBName    = "database_name"
	maxOpen      = 5
	maxIdle      = 2
	MaxLife      = 0
	keyServiceID = "serviceid"
)

func lockKey() string {
	serviceID := config.GetStringValueWithNameSpace(servicename.ServiceDomain, keyServiceID)
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixMigrate, serviceID)
}

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsn", "error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		username, password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(hostname string) (conn *sql.DB, err error) {
	hdsn, err := dsn(hostname)
	if err != nil {
		return nil, err
	}

	logger.Sugar().Warnw("open", "hdsn", hdsn)

	conn, err = sql.Open("mysql", hdsn)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.

	conn.SetConnMaxLifetime(time.Minute * MaxLife)
	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	return conn, nil
}

func Migrate(ctx context.Context) error {
	return nil
}
