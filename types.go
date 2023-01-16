package main

import (
	"time"

	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Discord_ID string
	Name       string
	Created    time.Time
	Updated    time.Time
}

type User struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Discord_ID string
	Nickname   string
	Created    time.Time
	Updated    time.Time
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
	ID          uint `gorm:"primaryKey"`
	IMDB_ID     string
	Title       string
	RuntimeMin  uint
	Description string
	Released    time.Time
	Created     time.Time
}

type GuildMovie struct {
	gorm.Model
	GuildID     uint
	MovieID     uint
	Guild       Guild `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Movie       Movie `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	NumViewings uint
	UserID      uint
	SuggestedBy User `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Created     time.Time
	Updated     time.Time
}

type EventLoc struct {
	gorm.Model
	ID          uint  `gorm:"primaryKey"`
	Guild       Guild `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	IsVC        bool
	Description string
}

type Event struct {
	gorm.Model
	ID                 uint `gorm:"primaryKey"`
	GuildID            uint
	Guild              Guild `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Topic              string
	LocID              uint
	Location           EventLoc `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Description        string
	StartTime          time.Time
	EndTime            time.Time
	Repeating          bool
	RepeatIntervalDays uint
	Completed          bool
	Created            time.Time
	Updated            time.Time
}
