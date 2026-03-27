package video

import (
	"context"
	model "w2-work4/biz/model/db"
	video "w2-work4/biz/repository/video"
)

type VideoSvc struct {
	videoRepo *video.VideoRepo
}

func NewVideoSvc(videorepo *video.VideoRepo) *VideoSvc {
	return &VideoSvc{videoRepo: videorepo}
}

func (s *VideoSvc) UploadVideo(ctx context.Context, userID int64, videoURL, coverURL, title, description string) error {
	video := &model.Video{
		UserID:      userID,
		VideoURL:    videoURL,
		CoverURL:    coverURL,
		Title:       title,
		Description: description,
	}
	return s.videoRepo.AddVideo(ctx, video)
}

func (s *VideoSvc) GetVideoByID(ctx context.Context, id int64) (*model.Video, error) {
	return s.videoRepo.GetVideoByID(ctx, id)
}

func (s *VideoSvc) GetVideosByUserID(ctx context.Context, userID int64) ([]*model.Video, error) {
	return s.videoRepo.GetVideosByUserID(ctx, userID)
}
