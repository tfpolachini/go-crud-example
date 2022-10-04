package service

import "errors"

var ErrInvalidAction = errors.New("invalid action for this status")
var ErrResourceAlreadyExists = errors.New("resource already exists")
var ErrResourceNotFound = errors.New("resource not found")
