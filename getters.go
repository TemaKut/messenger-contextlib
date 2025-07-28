package contextlib

import (
	"context"
	"github.com/TemaKut/messenger-contextlib/dto"
)

func User(ctx context.Context) (dto.User, error) {
	value := ctx.Value(ContextKeyUser)
	if value == nil {
		return dto.User{}, ErrUserNotPassed
	}

	valueUser, ok := value.(dto.User)
	if !ok {
		return dto.User{}, ErrUserNotPassed
	}

	return valueUser, nil
}
