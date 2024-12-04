package models

import (
	"api_pattern_go/api/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Endereco struct {
	Id       uuid.UUID `json:"id"`
	Endereco string    `json:"endereco"`
}

func (e *Endereco) BeforeCreate(_ *gorm.DB) (err error) {
	e.Id = uuid.New()
	return err
}

func (e *Endereco) TableName() string {
	return global.TableEndereco
}
