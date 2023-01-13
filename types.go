package main

import (
	"time"
)

type Guild struct {
	Id         uint
	Discord_id string
	Name       string
	Created    time.Time
	Updated    time.Time
}

type User struct {
	Id         uint
	Discord_id string
	Nickname   string
	Created    time.Time
	Updated    time.Time
}
