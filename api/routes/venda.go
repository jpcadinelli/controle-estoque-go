package routes

import (
	"api_pattern_go/api/controller/venda"
	"github.com/gin-gonic/gin"
)

func vendaRoutes(r *gin.RouterGroup) {
	r.POST(route, venda.Criar)
	r.GET(routeId, venda.Visualizar)
	r.POST(routeFiltro, venda.Listar)
	r.PUT(routeId, venda.Atualizar)
	r.DELETE(routeId, venda.Deletar)
}
