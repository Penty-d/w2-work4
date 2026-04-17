package errormessage

import "errors"

var (
	ErrInvalidUserID         = errors.New("invalid userID")
	ErrInvalidVideoID        = errors.New("invalid video ID")
	ErrInvalidCommentID      = errors.New("invalid comment ID")
	ErrVideoNotFound         = errors.New("video not found")
	ErrLikedVideosNotFound   = errors.New("liked videos not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrCannotFollowSelf      = errors.New("cannot follow yourself")
	ErrInvalidPageOrPageSize = errors.New("invalid page or page size")
	ErrMissingRequiredFields = errors.New("missing required fields")
)
