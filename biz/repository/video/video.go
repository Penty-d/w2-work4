package video

import (
	"context"
	model "w2-work4/biz/model/db"

	"gorm.io/gorm"
)

type VideoRepo struct {
	db *gorm.DB
}

func NewVideoRepo(db *gorm.DB) *VideoRepo {
	return &VideoRepo{db: db}
}

func (r *VideoRepo) AddVideo(ctx context.Context, video *model.Video) error {
	return r.db.WithContext(ctx).Create(video).Error
}

func (r *VideoRepo) GetVideoByID(ctx context.Context, id int64) (*model.Video, error) {
	var video model.Video
	err := r.db.WithContext(ctx).Model(&model.Video{}).Where("id = ?", id).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, err
}

func (r *VideoRepo) GetVideosByUserID(ctx context.Context, userID int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.db.WithContext(ctx).Model(&model.Video{}).Where("user_id = ?", userID).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, err
}

func (r *VideoRepo) GetVideosByKeywords(ctx context.Context, keywords []string) ([]*model.Video, error) {
	var videos []*model.Video
	kw := make([]string, 0, len(keywords))
	for _, keyword := range keywords {
		kw = append(kw, "%"+keyword+"%")
	}
	err := r.db.WithContext(ctx).Model(&model.Video{}).Where("description ILIKE ANY (array[?]) OR title ILIKE ANY (array[?])", kw, kw).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, err
}

func (r *VideoRepo) GetVideosByVisitCount(ctx context.Context, limit int) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.db.WithContext(ctx).Model(&model.Video{}).Order("visit_count DESC").Limit(limit).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, err
}

func (r *VideoRepo) GetVideosByLikeCount(ctx context.Context, limit int) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.db.WithContext(ctx).Model(&model.Video{}).Order("like_count DESC").Limit(limit).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, err
}

func (r *VideoRepo) GetVideosByCommentCount(ctx context.Context, limit int) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.db.WithContext(ctx).Model(&model.Video{}).Order("comment_count DESC").Limit(limit).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, err
}
