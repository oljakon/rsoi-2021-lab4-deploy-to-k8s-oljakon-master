package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const dsn = "postgresql://hippo:,7_lQeiIM%5DgiO%3EF9danx%29%3Aoh@hippo-primary.postgres-operator.svc:5432/hippo"

//const dsn = "postgresql://postgres:postgres@postgres-rsoi:5432/postgres?sslmode=disable"
//const dsn = "postgres://gpsjekmvjzncqf:fd435ed58327f419e2327c2c2dfd11a9bc6b9ea4891fe91549d3fdd329abf5fe@ec2-174-129-16-183.compute-1.amazonaws.com:5432/dah4v24ca3aaaf"

// CreateConnection to persons db
func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
