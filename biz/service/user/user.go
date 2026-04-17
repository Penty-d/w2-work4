package user

import (
	"context"
	"errors"
	"fmt"
	model "w2-work4/biz/model/db"
	errme "w2-work4/biz/model/errormessage"
	user "w2-work4/biz/repository/user"
	passwordutil "w2-work4/internal/utils/password"

	"gorm.io/gorm"
)

type UserSvc struct {
	userRepo *user.UserRepo
}

func NewUserSvc(userrepo *user.UserRepo) *UserSvc {
	return &UserSvc{userRepo: userrepo}
}

func (s *UserSvc) UserRegister(ctx context.Context, username, password string) error {
	if username == "" || password == "" {
		return errme.ErrMissingRequiredFields
	}
	_, err := s.userRepo.GetUserByUsername(ctx, username)
	if err == nil {
		return errors.New("username already exists")
	}
	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check username existence: %w", err)
	}
	passwordHash, err := passwordutil.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user := &model.User{
		Username:     username,
		PasswordHash: passwordHash,
	}
	if err := s.userRepo.AddUser(ctx, user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	return nil
}

func (s *UserSvc) UserLogin(ctx context.Context, username, password string) (*model.User, error) {
	if username == "" || password == "" {
		return nil, errme.ErrMissingRequiredFields
	}
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	if !passwordutil.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

func (s *UserSvc) UpdateAvatar(ctx context.Context, avatarurl string) error {
	if avatarurl == "" {
		return errme.ErrMissingRequiredFields
	}
	if err := s.userRepo.UpdateUserAvatar(ctx, avatarurl); err != nil {
		return fmt.Errorf("failed to update user avatar: %w", err)
	}
	return nil
}

func (s *UserSvc) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	if id <= 0 {
		return nil, errme.ErrInvalidUserID
	}
	user, err := s.userRepo.GetUserByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errme.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return user, nil
}

func (s *UserSvc) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, errme.ErrMissingRequiredFields
	}
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errme.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return user, nil
}

func (s *UserSvc) GetLikedVideosByUserID(ctx context.Context, userID int64) ([]*model.VideoLike, error) {
	if userID <= 0 {
		return nil, errme.ErrInvalidUserID
	}
	likes, err := s.userRepo.GetLikeByUserID(ctx, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errme.ErrLikedVideosNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get liked videos: %w", err)
	}
	return likes, nil
}
