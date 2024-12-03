package models

import (
	"api_pattern_go/api/global"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Usuario struct {
	Id             uuid.UUID   `json:"id"`
	PrimeiroNome   string      `json:"primeiroNome" validate:"required"`
	UltimoNome     string      `json:"ultimoNome" validate:"required"`
	CPF            string      `json:"cpf" validate:"required"`
	Email          string      `json:"email" validate:"required"`
	Password       string      `json:"password" validate:"required"`
	DataNascimento time.Time   `json:"dataNascimento" validate:"required"`
	Permissoes     []Permissao `json:"permissoes" gorm:"many2many:permissao_usuario;joinForeignKey:IdUsuario;joinReferences:IdPermissao"`
	CreatedAt      time.Time   `json:"createdAt"`
}

func (u *Usuario) BeforeCreate(_ *gorm.DB) (err error) {
	u.Id = uuid.New()
	u.CreatedAt = time.Now()
	return err
}

func (u *Usuario) TableName() string {
	return global.TableUsuario
}

type UsuarioDTOResponse struct {
	Id             uuid.UUID   `json:"id"`
	PrimeiroNome   string      `json:"primeiroNome"`
	UltimoNome     string      `json:"ultimoNome"`
	CPF            string      `json:"cpf"`
	Email          string      `json:"email"`
	DataNascimento time.Time   `json:"dataNascimento"`
	Permissoes     []Permissao `json:"permissoes"`
	CreatedAt      time.Time   `json:"createdAt"`
}

func (u *Usuario) UsuarioToDTOResponse() *UsuarioDTOResponse {
	return &UsuarioDTOResponse{
		Id:             u.Id,
		PrimeiroNome:   u.PrimeiroNome,
		UltimoNome:     u.UltimoNome,
		CPF:            u.CPF,
		Email:          u.Email,
		DataNascimento: u.DataNascimento,
		Permissoes:     u.Permissoes,
		CreatedAt:      u.CreatedAt,
	}
}
func (u *Usuario) UsuarioToDropdownUUID() *DropdownUUID {
	return &DropdownUUID{
		Label: fmt.Sprintf("%v %v (%v)", u.PrimeiroNome, u.UltimoNome, u.Email),
		Value: u.Id,
	}
}
