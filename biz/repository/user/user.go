package user

import (
	"context"
	model "w2-work4/biz/model/db"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) AddUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) UpdateUserAvatar(ctx context.Context, avatarurl string) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Select("avatar_url").Updates(map[string]string{"avatar_url": avatarurl}).Error
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetLikeByUserID(ctx context.Context, userID int64) ([]*model.VideoLike, error) {
	var likes []*model.VideoLike
	err := r.db.WithContext(ctx).Model(&model.VideoLike{}).Where("user_id = ?", userID).Find(&likes).Error
	if err != nil {
		return nil, err
	}
	return likes, nil
}
