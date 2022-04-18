package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Roach holds the connection pool to the database - created by a configuration
// object (`Config`).
type Roach struct {
	// Db holds a sql.DB pointer that represents a pool of zero or more
	// underlying connections - safe for concurrent use by multiple
	// goroutines -, with freeing/creation of new connections all managed
	// by `sql/database` package.
	Db   *sql.DB
	PgDb *pg.DB
	cfg  Config
}

// Config holds the configuration used for instantiating a new Roach.
type Config struct {
	// Address that locates our postgres instance
	Host string
	// Port to connect to
	Port string
	// User that has access to the database
	User string
	// Password so that the user can login
	Password string
	// Database to connect to (must have been created priorly)
	Database string
}

var (
	once sync.Once

	instance Roach
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(ctx context.Context, event *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d dbLogger) AfterQuery(ctx context.Context, event *pg.QueryEvent) error {
	query, _ := event.FormattedQuery()
	fmt.Printf("%s", query)
	println()
	return nil
}

// New returns a Roach with the sql.DB set with the postgres
// DB connection string in the configuration
func NewConnection(cfg Config) (roach Roach, err error) {
	once.Do(func() {
		if cfg.Host == "" || cfg.Port == "" || cfg.User == "" ||
			cfg.Password == "" || cfg.Database == "" {
			err = errors.Errorf(
				"All fields must be set (%s)",
				spew.Sdump(cfg))
			return
		}
		roach.cfg = cfg
		connection := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
		db, err := sql.Open("postgres", connection)
		if err != nil {
			err = errors.Wrapf(err,
				"Couldn't open connection to postgre database (%s)",
				spew.Sdump(cfg))
			return
		}

		dbOrm := pg.Connect(&pg.Options{
			Addr:     fmt.Sprintf(":%s", cfg.Port),
			User:     cfg.User,
			Password: cfg.Password,
			Database: cfg.Database,
		})
		ctx := context.Background()
		if err = dbOrm.Ping(ctx); err != nil {
			err = errors.Wrapf(err,
				"Couldn't ping postgre database (%s)",
				spew.Sdump(cfg))
			return
		}

		if err = db.Ping(); err != nil {
			err = errors.Wrapf(err,
				"Couldn't ping postgre database (%s)",
				spew.Sdump(cfg))
			return
		}
		dbOrm.AddQueryHook(dbLogger{})

		log.Print("Connect DB Success")
		// Config connection pool
		db.SetMaxOpenConns(70)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(5 * time.Minute)
		roach.Db = db
		roach.PgDb = dbOrm
		instance = roach
	})

	return instance, nil
}

// Todo: chưa check dc instance connection nil
func GetConnection() Roach {
	return instance
}

// Close performs the release of any resources that
// `sql/database` DB pool created. This is usually meant
// to be used in the exitting of a program or `panic`ing.
func (r *Roach) Close() (err error) {
	if r.Db == nil {
		return
	}

	if err = r.Db.Close(); err != nil {
		err = errors.Wrapf(err,
			"Errored closing database connection",
			spew.Sdump(r.cfg))
	}
	return
}
