package repository

import (
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProdutoRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models.Produto, error)
	Create(produto *models.Produto) error
	Update(produto *models.Produto, updateItems map[string]interface{}) (*models.Produto, error)
	Delete(id uuid.UUID) error
}

type produtoRepositoryImpl struct {
	db *gorm.DB
}

func NewProdutoRepository(db *gorm.DB) ProdutoRepository {
	return &produtoRepositoryImpl{db: db}
}

func (r *produtoRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models.Produto, error) {
	var produto models.Produto

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&produto, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrProdutoNaoEncontrado
	}

	return &produto, nil
}

func (r *produtoRepositoryImpl) Create(produto *models.Produto) error {
	return r.db.Create(produto).Error
}

func (r *produtoRepositoryImpl) Update(produto *models.Produto, updateItems map[string]interface{}) (*models.Produto, error) {
	tx := r.db.Model(produto).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrProdutoNaoEncontrado
	}

	return produto, nil
}

func (r *produtoRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Produto{}, "id = ?", id).Error
}
