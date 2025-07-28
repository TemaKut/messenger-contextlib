package contextlib

import (
	"context"
	"github.com/TemaKut/messenger-contextlib/dto"
)

func SetUser(ctx context.Context, user dto.User) context.Context {
	return context.WithValue(ctx, ContextKeyUser, user)
}
