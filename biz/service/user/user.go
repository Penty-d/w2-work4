package user

import (
	"context"
	"errors"
	model "w2-work4/biz/model/db"
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
	_, err := s.userRepo.GetUserByUsername(ctx, username)
	if err == nil {
		return errors.New("username already exists")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	passwordHash, err := passwordutil.HashPassword(password)
	if err != nil {
		return err
	}
	user := &model.User{
		Username:     username,
		PasswordHash: passwordHash,
	}
	return s.userRepo.AddUser(ctx, user)
}

func (s *UserSvc) UserLogin(ctx context.Context, username, password string) (*model.User, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if !passwordutil.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

func (s *UserSvc) UpdateAvatar(ctx context.Context, avatarurl string) error {
	return s.userRepo.UpdateUserAvatar(ctx, avatarurl)
}

func (s *UserSvc) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *UserSvc) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.userRepo.GetUserByUsername(ctx, username)
}
