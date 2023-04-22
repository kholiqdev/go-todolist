package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kholiqcode/go-todolist/internal/todo/dtos"
	"github.com/kholiqcode/go-todolist/internal/todo/service"
	"github.com/kholiqcode/go-todolist/utils"
)

type TodoHandler interface {
	getTodos(w http.ResponseWriter, r *http.Request)
	// getTodo(w http.ResponseWriter, r *http.Request)
	// createTodo(w http.ResponseWriter, r *http.Request)
	// updateTodo(w http.ResponseWriter, r *http.Request)
	MapRoutes()
}
type todoHandlerImpl struct {
	route   *chi.Mux
	todoSvc service.TodoService
}

func (h *todoHandlerImpl) getTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	activityGroupID := utils.ValidateQueryParamInt(r, "activity_group_id")

	request := dtos.GetTodosRequest{
		ActivityGroupID: int32(activityGroupID),
	}

	activityGroupsResp, err := h.todoSvc.FindAll(ctx, request)
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, activityGroupsResp, 200, "Success")
}
