package mock

import "errors"

var errTopicNotFound = errors.New("topic not found")

var errTopicAlreadyExists = errors.New("topic already exist")