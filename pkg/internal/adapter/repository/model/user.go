package model

import "time"

type User struct {
	ID          uint64     `db:"id"`
	Username    string     `db:"username"`
	FirstName   string     `db:"firstName"`
	LastName    string     `db:"lastName"`
	Email       string     `db:"email"`
	PicUrl      string     `db:"picUrl"`
	District    string     `db:"district"`
	City        string     `db:"city"`
	IsPremium   bool       `db:"isPremium"`
	IsValidated bool       `db:"isValidated"`
	CreatedAt   time.Time  `db:"createdAt"`
	UpdatedAt   time.Time  `db:"updatedAt"`
	DeletedAt   *time.Time `db:"deletedAt"`
}
