//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/Khan/genqlient"
	_ "github.com/Yamashou/gqlgenc"
)
