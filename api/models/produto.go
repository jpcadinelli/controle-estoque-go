package models

import (
	"api_pattern_go/api/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Produto struct {
	Id         uuid.UUID `json:"id"`
	Nome       string    `json:"nome"`
	Marca      string    `json:"marca"`
	Quantidade int       `json:"quantidade"`
	Unidade    string    `json:"unidade"`
}

func (p *Produto) BeforeCreate(_ *gorm.DB) (err error) {
	p.Id = uuid.New()
	return err
}

func (p *Produto) TableName() string {
	return global.TableProduto
}

func (p *Produto) ProdutoToDropdownUUID() *DropdownUUID {
	return &DropdownUUID{
		Label: p.Nome,
		Value: p.Id,
	}
}
