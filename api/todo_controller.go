package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snoeffels/mygo/models"
	"github.com/snoeffels/mygo/persistence"
	"net/http"
	"strconv"
)

type TodoController struct {
	TodoService persistence.TodoService
}

func ProvideTodoAPI(t persistence.TodoService) TodoController {
	return TodoController{TodoService: t}
}

// FindAll godoc
// @Summary FindAll Todo List
// @Description FindAll Todo List
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {array} models.Todo
// @Failure 500 {object} models.APIError
// @Router /todo [get]
func (t *TodoController) FindAll(c *gin.Context) {
	// Usage Roles
	//roles, ok := c.Get("roles")
	//if !ok {
	//	return
	//}

	todos, err := t.TodoService.FindAll()
	if err != nil {
		models.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todos)
}

// Create godoc
// @Summary Create Todo
// @Description Create Todo
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param todo body models.Todo true "Create Todo"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.APIError
// @Failure 500 {object} models.APIError
// @Router /todo [post]
func (t *TodoController) Create(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = t.TodoService.Create(&todo)
	if err != nil {
		models.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

// FindById godoc
// @Summary FindById Todo
// @Description FindById Todo
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path integer true "ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.APIError
// @Failure 404 {object} models.APIError
// @Router /todo/{id} [get]
func (t *TodoController) FindById(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := t.TodoService.FindById(uint(id))
	if err != nil {
		models.SendSimpleErrorResponse(c, http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Update godoc
// @Summary Update Todo
// @Description Update Todo
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path integer true "ID"
// @Param todo body models.Todo true "Update Todo"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.APIError
// @Failure 404 {object} models.APIError
// @Failure 500 {object} models.APIError
// @Router /todo/{id} [put]
func (t *TodoController) Update(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := t.TodoService.FindById(uint(id))
	if err != nil {
		models.SendErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var updateTodo models.Todo
	err = c.BindJSON(&updateTodo)
	if err != nil {
		models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = t.TodoService.Update(&todo, &updateTodo)
	if err != nil {
		models.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Delete godoc
// @Summary Delete Todo
// @Description Delete Todo
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path integer true "ID"
// @Success 204 {string} {}
// @Failure 400 {object} models.APIError
// @Failure 404 {object} models.APIError
// @Failure 500 {object} models.APIError
// @Router /todo/{id} [delete]
func (t *TodoController) Delete(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := t.TodoService.FindById(uint(id))
	if err != nil {
		models.SendErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	err = t.TodoService.DeleteById(&todo, uint(id))
	if err != nil {
		models.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, "")
}
