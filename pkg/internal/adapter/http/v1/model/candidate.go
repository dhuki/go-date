package model

type CandidateListPaginationReponse struct {
	CandidateList []CandidateListReponse `json:"candidationList"`
	Page          int                    `json:"page"`
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
