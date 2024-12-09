package repository

import (
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/models"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VendaRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models.Venda, error)
	FindWithFilter(filtro models.VendaFiltro, preloads ...string) ([]models.Venda, error)
	FindAll(preloads ...string) ([]models.Venda, error)
	Create(venda *models.Venda) error
	Update(venda *models.Venda, updateItems map[string]interface{}) (*models.Venda, error)
	Delete(id uuid.UUID) error
}

type vendaRepositoryImpl struct {
	db *gorm.DB
}

func NewVendaRepository(db *gorm.DB) VendaRepository {
	return &vendaRepositoryImpl{db: db}
}

func (r *vendaRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models.Venda, error) {
	var venda models.Venda

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&venda, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrVendaNaoEncontrada
	}

	return &venda, nil
}

func (r *vendaRepositoryImpl) FindWithFilter(filtro models.VendaFiltro, preloads ...string) ([]models.Venda, error) {
	var vendas []models.Venda

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = prepereFilterVenda(filtro, tx)

	tx = tx.Find(&vendas)
	if tx.Error != nil {
		return vendas, tx.Error
	}
	if tx.RowsAffected == 0 {
		return vendas, erros.ErrVendaNaoEncontrada
	}

	return vendas, nil
}

func (r *vendaRepositoryImpl) FindAll(preloads ...string) ([]models.Venda, error) {
	var vendas []models.Venda

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&vendas)
	if tx.Error != nil {
		return vendas, tx.Error
	}
	if tx.RowsAffected == 0 {
		return vendas, erros.ErrVendaNaoEncontrada
	}

	return vendas, nil
}

func (r *vendaRepositoryImpl) Create(venda *models.Venda) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var custoTotalVenda float64

		for _, vp := range venda.Produtos {
			estoqueRepo := NewEstoqueRepository(r.db)

			estoque, err := estoqueRepo.FindByIdProduto(vp.IdProduto)
			if err != nil {
				return err
			}

			custoTotalVenda += float64(vp.Quantidade) * estoque.Custo

			novaQuantidade := estoque.Quantidade - vp.Quantidade
			if novaQuantidade < 0 {
				return fmt.Errorf(erros.ErrQuantidadeInsuficiente, vp.IdProduto, estoque.Quantidade)
			}

			updateItems := map[string]interface{}{
				"quantidade": novaQuantidade,
			}

			_, err = estoqueRepo.Update(estoque, updateItems)
			if err != nil {
				return err
			}
		}

		venda.Custo = custoTotalVenda

		if err := tx.Create(venda).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *vendaRepositoryImpl) Update(venda *models.Venda, updateItems map[string]interface{}) (*models.Venda, error) {
	tx := r.db.Model(venda).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrVendaNaoEncontrada
	}

	return venda, nil
}

func (r *vendaRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Venda{}, "id = ?", id).Error
}

func prepereFilterVenda(filtro models.VendaFiltro, tx *gorm.DB) *gorm.DB {
	//if filtro.IdEnderecoPadrao != nil {
	//	tx = tx.Where("id_endereco_padrao = ?", *filtro.IdEnderecoPadrao)
	//}

	//if filtro.Nome != nil {
	//	tx = tx.Where("LOWER(nome) LIKE ?", "%"+strings.ToLower(*filtro.Nome)+"%")
	//}

	//if filtro.Referencia != nil {
	//	tx = tx.Where("LOWER(referencia) LIKE ?", "%"+strings.ToLower(*filtro.Referencia)+"%")
	//}

	//if filtro.Telefone != nil {
	//	tx = tx.Where("LOWER(telefone) LIKE ?", "%"+strings.ToLower(*filtro.Telefone)+"%")
	//}

	//if filtro.Whatsapp != nil {
	//	tx = tx.Where("LOWER(whatsapp) LIKE ?", "%"+strings.ToLower(*filtro.Whatsapp)+"%")
	//}

	//if filtro.Instagram != nil {
	//	tx = tx.Where("LOWER(instagram) LIKE ?", "%"+strings.ToLower(*filtro.Instagram)+"%")
	//}

	return tx
}
