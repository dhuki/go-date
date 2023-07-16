package domain

import (
	"errors"
	"strings"
)

var (
	swapDirectionLeft  = "Left"
	swapDirectionRight = "Right"

	transalteDirectionLeft  = "DISLIKE"
	transalteDirectionRight = "LIKE"

	KeyLastPagination = "last.page.candidate"

	ErrInvalidSwipeAction = errors.New("terjadi kesalahan swipe direction tidak valid")
)

func TranslateSwipeAction(swipeDirection string) (string, error) {
	switch swipeDirection {
	case strings.ToLower(swapDirectionLeft):
		return transalteDirectionLeft, nil
	case strings.ToLower(swapDirectionRight):
		return transalteDirectionRight, nil
	default:
		return "", ErrInvalidSwipeAction
	}

}
