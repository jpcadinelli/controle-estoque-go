package routes

import (
	"api_pattern_go/api/controller/cliente"
	"github.com/gin-gonic/gin"
)

func clienteRoutes(r *gin.RouterGroup) {
	r.POST(route, cliente.Criar)
	r.GET(routeId, cliente.Visualizar)
	r.POST(routeFiltro, cliente.Listar)
	r.GET(routeDropdown, cliente.Dropdown)
	r.PUT(route, cliente.Atualizar)
	r.DELETE(routeId, cliente.Deletar)
}
