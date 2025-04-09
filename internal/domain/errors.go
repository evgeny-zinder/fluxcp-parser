package domain

import "errors"

var (
	ErrNoSessionCookie = errors.New("no session cookie")
	ErrParsingFailed   = errors.New("parsing failed")
)
