package models

import (
	"errors"
	"regexp"
	"time"

	"../helpers"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:not null`
	Email    string `gorm:"unique"`
	Password string
}

type Article struct {
	gorm.Model
	Title       string
	Content     string
	Text        string
	ImageURL    string
	IsPublish   bool
	PublishTime time.Time
	UserID      uint
	User        User
}

func (a *Article) FormattedDate() string {
	return a.PublishTime.Format("January 02, 2006")
}

func (u *User) IsValid() error {
	if u.Name == "" {
		return errors.New("Name cannot be empty")
	}

	match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", u.Email)
	if !match {
		return errors.New("Invalid Email")
	}

	if len(u.Password) < 6 {
		return errors.New("Password minimum length is 6 character ")
	}

	return nil

}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if err = u.IsValid(); err == nil {
		hashedPassword, err := helpers.HashPassword(u.Password)

		u.Password = string(hashedPassword)
		return err
	}

	return nil
}
