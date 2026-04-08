package video

import (
	"context"

	model "w2-work4/biz/model/db"
	video "w2-work4/biz/repository/video"
)

type VideoSvc struct {
	videoRepo *video.VideoRepo
	rdb       *video.VideoCacheRepo
}

func NewVideoSvc(videorepo *video.VideoRepo, rdb *video.VideoCacheRepo) *VideoSvc {
	return &VideoSvc{videoRepo: videorepo, rdb: rdb}
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

func (s *VideoSvc) GetVideo(ctx context.Context, id int64) (*model.Video, error) {
	return s.videoRepo.GetVideoByID(ctx, id)
}

func (s *VideoSvc) GetVideos(ctx context.Context, ids []int64) ([]*model.Video, error) {
	return s.videoRepo.GetVideosByIDs(ctx, ids)
}

func (s *VideoSvc) GetVideosByUserID(ctx context.Context, userID int64, page, pagesize int) ([]*model.Video, error) {
	return s.videoRepo.GetVideosByUserID(ctx, userID, page, pagesize)
}

func (s *VideoSvc) GetHotVideos(ctx context.Context, Page, PageSize int) ([]*model.Video, error) {
	ids, err := s.rdb.GetHotVideos(ctx, Page, PageSize)
	if err != nil {
		return nil, err
	}
	return s.videoRepo.GetVideosByIDs(ctx, ids)
}

func (s *VideoSvc) SearchVideos(ctx context.Context, keywords []string, Page, Pagesize int) ([]*model.Video, error) {
	s.rdb.IncreaseSearchCount(ctx, keywords)
	ids, err := s.rdb.GetVideoIDsByKeywords(ctx, keywords, Page, Pagesize)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		res, err := s.videoRepo.GetVideosByKeywords(ctx, keywords, Page, Pagesize)
		if err != nil {
			return nil, err
		}
		if len(res) > 0 {
			videoIDs := make([]int64, 0, len(res))
			for _, video := range res {
				videoIDs = append(videoIDs, video.ID)
			}
			s.rdb.CacheVideoIDsByKeywords(ctx, keywords, videoIDs)
		}
		return res, nil
	}
	return s.videoRepo.GetVideosByIDs(ctx, ids)
}
