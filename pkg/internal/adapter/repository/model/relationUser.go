package model

import "time"

type RelationUser struct {
	ID           uint64     `db:"id"`
	UserID       uint64     `db:"userId"`
	CandidateID  uint64     `db:"candidateId"`
	RelationType string     `db:"relationType"`
	CreatedAt    time.Time  `db:"createdAt"`
	UpdatedAt    time.Time  `db:"updatedAt"`
	DeletedAt    *time.Time `db:"deletedAt"`
}
