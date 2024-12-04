package models

import (
	"api_pattern_go/api/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Estoque struct {
	Id         uuid.UUID `json:"id"`
	IdProduto  uuid.UUID `json:"idProduto"`
	Quantidade int       `json:"quantidade"`
	Custo      float64   `json:"custo"`
	Data       time.Time `json:"data"`
}

func (e *Estoque) BeforeCreate(_ *gorm.DB) (err error) {
	e.Id = uuid.New()
	e.Custo = e.Custo / float64(e.Quantidade)
	e.Data = time.Now()
	return err
}

func (e *Estoque) TableName() string {
	return global.TableEstoque
}
