// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int32
	Username  string
	Password  string
	Email     string
	CreatedAt pgtype.Timestamp
	LastLogin pgtype.Timestamp
}

type UserProfile struct {
	ID          int32
	UserID      int32
	PhoneNumber string
	Address     string
	IsVerified  bool
}