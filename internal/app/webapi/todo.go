package webapi

import (
	"net/http"
	"proposal/internal/app/usecase"
	"proposal/internal/pkgs/maybe"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoUseCase usecase.TodoUseCase
}

var _ Controller = (*TodoController)(nil)

// Attach implements Controller.
func (c *TodoController) Attach(r gin.IRouter) {
	r.POST("/", c.addTodo)
}

func (c *TodoController) addTodo(ctx *gin.Context) {
	var request struct {
		Title   string     `json:"title"`
		DueDate *time.Time `json:"due_date"`
	}

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = c.todoUseCase.AddTodo(
		ctx,
		usecase.AddTodoInput{
			Title:   request.Title,
			DueDate: maybe.FromPointer(request.DueDate),
		},
	)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
