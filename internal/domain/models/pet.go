package models

type Pet struct {
	ID     uint64 `json:"id"`
	SlotID uint64 `json:"slot_id"`
	UserID uint64 `json:"user_id"`
}
