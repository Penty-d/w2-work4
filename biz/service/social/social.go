package social

import (
	"context"
	"fmt"
	model "w2-work4/biz/model/db"
	errme "w2-work4/biz/model/errormessage"
	"w2-work4/biz/repository/social"
)

type FollowSvc struct {
	followRepo *social.FollowRepo
}

func NewFollowSvc(followRepo *social.FollowRepo) *FollowSvc {
	return &FollowSvc{followRepo: followRepo}
}

func (s *FollowSvc) FollowUser(ctx context.Context, userID, followUserID int64) error {
	if userID <= 0 || followUserID <= 0 {
		return errme.ErrInvalidUserID
	}
	if userID == followUserID {
		return errme.ErrCannotFollowSelf
	}
	follow := &model.Follow{
		UserID:       userID,
		FollowUserID: followUserID,
	}
	err := s.followRepo.CreateFollow(ctx, follow)
	if err != nil {
		return fmt.Errorf("failed to follow user: %w", err)
	}
	return nil
}
