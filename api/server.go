package api

import (
	db "simplebank/db/sqlc"
	"simplebank/token"
	"simplebank/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	config     util.Config
}

func NewServer(store *db.Store, config util.Config) (*Server, error) {
	server := &Server{store: store, config: config}
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return server, err
	}
	server.tokenMaker = tokenMaker
	server.SetUpRoter()
	return server, nil
}

func (server *Server) SetUpRoter() {
	router := gin.Default()
	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRouter.POST("/accounts", server.CreateAccount)
	authRouter.GET("/accounts/:id", server.GetAccount)
	authRouter.GET("/accounts", server.ListAccount)
	authRouter.POST("/transfer", server.Transfer)
	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.LoginUser)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func errorMessage(errorMessage string) gin.H {
	return gin.H{"error": errorMessage}
}
