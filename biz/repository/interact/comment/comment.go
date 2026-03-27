package comment

import (
	"context"
	model "w2-work4/biz/model/db"

	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

func (r *CommentRepo) CreateComment(ctx context.Context, comment *model.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *CommentRepo) GetCommentsByVideoID(ctx context.Context, videoID int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := r.db.WithContext(ctx).Model(&model.Comment{}).Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (r *CommentRepo) DeleteCommentByIDAndUserID(ctx context.Context, commentID int64, userID int64) error {
	return r.db.WithContext(ctx).Where("id = ? AND user_id = ?", commentID, userID).Delete(&model.Comment{}).Error
}
