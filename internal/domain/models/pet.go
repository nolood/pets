package models

const (
	PetRarityCommon = iota + 1
	PetRarityUncommon
	PetRarityRare
	PetRarityEpic
	PetRarityLegendary
)

type Pet struct {
	ID     uint64 `json:"id"`
	Rarity uint   `json:"rarity"`
	Image  string `json:"image"`
	SlotID uint64 `json:"slot_id"`
	UserID uint64 `json:"user_id"`
}
