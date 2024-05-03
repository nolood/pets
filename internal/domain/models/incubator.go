package models

import "time"

type Egg struct {
	ID     uint64 `json:"id"`
	Rarity uint   `json:"rarity"`
	Image  string `json:"image"`
}

type UserEgg struct {
	ID         uint64     `json:"id"`
	UserID     uint64     `json:"user_id"`
	Egg        Egg        `json:"egg"`
	HatchTime  time.Time  `json:"hatch_time"`
	HatchStart *time.Time `json:"hatch_start"`
	HatchEnd   *time.Time `json:"hatch_end"`
}

type Incubator struct {
	ID     uint64   `json:"id"`
	UserID uint64   `json:"user_id"`
	Egg    *UserEgg `json:"egg"`
}
