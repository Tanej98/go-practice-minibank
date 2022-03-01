package api

import (
	db "github.com/Tanej98/minibank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	bank   *db.Bank
	router *gin.Engine
}

func NewServer(bank *db.Bank) *Server {
	server := &Server{bank: bank}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
