package user

import "gorm.io/gorm"

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo()(*UserRepo){
	return 
}