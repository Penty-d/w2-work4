package video

import (
	"context"
	"strconv"

	model "w2-work4/biz/model/db"
	video "w2-work4/biz/repository/video"

	"github.com/go-redis/redis/v8"
)

const (
	HotVideosKey = "hot_videos:"
)

type VideoSvc struct {
	videoRepo *video.VideoRepo
	rdb       *redis.Client
} //

func NewVideoSvc(videorepo *video.VideoRepo, rdb *redis.Client) *VideoSvc {
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

func (s *VideoSvc) GetVideoByID(ctx context.Context, id int64) (*model.Video, error) {
	return s.videoRepo.GetVideoByID(ctx, id)
}

func (s *VideoSvc) GetVideosByIDs(ctx context.Context, ids []int64) ([]*model.Video, error) {
	return s.videoRepo.GetVideosByIDs(ctx, ids)
}

func (s *VideoSvc) GetVideosByUserID(ctx context.Context, userID int64) ([]*model.Video, error) {
	return s.videoRepo.GetVideosByUserID(ctx, userID)
}

func (s *VideoSvc) GetHotVideos(ctx context.Context, limit int) ([]*model.Video, error) {
	hotVideoIDs, err := s.rdb.ZRevRange(ctx, "hot_videos", 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}
	var hotVideos []*model.Video
	for _, idStr := range hotVideoIDs {
		videoID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			continue
		}
		video, err := s.GetVideoByID(ctx, videoID)
		if err == nil {
			hotVideos = append(hotVideos, video)
		}
	}
	return hotVideos, nil
}

/*func (s *VideoSvc) GetVideosByKeywords(ctx context.Context, keywords []string) ([]*model.Video, error) {
	if len(keywords) == 0 {
		return []*model.Video{}, nil
	}
	formattedKeywords := kw.FormatKeywords(keywords)
	cached, err := s.rdb.Get(ctx,HotVideosKey + formattedKeywords).Result()
	if err == nil {
		videoIDs := strings.Split(cached, "|")
		videoIDsInt := make([]int64, 0, len(videoIDs))
		for _, idStr := range videoIDs {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err == nil {
				videoIDsInt = append(videoIDsInt, id)
			}
		}
		var videos []*model.Video
		videos, err1 := s.videoRepo.GetVideosByIDs(ctx, videoIDsInt)
		if err1 == nil {
			return videos, nil
		}

	} else if err != redis.Nil {
		return nil, err
	}
	videos, err := s.videoRepo.GetVideosByKeywords(ctx, keywords)
	if err != nil {
		return nil, err
	}
	var videoIDs []string
	for _, video := range videos {
		videoIDs = append(videoIDs, strconv.FormatInt(video.ID, 10))
	}
	s.rdb.Set(ctx, HotVideosKey + formattedKeywords, strings.Join(videoIDs, "|"), 0)
	return videos, nil
}
*/
