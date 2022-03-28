package api

import (
	"context"
	"fmt"
	db "go-lang-app/db/sqlc"
	"go-lang-app/token"
	"go-lang-app/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		fmt.Println(v)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/ping", server.ping)
		v1.POST("/register", server.createUser)
		v1.POST("/login", server.loginUser)
		v1.POST("/refresh-token", server.renewAccessToken)

	}

	user := v1.Group("user").Use(authMiddleware(server.tokenMaker))
	{
		user.GET("/profile", server.profile)
		user.PUT("/change-password", server.changePassword)
		user.PUT("/change-profile-pic", server.changeProfilePic)
		user.GET("/logout", server.logout)

	}

	todo := v1.Group("/todo").Use(authMiddleware(server.tokenMaker))
	{
		todo.GET("/", server.getTodos)
		todo.POST("/", server.createTodo)
		todo.GET("/:id", server.getTodo)
		todo.PUT("/:id", server.updateTodo)
		todo.DELETE("/:id", server.deleteTodo)
	}

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: server.router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
