package schema

import "embed"

// FS embeds the goose migration files so the migrate binary is self-contained
// and doesn't need the sql/schema directory shipped alongside it at runtime.
//
//go:embed *.sql
var FS embed.FS
