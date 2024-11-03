package domain_error

import "errors"

var (
	ErrProblemNotFound  = errors.New("проблема не найдена")
	ErrUserNotFound     = errors.New("пользователь не найден")
	ErrUserAlreadyExist = errors.New("пользователь с таким телефоном уже существует")
)
