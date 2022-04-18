package models

import "time"

type BaseModel struct {
	CreatedAt  time.Time
	ModifiedAt time.Time
	ModifiedBy int
	CreatedBy  int
}

func NewBaseModel() BaseModel {
	return BaseModel{
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
}
