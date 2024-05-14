package models

type Slot struct {
	ID          uint64   `json:"id"`
	FarmID      uint64   `json:"farm_id"`
	PetID       *uint64  `json:"pet_id"`
	Index       int      `json:"index"`
	Charge      int      `json:"charge"`
	IsAvailable bool     `json:"is_available"`
	Price       int      `json:"price"`
	Pet         *UserPet `json:"pet"`
}

type Farm struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"user_id"`
	Slots  []Slot `json:"slots"`
}
