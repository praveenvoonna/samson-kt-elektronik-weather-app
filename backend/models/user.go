package models

import (
	"strings"
	"time"

	"go.uber.org/zap"
)

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DateOfBirth Date   `json:"date_of_birth"`
}

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte, logger *zap.Logger) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}
