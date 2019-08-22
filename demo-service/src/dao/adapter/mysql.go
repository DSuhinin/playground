package adapter

import (
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/KWRI/demo-service/core/errors"
	"github.com/KWRI/demo-service/core/http/health"
	"github.com/KWRI/demo-service/core/log"
)

//
// Provider provides and interface to work with database adapter.
//
type Provider interface {
	//
	// GetDB returns a DB instance.
	//
	GetDB() *sqlx.DB
}

const (
	//
	// ServiceName name returning in health info.
	//
	ServiceName = "mysql_db"
)

//
// Connection wraps cassandra session.
//
type Connection struct {
	db     *sqlx.DB
	logger log.Logger
}

//
// NewConnection returns MySQL connection object.
//
func NewConnection(logger log.Logger, connectionDSN string) (*Connection, error) {

	db, err := sqlx.Open("mysql", connectionDSN)
	if err != nil {
		return nil, errors.WithMessage(err, "database connection error")
	}

	return &Connection{
		db:     db,
		logger: logger,
	}, nil
}

//
// GetDB returns a DB instance.
//
func (c Connection) GetDB() *sqlx.DB {

	return c.db
}

//
// GetHealth returns the Cassandra Health info.
//
func (c Connection) GetHealth() (*health.Data, error) {

	h := health.Data{
		Name:   ServiceName,
		Status: http.StatusOK,
	}

	startTime := time.Now()
	defer func() {
		h.Latency = time.Since(startTime).Seconds()
	}()

	if err := c.db.Ping(); err != nil {
		h.Status = http.StatusBadRequest
		return &h, errors.New("mysql health check error:%+v", err)

	}

	return &h, nil
}

//
// Shutdown closes connection with Cassandra service. Method needs for graceful shutdown.
//
func (c Connection) Shutdown() {

	if err := c.db.Close(); err != nil {
		c.logger.Info("impossible to close mysql connection: %+v", err)
	} else {
		c.logger.Info("mysql connection has been closed")
	}
}
