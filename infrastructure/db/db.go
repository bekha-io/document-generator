package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS templates (
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	template_file TEXT NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS documents (
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	template_id VARCHAR(255) NOT NULL,
	template_data JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

	CONSTRAINT fk_template
      FOREIGN KEY(template_id) 
	  REFERENCES templates(id)
	  ON DELETE SET NULL

);
`

func NewSqlxConnection(driverName, dsn string) *sqlx.DB {
	db := sqlx.MustConnect("postgres", dsn)
	_ = db.MustExec(schema)
	return db
}
