package comment

import (
	"context"
	"fmt"

	model "w2-work4/biz/model/db"
	errme "w2-work4/biz/model/errormessage"
	comment "w2-work4/biz/repository/interact/comment"
)

type CommentSvc struct {
	commentRepo *comment.CommentRepo
}

func NewCommentSvc(commentRepo *comment.CommentRepo) *CommentSvc {
	return &CommentSvc{commentRepo: commentRepo}
}

func (s *CommentSvc) CommentVideo(ctx context.Context, userID, videoID int64, content string) error {
	if userID <= 0 {
		return errme.ErrInvalidUserID
	}
	if videoID <= 0 {
		return errme.ErrInvalidVideoID
	}
	if content == "" {
		return errme.ErrMissingRequiredFields
	}
	comment := &model.Comment{
		UserID:  userID,
		VideoID: videoID,
		Content: content,
	}
	err := s.commentRepo.CreateComment(ctx, comment)
	if err != nil {
		return fmt.Errorf("failed to create comment: %w", err)
	}
	return nil
}

func (s *CommentSvc) GetComments(ctx context.Context, VideoID int64, page, pagesize int) ([]*model.Comment, error) {
	if VideoID <= 0 {
		return nil, errme.ErrInvalidVideoID
	}
	if page <= 0 || pagesize <= 0 {
		return nil, errme.ErrInvalidPageOrPageSize
	}
	comments, err := s.commentRepo.GetCommentsByVideoID(ctx, VideoID, page, pagesize)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}
	return comments, nil
}

func (s *CommentSvc) DeleteComment(ctx context.Context, commentID, userID int64) error {
	if commentID <= 0 {
		return errme.ErrInvalidCommentID
	}
	if userID <= 0 {
		return errme.ErrInvalidUserID
	}
	err := s.commentRepo.DeleteCommentByIDAndUserID(ctx, commentID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}
	return nil
}
