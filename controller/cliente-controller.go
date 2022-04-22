package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/models"
    "github.com/mariarobertap/api-vidroglass/service"
	"strconv"

)



type ClienteController interface {
	FindAll(ctx *gin.Context) 
	Save(ctx *gin.Context) 
	GetClientById(ctx *gin.Context)
	
}

type controllerCliente struct {
	service service.ClienteService
}


func NewClienteController(service service.ClienteService) ClienteController {
	return &controllerCliente {
		service: service,
	}
}

func (c *controllerCliente) FindAll(ctx *gin.Context){
	clientes, err := c.service.FindAll()
	if(err != nil){
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	ctx.JSON(200, clientes)


}

func (c *controllerCliente) Save(ctx *gin.Context){
	var cliente models.Cliente
	ctx.BindJSON(&cliente)
	id, err := c.service.Save(cliente)
	if(err != nil){
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()}
		ctx.JSON(400, response) 
		return
	}
	cliente.Id_cliente = id
	response := models.GoodResponse{"Objeto criado", "Ok", cliente}

	ctx.JSON(200, response)

}


func (c *controllerCliente) GetClientById(ctx *gin.Context){
	id_cliente := ctx.Param("id_cliente") 
	teste, err := strconv.Atoi(id_cliente)
	customer, err := c.service.GetClientById(teste)
	if(err != nil){
		fmt.Println(err)
		response := models.BadResponse{"Cliente não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, customer)


}