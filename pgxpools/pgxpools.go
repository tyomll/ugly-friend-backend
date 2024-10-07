package pgxpools

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type ConfigConnectPgxPool struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func ConnectDB(configObj *ConfigConnectPgxPool) *pgxpool.Pool {
	logrus.Info("🟨 ConnectDB")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configObj.Host,
		configObj.Port,
		configObj.User,
		configObj.Password,
		configObj.Name, configObj.SSLMode)

	configDB, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal("📛 error: failed to parse config: ", err)
		return nil
	}

	configDB.MaxConns = 50
	configDB.MinConns = 5
	configDB.MaxConnLifetime = time.Second * 45
	configDB.MaxConnIdleTime = time.Second * 45
	configDB.HealthCheckPeriod = time.Second * 15

	pool, err := pgxpool.NewWithConfig(context.Background(), configDB)
	if err != nil {
		log.Fatal("📛 error: failed to connect to database: ", err)
		return nil
	}

	return pool
}
