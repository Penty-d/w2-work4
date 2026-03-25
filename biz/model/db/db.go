package dbmodel

import "time"

type User struct {
	ID           int64      `gorm:"primaryKey;autoIncrement"`
	Username     string     `gorm:"type:varchar(64);not null;uniqueIndex"`
	PasswordHash string     `gorm:"type:text;not null"`
	AvatarURL    string     `gorm:"type:text;not null;default:''"`
	MFASecret    *string    `gorm:"type:text"`
	MFAEnabled   bool       `gorm:"not null;default:false"`
	CreatedAt    time.Time  `gorm:"not null;default:now()"`
	UpdatedAt    time.Time  `gorm:"not null;default:now()"`
	DeletedAt    *time.Time `gorm:"index"`
}

type Video struct {
	ID           int64      `gorm:"primaryKey;autoIncrement"`
	UserID       int64      `gorm:"not null;index"`
	VideoURL     string     `gorm:"type:text;not null"`
	CoverURL     string     `gorm:"type:text;not null"`
	Title        string     `gorm:"type:varchar(255);not null"`
	Description  string     `gorm:"type:text;not null;default:''"`
	VisitCount   int64      `gorm:"not null;default:0"`
	LikeCount    int64      `gorm:"not null;default:0"`
	CommentCount int64      `gorm:"not null;default:0"`
	CreatedAt    time.Time  `gorm:"not null;default:now();index"`
	UpdatedAt    time.Time  `gorm:"not null;default:now()"`
	DeletedAt    *time.Time `gorm:"index"`
}

type Comment struct {
	ID         int64      `gorm:"primaryKey;autoIncrement"`
	UserID     int64      `gorm:"not null;index"`
	VideoID    int64      `gorm:"not null;index"`
	ParentID   *int64     `gorm:"index"`
	Content    string     `gorm:"type:text;not null"`
	LikeCount  int64      `gorm:"not null;default:0"`
	ChildCount int64      `gorm:"not null;default:0"`
	CreatedAt  time.Time  `gorm:"not null;default:now();index"`
	UpdatedAt  time.Time  `gorm:"not null;default:now()"`
	DeletedAt  *time.Time `gorm:"index"`
}

type VideoLike struct {
	UserID    int64     `gorm:"primaryKey"`
	VideoID   int64     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

type Follow struct {
	UserID       int64     `gorm:"primaryKey"`
	FollowUserID int64     `gorm:"primaryKey"`
	CreatedAt    time.Time `gorm:"not null;default:now()"`
}

type RefreshToken struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	UserID    int64     `gorm:"not null;index"`
	TokenHash string    `gorm:"type:text;not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"not null;index"`
	Revoked   bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}
