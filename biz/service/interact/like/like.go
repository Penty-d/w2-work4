package like

import (
	"context"
	"fmt"
	model "w2-work4/biz/model/db"
	errme "w2-work4/biz/model/errormessage"
	"w2-work4/biz/repository/interact/like"
)

type LikeSvc struct {
	likeRepo *like.LikeRepo
	rdb      *like.LikeCacheRepo
}

func NewLikeSvc(likeRepo *like.LikeRepo, rdb *like.LikeCacheRepo) *LikeSvc {
	return &LikeSvc{likeRepo: likeRepo, rdb: rdb}
}

func (s *LikeSvc) LikeVideo(ctx context.Context, userID, videoID int64) error {
	if userID <= 0 {
		return errme.ErrInvalidUserID
	}
	if videoID <= 0 {
		return errme.ErrInvalidVideoID
	}
	like := &model.VideoLike{
		UserID:  userID,
		VideoID: videoID,
	}
	err := s.likeRepo.CreateLike(ctx, like)
	if err != nil {
		return fmt.Errorf("failed to like video: %w", err)
	}
	err = s.rdb.IncreaseVideoLikeCount(ctx, videoID)
	if err != nil {
		return fmt.Errorf("failed to increase video like count: %w", err)
	}
	return nil
}

func (s *LikeSvc) UnlikeVideo(ctx context.Context, userID, videoID int64) error {
	if userID <= 0 {
		return errme.ErrInvalidUserID
	}
	if videoID <= 0 {
		return errme.ErrInvalidVideoID
	}
	err := s.likeRepo.DeleteLikeByVideoIDAndUserID(ctx, videoID, userID)
	if err != nil {
		return fmt.Errorf("failed to unlike video: %w", err)
	}
	err = s.rdb.DecreaseVideoLikeCount(ctx, videoID)
	if err != nil {
		return fmt.Errorf("failed to decrease video like count: %w", err)
	}
	return nil
}

func (s *LikeSvc) GetVideoLikeUsers(ctx context.Context, videoID int64) ([]int64, error) {
	if videoID <= 0 {
		return nil, errme.ErrInvalidVideoID
	}
	userIDs, err := s.likeRepo.GetUserIDsByVideoID(ctx, videoID)
	if err != nil {
		return nil, fmt.Errorf("failed to get video like users: %w", err)
	}
	return userIDs, nil
}
