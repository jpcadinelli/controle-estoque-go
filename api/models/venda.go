package models

import (
	"api_pattern_go/api/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Venda struct {
	Id         uuid.UUID      `json:"id"`
	IdCliente  uuid.UUID      `json:"idCliente"`
	IdEndereco uuid.UUID      `json:"idEndereco"`
	Custo      float64        `json:"custo"`
	Entrega    float64        `json:"entrega"`
	Valor      float64        `json:"valor"`
	Pago       bool           `json:"pago"`
	Data       time.Time      `json:"data"`
	Produtos   []VendaProduto `json:"produtos" gorm:"foreignKey:IdVenda;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (v *Venda) BeforeCreate(_ *gorm.DB) (err error) {
	v.Id = uuid.New()
	v.Data = time.Now()
	return err
}

func (v *Venda) TableName() string {
	return global.TableVenda
}

type VendaProduto struct {
	Id         uuid.UUID `json:"id"`
	IdVenda    uuid.UUID `json:"idVenda" gorm:"column:id_venda"`
	IdProduto  uuid.UUID `json:"idProduto" gorm:"column:id_produto"`
	Quantidade int       `json:"quantidade"`
}

func (vp *VendaProduto) BeforeCreate(_ *gorm.DB) (err error) {
	vp.Id = uuid.New()
	return err
}

func (vp *VendaProduto) TableName() string {
	return global.TableVendaProduto
}

type VendaFiltro struct {
	IdCliente  *uuid.UUID `json:"idCliente"`
	IdEndereco *uuid.UUID `json:"idEndereco"`
	Custo      *float64   `json:"custo"`
	Entrega    *float64   `json:"entrega"`
	Valor      *float64   `json:"valor"`
	Pago       *bool      `json:"pago"`
	DataInicio *time.Time `json:"dataInicio"`
	DataFim    *time.Time `json:"dataFim"`
}
