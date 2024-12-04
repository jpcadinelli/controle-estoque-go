package routes

import (
	"api_pattern_go/api/controller/produto"
	"github.com/gin-gonic/gin"
)

func produtoRoutes(r *gin.RouterGroup) {
	r.POST(route, produto.Criar)
	r.GET(routeId, produto.Visualizar)
	r.GET(routeDropdown, produto.Dropdown)
	r.PUT(route, produto.Atualizar)
	r.DELETE(routeId, produto.Deletar)
}
