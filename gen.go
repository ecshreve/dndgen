package ent

import (
	_ "github.com/hedwigz/entviz"
)

//go:generate go run -mod=mod ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
