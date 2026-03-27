package social

import (
	"context"
	model "w2-work4/biz/model/db"

	"gorm.io/gorm"
)

type FollowRepo struct {
	db *gorm.DB
}

func NewFollowRepo(db *gorm.DB) *FollowRepo {
	return &FollowRepo{db: db}
}

func (r *FollowRepo) CreateFollow(ctx context.Context, follow *model.Follow) error {
	return r.db.WithContext(ctx).Create(follow).Error
}

func (r *FollowRepo) DeleteFollowByUserIDAndFollowUserID(ctx context.Context, userID int64, followUserID int64) error {
	return r.db.WithContext(ctx).Where("user_id = ? AND follow_user_id = ?", userID, followUserID).Delete(&model.Follow{}).Error
}

func (r *FollowRepo) GetFollowByUserID(ctx context.Context, userID int64) ([]*model.Follow, error) {
	var follows []*model.Follow
	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("user_id = ?", userID).Find(&follows).Error
	if err != nil {
		return nil, err
	}
	return follows, err
}

func (r *FollowRepo) GetFollowByFollowUserID(ctx context.Context, followUserID int64) ([]*model.Follow, error) {
	var follows []*model.Follow
	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("follow_user_id = ?", followUserID).Find(&follows).Error
	if err != nil {
		return nil, err
	}
	return follows, err
}

func (r *FollowRepo) GetFollowByUserIDAndFollowUserID(ctx context.Context, userID int64, followUserID int64) (*model.Follow, error) {
	var follow model.Follow
	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("user_id = ? AND follow_user_id = ?", userID, followUserID).First(&follow).Error
	if err != nil {
		return nil, err
	}
	return &follow, err
}
