/*
	This file (fs.go) is responsible for embedding
	SQL migration files into the application
*/

package migrations

import (
	"embed"
	_ "embed"
)

//go:embed *.sql
var FS embed.FS
