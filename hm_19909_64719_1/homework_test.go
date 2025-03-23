package homework

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestDatabase struct {
	instance testcontainers.Container
}

func NewTestDatabase(t *testing.T) *TestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:12",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "postgres",
			"POSTGRES_SSL":      "off",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	return &TestDatabase{
		instance: postgres,
	}
}

func (db *TestDatabase) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.instance.MappedPort(ctx, "5432")
	require.NoError(t, err)
	return p.Int()
}

func (db *TestDatabase) ConnectionString(t *testing.T) string {
	return fmt.Sprintf("postgres://postgres:postgres@127.0.0.1:%d/postgres", db.Port(t))
}

func (db *TestDatabase) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, db.instance.Terminate(ctx))
}

func TestPostgresqlPing(t *testing.T) {
	testDatabase := NewTestDatabase(t)
	defer testDatabase.Close(t)

	connStr := testDatabase.ConnectionString(t)

	db, err := sql.Open("postgres", connStr+"?sslmode=disable")
	require.NoError(t, err)

	defer func(db *sql.DB) {
		err = db.Close()
		require.NoError(t, err)
	}(db)

	if err = db.Ping(); err != nil {
		require.NoError(t, err)
	}
}

func TestPostgresqlCreateDatabase(t *testing.T) {
	testDatabase := NewTestDatabase(t)
	defer testDatabase.Close(t)

	connStr := testDatabase.ConnectionString(t)

	db, err := sql.Open("postgres", connStr+"?sslmode=disable")
	require.NoError(t, err)

	createTablePostgres := `create table Students (Name TEXT, Roll_number INT PRIMARY KEY);`
	_, err = db.Exec(createTablePostgres)
	require.NoError(t, err)

	defer func(db *sql.DB) {
		err = db.Close()
		require.NoError(t, err)
	}(db)

	if err = db.Ping(); err != nil {
		require.NoError(t, err)
	}
}
