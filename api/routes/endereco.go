package routes

import (
	"api_pattern_go/api/controller/endereco"
	"github.com/gin-gonic/gin"
)

func enderecoRoutes(r *gin.RouterGroup) {
	r.POST(route, endereco.Criar)
	r.PUT(route, endereco.Atualizar)
}
