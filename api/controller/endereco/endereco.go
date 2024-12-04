package endereco

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

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoEnderecoCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var e models.Endereco
	if err = ginctx.ShouldBindJSON(&e); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewEnderecoRepository(dbConetion.DB).Create(&e); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, e))
}

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoEnderecoAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var e models.Endereco
	if err = ginctx.ShouldBindJSON(&e); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	eOld, err := repository.NewEnderecoRepository(dbConetion.DB).FindById(e.Id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	updateItems := map[string]interface{}{
		"endereco": e.Endereco,
	}

	eOld, err = repository.NewEnderecoRepository(dbConetion.DB).Update(eOld, updateItems)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, eOld))
}
