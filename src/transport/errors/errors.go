package errors

import "errors"

var (
	ErrUidContextIssue = errors.New("there is an issue with UID in the request")
)
