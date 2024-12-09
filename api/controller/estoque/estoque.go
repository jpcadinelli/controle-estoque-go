package estoque

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/global/enum"
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"api_pattern_go/api/repository"
	"api_pattern_go/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoEstoqueCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var e models.Estoque
	if err = ginctx.ShouldBindJSON(&e); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	eOld, err := repository.NewEstoqueRepository(dbConetion.DB).FindByIdProduto(e.IdProduto)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	if eOld == nil {
		if err = repository.NewEstoqueRepository(dbConetion.DB).Create(&e); err != nil {
			ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
			return
		}
	} else {
		updateItems := map[string]interface{}{
			"quantidade": e.Quantidade + eOld.Quantidade,
			"custo":      ((e.Custo) + (eOld.Custo * float64(eOld.Quantidade))) / float64(e.Quantidade+eOld.Quantidade),
		}

		eOld, err = repository.NewEstoqueRepository(dbConetion.DB).Update(eOld, updateItems)
		if err != nil {
			ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
			return
		}
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, e))
}

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoEstoqueAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var e models.Estoque
	if err = ginctx.ShouldBindJSON(&e); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	eOld, err := repository.NewEstoqueRepository(dbConetion.DB).FindById(e.Id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	updateItems := map[string]interface{}{
		"quantidade": e.Quantidade,
		"custo":      e.Custo / float64(e.Quantidade),
	}

	eOld, err = repository.NewEstoqueRepository(dbConetion.DB).Update(eOld, updateItems)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, eOld))
}
