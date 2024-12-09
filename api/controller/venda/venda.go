package venda

import (
	dbConetion "api_pattern_go/api/database/conection"
	"api_pattern_go/api/global/enum"
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/middleware"
	"api_pattern_go/api/models"
	"api_pattern_go/api/repository"
	"api_pattern_go/api/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoVendaCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var venda models.Venda
	if err = ginctx.ShouldBindJSON(&venda); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewVendaRepository(dbConetion.DB).Create(&venda); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, venda))
}

func Visualizar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoVendaVisualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	venda, err := repository.NewVendaRepository(dbConetion.DB).FindById(id, "Produtos")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, venda))
}

func Listar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoVendaListar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var filtro models.VendaFiltro
	if err = ginctx.ShouldBindJSON(&filtro); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	vendas, err := repository.NewVendaRepository(dbConetion.DB).FindWithFilter(filtro, "Produtos")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, vendas))
}
