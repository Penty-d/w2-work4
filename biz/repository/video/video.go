package video

import (
	"context"
	"strconv"
	"time"
	model "w2-work4/biz/model/db"
	kw "w2-work4/internal/utils/keywords"

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
	SearchResultsKeyPrefix   = "search_results:"
	SearchKeywordsKeyPrefix  = "search_keywords:"
	HotVideosKey             = "hot_videos"
)

func NewVideoCacheRepo(rdb *redis.Client) *VideoCacheRepo {
	return &VideoCacheRepo{rdb: rdb}
}

func (r *VideoCacheRepo) GetSearchResults(ctx context.Context, keywords []string) ([]int64, error) {
	kws := kw.FormatKeywords(keywords)
	ids, err := r.rdb.SMembers(ctx, SearchResultsKeyPrefix+kws).Result()
	if err != nil {
		return nil, err
	}
	var videoIDs []int64
	for _, id := range ids {
		videoID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}
		videoIDs = append(videoIDs, videoID)
	}
	return videoIDs, nil
}

func (r *VideoCacheRepo) CacheSearchResults(ctx context.Context, keywords []string, videoIDs []int64) error {
	kws := kw.FormatKeywords(keywords)
	members := make([]any, 0, len(videoIDs))
	for _, videoID := range videoIDs {
		members = append(members, videoID)
	}
	return r.rdb.SAdd(ctx, SearchResultsKeyPrefix+kws, members...).Err()
}

func (r *VideoCacheRepo) CacheHotVideos(ctx context.Context, videoIDs []int64) error {
	members := make([]any, 0, len(videoIDs))
	for _, videoID := range videoIDs {
		members = append(members, videoID)
	}
	return r.rdb.SAdd(ctx, HotVideosKey, members...).Err()
}

func (r *VideoCacheRepo) IncreaseSearchCount(ctx context.Context, keywords []string) error {
	kws := kw.FormatKeywords(keywords)
	return r.rdb.ZIncrBy(ctx, SearchKeywordsKeyPrefix+time.Now().Format("2006-01-02"), 1, kws).Err()
}

func (r *VideoCacheRepo) DeleteSearchCount(ctx context.Context, day time.Time) error {
	key := SearchKeywordsKeyPrefix + day.Format("2006-01-02")
	return r.rdb.Del(ctx, key).Err()
}

func (r *VideoCacheRepo) DeleteHotVideosCache(ctx context.Context) error {
	return r.rdb.Del(ctx, HotVideosKey).Err()
}

func (r *VideoCacheRepo) DeleteSearchResultsCache(ctx context.Context, keywords []string) error {
	kws := kw.FormatKeywords(keywords)
	return r.rdb.Del(ctx, SearchResultsKeyPrefix+kws).Err()
}

func (r *VideoCacheRepo) GetHotSearchKeywords(ctx context.Context, limit int) ([]string, error) {
	key := SearchKeywordsKeyPrefix + time.Now().Format("2006-01-02")
	result, err := r.rdb.ZRevRangeWithScores(ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}
	var keywords []string
	for _, item := range result {
		keywords = append(keywords, item.Member.(string))
	}
	return keywords, nil
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

func (r *VideoCacheRepo) IncreaseVideoLikeCount(ctx context.Context, videoID int64) error {
	key := VideoLikeCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Incr(ctx, key).Err()
}

func (r *VideoCacheRepo) DecreaseVideoLikeCount(ctx context.Context, videoID int64) error {
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

func (r *VideoCacheRepo) IncreaseVideoVisitCount(ctx context.Context, videoID int64) error {
	key := VideoVisitCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Incr(ctx, key).Err()
}

func (r *VideoCacheRepo) DeleteVideoVisitCount(ctx context.Context, videoID int64) error {
	key := VideoVisitCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Del(ctx, key).Err()
}

func (r *VideoCacheRepo) DeleteVideoLikeCount(ctx context.Context, videoID int64) error {
	key := VideoLikeCountKeyPrefix + strconv.FormatInt(videoID, 10)
	return r.rdb.Del(ctx, key).Err()
}
