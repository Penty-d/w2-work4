package like

import (
	"context"
	"strconv"
	"time"
	model "w2-work4/biz/model/db"

	"github.com/go-redis/redis/v8"
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

func (r *LikeRepo) GetUserIDsByVideoID(ctx context.Context, videoID int64) ([]int64, error) {
	var userIDs []int64
	err := r.db.WithContext(ctx).Model(&model.VideoLike{}).Where("video_id = ?", videoID).Pluck("user_id", &userIDs).Error
	if err != nil {
		return nil, err
	}
	return userIDs, nil
}

type LikeCacheRepo struct {
	rdb *redis.Client
}

func NewLikeCacheRepo(rdb *redis.Client) *LikeCacheRepo {
	return &LikeCacheRepo{rdb: rdb}
}

const (
	VideoLikeCountKeyPrefix = "video_like_count:"
	VideoLikeCountKeyTTL    = 24 * time.Hour
)

func (r *LikeCacheRepo) IncreaseVideoLikeCount(ctx context.Context, videoID int64) error {
	key := VideoLikeCountKeyPrefix + strconv.FormatInt(videoID, 10)
	pipe := r.rdb.TxPipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, VideoLikeCountKeyTTL)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *LikeCacheRepo) DecreaseVideoLikeCount(ctx context.Context, videoID int64) error {
	key := VideoLikeCountKeyPrefix + strconv.FormatInt(videoID, 10)
	pipe := r.rdb.TxPipeline()
	pipe.Decr(ctx, key)
	pipe.Expire(ctx, key, VideoLikeCountKeyTTL)
	_, err := pipe.Exec(ctx)
	return err
}
