package cerr

import "errors"

var (
	ErrUserAlreadyExists        = errors.New("user with this nickname already exists")
	ErrUserDoesntExist          = errors.New("user doesn't exist")
	ErrCantFindUser             = errors.New("can't find user with id #42")
	ErrNicknameParamNotProvided = errors.New("nickname is not provided")
	ErrEmailIsInUse             = errors.New("email is in use")
	ErrForumAlreadyExists       = errors.New("forum already exists")
	ErrForumDoesntExist         = errors.New("forum doesn't exist")
	ErrWrongSlugProvided        = errors.New("slug not provided")
	ErrThreadDoesntExist        = errors.New("thread doesn't exist")
	ErrThreadAlreadyExists      = errors.New("thread already exists")
)
