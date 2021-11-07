package logging

import "context"

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

func (c ContextKey) GetFromContext(ctx context.Context) string {
	if val, ok := ctx.Value(c).(string); !ok {
		return ""
	} else {
		return val
	}
}

func (c ContextKey) AddToContext(ctx context.Context, value string) context.Context {
	ctx = context.WithValue(ctx, c, value)
	return ctx
}
