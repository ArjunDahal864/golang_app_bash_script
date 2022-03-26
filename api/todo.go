package api

import (
	"database/sql"
	"errors"
	db "go-lang-app/db/sqlc"
	"go-lang-app/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

type todoUpdateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsCompleted bool   `json:"is_completed" binding:"required"`
}

func (s *Server) createTodo(ctx *gin.Context) {
	var req todoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	args := db.CreateTodoParams{
		Title: req.Title,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
		IsCompleted: req.IsCompleted,
		UserID:      user.UserID,
	}
	todo, err := s.store.CreateTodo(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, todo)
}

func (s *Server) getTodo(ctx *gin.Context) {
	todoID := ctx.Param("id")
	// convert to int32
	todoIDInt, err := strconv.ParseInt(todoID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	todo, err := s.store.GetTodo(ctx, int32(todoIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (s *Server) getTodos(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	// convert to int32
	limitInt, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	offsetInt, err := strconv.ParseInt(offset, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	args := db.GetTodosParams{
		UserID: user.UserID,
		Limit:  int32(limitInt),
		Offset: int32(offsetInt),
	}

	todos, err := s.store.GetTodos(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (s *Server) updateTodo(ctx *gin.Context) {
	todoID := ctx.Param("id")
	todoIDInt, err := strconv.ParseInt(todoID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req todoUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	todo, err := s.store.GetTodo(ctx, int32(todoIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if todo.UserID != user.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Unauthorized")))
		return
	}

	args := db.UpdateTodoParams{
		Title: req.Title,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
		IsCompleted: req.IsCompleted,
		ID:          int32(todoIDInt),
	}

	err = s.store.UpdateTodo(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, req)
}

func (s *Server) deleteTodo(ctx *gin.Context) {
	todoID := ctx.Param("todo_id")
	todoIDInt, err := strconv.ParseInt(todoID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	todo, err := s.store.GetTodo(ctx, int32(todoIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if todo.UserID != user.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Unauthorized")))
		return
	}

	err = s.store.DeleteTodo(ctx, int32(todoIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Todo deleted"})
}
