package storage

import "database/sql"

type UserEgg struct {
	ID         uint64       `db:"id"`
	UserID     uint64       `db:"user_id"`
	EggID      uint64       `db:"egg_id"`
	HatchTime  sql.NullTime `db:"hatch_time"`
	HatchStart sql.NullTime `db:"hatch_start"`
	HatchEnd   sql.NullTime `db:"hatch_end"`
}

type UserPet struct {
	ID     uint64 `db:"id"`
	UserID uint64 `db:"user_id"`
	PetID  uint64 `db:"pet_id"`
	Level  uint   `db:"level"`
}

type Egg struct {
	ID     uint64 `db:"id"`
	Rarity uint   `db:"rarity"`
	Image  string `db:"image"`
}

type Farm struct {
	ID     uint64 `db:"id"`
	UserID uint64 `db:"user_id"`
}

type Pet struct {
	ID     uint64 `db:"id"`
	Rarity uint   `db:"rarity"`
	Image  string `db:"image"`
}

type Slot struct {
	ID     uint64 `db:"id"`
	FarmID uint64 `db:"farm_id"`
	PetID  uint64 `db:"pet_id"`
	Charge int    `db:"charge"`
	Index  int    `db:"index"`
}

type Incubator struct {
	ID    uint64 `db:"id"`
	User  uint64 `db:"user_id"`
	EggID uint64 `db:"egg_id"`
}

type User struct {
	ID           uint64         `db:"id"`
	TgID         uint64         `db:"tg_id"`
	Username     string         `db:"username"`
	Lastname     sql.NullString `db:"lastname"`
	Firstname    sql.NullString `db:"firstname"`
	LanguageCode sql.NullString `db:"language_code"`
	IsPremium    sql.NullString `db:"is_premium"`
	PhotoUrl     sql.NullString `db:"photo_url"`
}
