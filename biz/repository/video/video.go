package video

import (
	"context"
	"strconv"
	model "w2-work4/biz/model/db"

	"github.com/go-redis/redis/v8"

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
	return &video, nil
}

func (r *VideoRepo) GetVideosByIDs(ctx context.Context, ids []int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.db.WithContext(ctx).Model(&model.Video{}).Where("id IN ?", ids).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, err
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

type VideoCacheRepo struct {
	rdb *redis.Client
}

const (
	VideoLikeCountKeyPrefix  = "video_like_count:"
	VideoVisitCountKeyPrefix = "video_visit_count:"
)

func NewVideoCacheRepo(rdb *redis.Client) *VideoCacheRepo {
	return &VideoCacheRepo{rdb: rdb}
}

func (r *VideoCacheRepo) GetVideoLikeCount(ctx context.Context, videoID int64) (int64, error) {
	key := VideoLikeCountKeyPrefix + strconv.FormatInt(videoID, 10)
	countStr, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil //没有缓存，默认0
		}
		return 0, err
	}
	return strconv.ParseInt(countStr, 10, 64)
}

func (r *VideoCacheRepo) IncrementVideoLikeCount(ctx context.Context, videoID int64) error {
	key := VideoLikeCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Incr(ctx, key).Err()
}

func (r *VideoCacheRepo) DecrementVideoLikeCount(ctx context.Context, videoID int64) error {
	key := VideoLikeCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Decr(ctx, key).Err()
}

func (r *VideoCacheRepo) GetVideoVisitCount(ctx context.Context, videoID int64) (int64, error) {
	key := VideoVisitCountKeyPrefix + strconv.FormatInt(videoID, 10)
	countStr, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil //没有缓存，默认0
		}
		return 0, err
	}
	return strconv.ParseInt(countStr, 10, 64)
}

func (r *VideoCacheRepo) IncrementVideoVisitCount(ctx context.Context, videoID int64) error {
	key := VideoVisitCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Incr(ctx, key).Err()
}

// 删除视频用得到
func (r *VideoCacheRepo) DeleteVideoVisitCount(ctx context.Context, videoID int64) error {
	key := VideoVisitCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Del(ctx, key).Err()
}

func (r *VideoCacheRepo) DeleteVideoLikeCount(ctx context.Context, videoID int64) error {
	key := VideoLikeCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Del(ctx, key).Err()
}
