package repository

import (
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EnderecoRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models.Endereco, error)
	Create(endereco *models.Endereco) error
	Update(endereco *models.Endereco, updateItems map[string]interface{}) (*models.Endereco, error)
}

type enderecoRepositoryImpl struct {
	db *gorm.DB
}

func NewEnderecoRepository(db *gorm.DB) EnderecoRepository {
	return &enderecoRepositoryImpl{db: db}
}

func (r *enderecoRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models.Endereco, error) {
	var endereco models.Endereco

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&endereco, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrEnderecoNaoEncontrado
	}

	return &endereco, nil
}

func (r *enderecoRepositoryImpl) Create(endereco *models.Endereco) error {
	return r.db.Create(endereco).Error
}

func (r *enderecoRepositoryImpl) Update(endereco *models.Endereco, updateItems map[string]interface{}) (*models.Endereco, error) {
	tx := r.db.Model(endereco).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrEnderecoNaoEncontrado
	}

	return endereco, nil
}
