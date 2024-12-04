package routes

import (
	"api_pattern_go/api/controller/estoque"
	"github.com/gin-gonic/gin"
)

func estoqueRoutes(r *gin.RouterGroup) {
	r.POST(route, estoque.Criar)
	r.PUT(route, estoque.Atualizar)
}
