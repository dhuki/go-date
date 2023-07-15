package domain

import (
	"errors"
	"strings"
)

var (
	swapDirectionLeft  = "Left"
	swapDirectionRight = "Right"

	transalteDirectionLeft  = "Dislike"
	transalteDirectionRight = "Like"

	KeyLastPagination = "last.page.candidate"
)

func TranslateSwipeAction(swipeDirection string) (string, error) {
	switch swipeDirection {
	case strings.ToLower(swapDirectionLeft):
		return transalteDirectionLeft, nil
	case strings.ToLower(swapDirectionRight):
		return transalteDirectionRight, nil
	default:
		return "", errors.New("erro happen") // todo
	}

}
