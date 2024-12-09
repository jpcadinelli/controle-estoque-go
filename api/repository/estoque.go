package repository

import (
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EstoqueRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models.Estoque, error)
	FindByIdProduto(idProduto uuid.UUID, preloads ...string) (*models.Estoque, error)
	Create(estoque *models.Estoque) error
	Update(estoque *models.Estoque, updateItems map[string]interface{}) (*models.Estoque, error)
}

type estoqueRepositoryImpl struct {
	db *gorm.DB
}

func NewEstoqueRepository(db *gorm.DB) EstoqueRepository {
	return &estoqueRepositoryImpl{db: db}
}

func (r *estoqueRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models.Estoque, error) {
	var estoque models.Estoque

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&estoque, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrEstoqueNaoEncontrado
	}

	return &estoque, nil
}

func (r *estoqueRepositoryImpl) FindByIdProduto(idProduto uuid.UUID, preloads ...string) (*models.Estoque, error) {
	var estoque *models.Estoque

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&estoque, "id_produto = ? AND quantidade > 0", idProduto)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return estoque, nil
}

func (r *estoqueRepositoryImpl) Create(estoque *models.Estoque) error {
	return r.db.Create(estoque).Error
}

func (r *estoqueRepositoryImpl) Update(estoque *models.Estoque, updateItems map[string]interface{}) (*models.Estoque, error) {
	tx := r.db.Model(estoque).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrEstoqueNaoEncontrado
	}

	return estoque, nil
}
