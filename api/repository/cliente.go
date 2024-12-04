package repository

import (
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type ClienteRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models.Cliente, error)
	FindWithFilter(filtro models.ClienteFiltro, preloads ...string) ([]models.Cliente, error)
	FindAll(preloads ...string) ([]models.Cliente, error)
	Create(cliente *models.Cliente) error
	Update(cliente *models.Cliente, updateItems map[string]interface{}) (*models.Cliente, error)
	Delete(id uuid.UUID) error
}

type clienteRepositoryImpl struct {
	db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) ClienteRepository {
	return &clienteRepositoryImpl{db: db}
}

func (r *clienteRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models.Cliente, error) {
	var cliente models.Cliente

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&cliente, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrClienteNaoEncontrado
	}

	return &cliente, nil
}

func (r *clienteRepositoryImpl) FindWithFilter(filtro models.ClienteFiltro, preloads ...string) ([]models.Cliente, error) {
	var clientes []models.Cliente

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = prepereFilterCliente(filtro, tx)

	tx = tx.Find(&clientes)
	if tx.Error != nil {
		return clientes, tx.Error
	}
	if tx.RowsAffected == 0 {
		return clientes, erros.ErrClienteNaoEncontrado
	}

	return clientes, nil
}

func (r *clienteRepositoryImpl) FindAll(preloads ...string) ([]models.Cliente, error) {
	var clientes []models.Cliente

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&clientes)
	if tx.Error != nil {
		return clientes, tx.Error
	}
	if tx.RowsAffected == 0 {
		return clientes, erros.ErrClienteNaoEncontrado
	}

	return clientes, nil
}

func (r *clienteRepositoryImpl) Create(cliente *models.Cliente) error {
	return r.db.Create(cliente).Error
}

func (r *clienteRepositoryImpl) Update(cliente *models.Cliente, updateItems map[string]interface{}) (*models.Cliente, error) {
	tx := r.db.Model(cliente).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrClienteNaoEncontrado
	}

	return cliente, nil
}

func (r *clienteRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Cliente{}, "id = ?", id).Error
}

func prepereFilterCliente(filtro models.ClienteFiltro, tx *gorm.DB) *gorm.DB {
	if filtro.IdEnderecoPadrao != nil {
		tx = tx.Where("id_endereco_padrao = ?", *filtro.IdEnderecoPadrao)
	}

	if filtro.Nome != nil {
		tx = tx.Where("LOWER(nome) LIKE ?", "%"+strings.ToLower(*filtro.Nome)+"%")
	}

	if filtro.Referencia != nil {
		tx = tx.Where("LOWER(referencia) LIKE ?", "%"+strings.ToLower(*filtro.Referencia)+"%")
	}

	if filtro.Telefone != nil {
		tx = tx.Where("LOWER(telefone) LIKE ?", "%"+strings.ToLower(*filtro.Telefone)+"%")
	}

	if filtro.Whatsapp != nil {
		tx = tx.Where("LOWER(whatsapp) LIKE ?", "%"+strings.ToLower(*filtro.Whatsapp)+"%")
	}

	if filtro.Instagram != nil {
		tx = tx.Where("LOWER(instagram) LIKE ?", "%"+strings.ToLower(*filtro.Instagram)+"%")
	}

	return tx
}
