package ctx

import "context"

type Context struct {
	Context context.Context
	Data    map[string]any
}
