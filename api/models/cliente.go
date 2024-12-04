package models

import (
	"api_pattern_go/api/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cliente struct {
	Id               uuid.UUID `json:"id"`
	IdEnderecoPadrao uuid.UUID `json:"idEnderecoPadrao"`
	Nome             string    `json:"nome"`
	Referencia       string    `json:"referencia"`
	Telefone         string    `json:"telefone"`
	Whatsapp         string    `json:"whatsapp"`
	Instagram        string    `json:"instagram"`
}

func (c *Cliente) BeforeCreate(_ *gorm.DB) (err error) {
	c.Id = uuid.New()
	return err
}

func (c *Cliente) TableName() string {
	return global.TableCliente
}

func (c *Cliente) ClienteToDropdownUUID() *DropdownUUID {
	return &DropdownUUID{
		Label: c.Nome,
		Value: c.Id,
	}
}

type ClienteFiltro struct {
	IdEnderecoPadrao *uuid.UUID `json:"idEnderecoPadrao"`
	Nome             *string    `json:"nome"`
	Referencia       *string    `json:"referencia"`
	Telefone         *string    `json:"telefone"`
	Whatsapp         *string    `json:"whatsapp"`
	Instagram        *string    `json:"instagram"`
}
