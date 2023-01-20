package main

import (
	"time"

	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	Discord_ID string
	Name       string
}

type User struct {
	gorm.Model
	Discord_ID string
	Nickname   string
	State      string
}

type GuildMember struct {
	gorm.Model
	UserID  uint
	GuildID uint
	User    User  `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Guild   Guild `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}

type Movie struct {
	gorm.Model
	IMDB_ID     string
	Title       string
	RuntimeMin  uint
	Description string
	Released    time.Time
}

type GuildMovie struct {
	gorm.Model
	GuildID     uint
	MovieID     uint
	Guild       Guild `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Movie       Movie `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	NumViewings uint
	UserID      uint
	SuggestedBy User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}

type EventLoc struct {
	gorm.Model
	GuildID     uint
	Guild       Guild `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	IsVC        bool
	Description string
}

type Event struct {
	gorm.Model
	GuildID            uint
	Guild              Guild `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Topic              string
	LocationID         uint
	Location           EventLoc `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Description        string
	StartTime          time.Time
	EndTime            time.Time
	Repeating          bool
	RepeatIntervalDays uint
	Completed          bool
}
