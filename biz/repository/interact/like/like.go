package like

import (
	"context"
	model "w2-work4/biz/model/db"

	"gorm.io/gorm"
)

type LikeRepo struct {
	db *gorm.DB
}

func NewLikeRepo(db *gorm.DB) *LikeRepo {
	return &LikeRepo{db: db}
}

func (r *LikeRepo) CreateLike(ctx context.Context, like *model.VideoLike) error {
	return r.db.WithContext(ctx).Create(like).Error
}

func (r *LikeRepo) DeleteLikeByVideoIDAndUserID(ctx context.Context, videoID int64, userID int64) error {
	return r.db.WithContext(ctx).Where("video_id = ? AND user_id = ?", videoID, userID).Delete(&model.VideoLike{}).Error
}

func (r *LikeRepo) GetLikeByVideoID(ctx context.Context, videoID int64) (*model.VideoLike, error) {
	var like model.VideoLike
	err := r.db.WithContext(ctx).Model(&model.VideoLike{}).Where("video_id = ?", videoID).First(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, err
}

func (r *LikeRepo) GetLikeByUserID(ctx context.Context, userID int64) ([]*model.VideoLike, error) {
	var likes []*model.VideoLike
	err := r.db.WithContext(ctx).Model(&model.VideoLike{}).Where("user_id = ?", userID).Find(&likes).Error
	if err != nil {
		return nil, err
	}
	return likes, err
}
