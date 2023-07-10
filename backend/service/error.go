package service

import "errors"

var ErrRoleDoesNotExist = errors.New("role does not exist")
var ErrPermissionDenied = errors.New("permission denied")
var ErrInvalidContextUserData = errors.New("invalid context user data")
