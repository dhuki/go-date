package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/dhuki/go-date/config"
	modelReq "github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	"github.com/dhuki/go-date/pkg/internal/core/candidate/domain"
	"github.com/dhuki/go-date/pkg/internal/core/candidate/port"
	"github.com/dhuki/go-date/pkg/logger"
)

type candidateServiceImpl struct {
	repository port.CandidateRepository
	redisLib   port.RedisLibs
}

func NewCandidateService(repository port.CandidateRepository, redisLib port.RedisLibs) port.CandidateService {
	return candidateServiceImpl{
		repository: repository,
		redisLib:   redisLib,
	}
}

func (u candidateServiceImpl) SwipeAction(ctx context.Context, candidateID uint64, swipeDirection string) (err error) {
	ctxName := fmt.Sprintf("%T.SwipeAction", u)

	userID, err := strconv.ParseUint(fmt.Sprint(ctx.Value(config.ValueUserIDctx)), 10, 64)
	if err != nil {
		return err
	}

	relationType, err := domain.TranslateSwipeAction(swipeDirection)
	if err != nil {
		logger.Error(ctx, ctxName, "domain.TranslateSwipeAction, got err: %v", err)
		return
	}

	relationUser, err := u.repository.GetRelationUserByUserIdAndCandidate(ctx, userID, candidateID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(ctx, ctxName, "u.repository.GetRelationUserByUserIdAndCandidate, got err: %v", err)
		return
	}

	tx, err := u.repository.Start(ctx)
	if err != nil {
		logger.Error(ctx, ctxName, "u.repository.Start, got err: %v", err)
		return
	}
	defer func() {
		if err := u.repository.Finish(ctx, tx, err); err != nil {
			logger.Error(ctx, ctxName, "u.repository.Finish, got err: %v", err)
		}
	}()

	if err = u.repository.UpsertRelationUser(ctx, tx, modelRepo.RelationUser{
		ID:           relationUser.ID,
		UserID:       userID,
		CandidateID:  candidateID,
		RelationType: relationType,
		CreatedAt:    relationUser.CreatedAt,
	}); err != nil {
		logger.Error(ctx, ctxName, "u.repository.UpsertRelationUser, got err: %v", err)
		return
	}

	return nil
}

func (u candidateServiceImpl) GetListCandidate(ctx context.Context, limit int) (resp modelReq.CandidateListPaginationReponse, err error) {
	ctxName := fmt.Sprintf("%T.GetListCandidate", u)

	userID, err := strconv.ParseUint(fmt.Sprint(ctx.Value(config.ValueUserIDctx)), 10, 64)
	if err != nil {
		logger.Error(ctx, ctxName, "strconv.ParseUint, got err: %v", err)
		return
	}

	user, err := u.repository.GetUserByID(ctx, userID)
	if err != nil {
		logger.Error(ctx, ctxName, "u.repository.GetUserByID, got err: %v", err)
		return
	}

	countTotalUser, err := u.repository.CountTotalUserPagination(ctx, user.ID, user.Gender)
	if err != nil {
		logger.Error(ctx, ctxName, "u.repository.CountTotalUserPagination, got err: %v", err)
		return
	}

	if resp.TotalPage = u.getTotalPage(countTotalUser, limit); resp.TotalPage == 0 {
		return
	}

	page, keyLastPagination := 1, fmt.Sprintf("%s.%d", domain.KeyLastPagination, userID)
	if value := u.redisLib.Get(keyLastPagination); len(value) > 0 {
		if page, err = strconv.Atoi(value); err != nil {
			logger.Error(ctx, ctxName, "strconv.Atoi, got err: %v", err)
			return
		}
		if page+1 <= resp.TotalPage+1 {
			page++
		}
	}
	offset := (page - 1) * limit

	users, err := u.repository.GetUserPagination(ctx, user.ID, user.Gender, limit, offset)
	if err != nil {
		logger.Error(ctx, ctxName, "u.repoUser.GetUserPagination, got err: %v", err)
		return
	}

	for _, v := range users {
		resp.CandidateList = append(resp.CandidateList, modelReq.CandidateListReponse{
			ID:        v.ID,
			Username:  v.Username,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			PicUrl:    v.PicUrl,
			District:  v.District,
			City:      v.City,
		})
	}
	resp.Page = page

	if err = u.redisLib.Set(keyLastPagination, page, config.Conf.Redis.LastPageCandidateTTL); err != nil {
		logger.Warn(ctx, ctxName, "u.redisLib.Set, got err: %v", err)
	}

	return
}
