package jwt

import "context"

func Authenticate(ctx context.Context) (context.Context, error) {
	return auth(ctx)
}

func auth(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
