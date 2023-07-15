package service

import (
	"context"
	"fmt"

	"github.com/dhuki/go-date/config"
	"github.com/dhuki/go-date/pkg/logger"
)

func (u userServiceImpl) temporarySuspendAccount(ctx context.Context, keyFailedAttempt, keySuspend string) (err error) {
	ctxName := fmt.Sprintf("%T.LogtemporarySuspendAccountin", u)

	if err = u.redisLib.Delete(keyFailedAttempt); err != nil {
		logger.Error(ctx, ctxName, "u.redisLib.Delete, got err: %v", err)
		return
	}
	if err = u.redisLib.Set(keySuspend, true, config.Conf.Redis.FailedLoginIssuspendTTL); err != nil {
		logger.Error(ctx, ctxName, "u.redisLib.Set, got err: %v", err)
		return
	}
	return
}
