package video

import (
	"context"
	"errors"
	"fmt"

	model "w2-work4/biz/model/db"
	errme "w2-work4/biz/model/errormessage"
	video "w2-work4/biz/repository/video"

	"gorm.io/gorm"
)

type VideoSvc struct {
	videoRepo *video.VideoRepo
	rdb       *video.VideoCacheRepo
}

func NewVideoSvc(videorepo *video.VideoRepo, rdb *video.VideoCacheRepo) *VideoSvc {
	return &VideoSvc{videoRepo: videorepo, rdb: rdb}
}

func (s *VideoSvc) UploadVideo(ctx context.Context, userID int64, videoURL, coverURL, title, description string) error {
	if videoURL == "" || coverURL == "" || title == "" || description == "" {
		return errme.ErrMissingRequiredFields
	}
	if userID <= 0 {
		return errme.ErrInvalidUserID
	}
	video := &model.Video{
		UserID:      userID,
		VideoURL:    videoURL,
		CoverURL:    coverURL,
		Title:       title,
		Description: description,
	}
	err := s.videoRepo.AddVideo(ctx, video)
	if err != nil {
		return fmt.Errorf("failed to upload video: %w", err)
	}
	return nil
}

func (s *VideoSvc) GetVideo(ctx context.Context, id int64) (*model.Video, error) {
	if id <= 0 {
		return nil, errme.ErrInvalidVideoID
	}
	video, err := s.videoRepo.GetVideoByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errme.ErrVideoNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get video: %w", err)
	}
	visitcount, err := s.rdb.GetVideoVisitCount(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get video visit count: %w", err)
	}
	likecount, err := s.rdb.GetVideoLikeCount(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get video like count: %w", err)
	}
	video.LikeCount = likecount
	video.VisitCount = visitcount
	return video, nil
}

func (s *VideoSvc) GetVideos(ctx context.Context, ids []int64) ([]*model.Video, error) {
	if len(ids) == 0 {
		return nil, errme.ErrMissingRequiredFields
	}
	videos, err := s.videoRepo.GetVideosByIDs(ctx, ids)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errme.ErrVideoNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get videos: %w", err)
	}
	return videos, nil
}

func (s *VideoSvc) GetVideosByUserID(ctx context.Context, userID int64, page, pagesize int) ([]*model.Video, error) {
	if userID <= 0 {
		return nil, errme.ErrInvalidUserID
	}
	if page <= 0 || pagesize <= 0 {
		return nil, errme.ErrInvalidPageOrPageSize
	}
	videos, err := s.videoRepo.GetVideosByUserID(ctx, userID, page, pagesize)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errme.ErrVideoNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get videos by user ID: %w", err)
	}
	return videos, nil
}

func (s *VideoSvc) GetHotVideos(ctx context.Context, Page, PageSize int) ([]*model.Video, error) {
	if Page <= 0 || PageSize <= 0 {
		return nil, errme.ErrInvalidPageOrPageSize
	}
	ids, err := s.rdb.GetHotVideos(ctx, Page, PageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get hot videos: %w", err)
	}
	videos, err := s.videoRepo.GetVideosByIDs(ctx, ids)
	if errors.Is(err, gorm.ErrRecordNotFound) { //hot videos丢失
		return nil, errme.ErrVideoNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get hot videos: %w", err)
	}
	return videos, nil
}

func (s *VideoSvc) SearchVideos(ctx context.Context, keywords []string, Page, Pagesize int) ([]*model.Video, error) {
	if len(keywords) == 0 {
		return nil, errme.ErrMissingRequiredFields
	}
	if Page <= 0 || Pagesize <= 0 {
		return nil, errme.ErrInvalidPageOrPageSize
	}
	if err := s.rdb.IncreaseSearchCount(ctx, keywords); err != nil {
		return nil, fmt.Errorf("failed to increase search count: %w", err)
	}
	ids, err := s.rdb.GetVideoIDsByKeywords(ctx, keywords, Page, Pagesize)
	if err != nil {
		return nil, fmt.Errorf("failed to get video IDs by keywords: %w", err)
	}
	if len(ids) == 0 {
		res, err := s.videoRepo.GetVideosByKeywords(ctx, keywords, Page, Pagesize)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errme.ErrVideoNotFound
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get videos by keywords: %w", err)
		}
		if len(res) > 0 {
			videoIDs := make([]int64, 0, len(res))
			for _, video := range res {
				videoIDs = append(videoIDs, video.ID)
			}
			if err := s.rdb.CacheVideoIDsByKeywords(ctx, keywords, videoIDs); err != nil {
				return nil, fmt.Errorf("failed to cache video IDs by keywords: %w", err)
			}
		}
		return res, nil
	}
	videos, err := s.videoRepo.GetVideosByIDs(ctx, ids)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errme.ErrVideoNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get videos by IDs: %w", err)
	}
	return videos, nil

}

func (s *VideoSvc) DeleteVideo(ctx context.Context, id int64) error {
	if id <= 0 {
		return errme.ErrInvalidVideoID
	}
	err := s.videoRepo.DeleteVideoByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errme.ErrVideoNotFound
	}
	if err != nil {
		return fmt.Errorf("failed to delete video: %w", err)
	}
	err = s.rdb.DeleteVideoLikeCount(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete video like count: %w", err)
	}
	err = s.rdb.DeleteVideoVisitCount(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete video visit count: %w", err)
	}
	return nil
}
