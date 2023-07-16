package model

import "errors"

var (
	ErrCandidateIsEmpty      = errors.New("Terjadi kesalahan inputan request candidateId kosong")
	ErrSwipeDirectionIsEmpty = errors.New("Terjadi kesalahan inputan request swipe direction kosong")
	ErrLimitIsEmpty          = errors.New("Terjadi kesalahan inputan request limit kosong")
)

type CandidateListPaginationReponse struct {
	CandidateList []CandidateListReponse `json:"candidationList"`
	Page          int                    `json:"page"`
	TotalPage     int                    `json:"totalPage"`
}

type CandidateListReponse struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PicUrl    string `json:"picUrl"`
	District  string `json:"district"`
	City      string `json:"city"`
}
