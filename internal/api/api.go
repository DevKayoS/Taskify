package api

import (
	"github.com/DevKayoS/taskify/internal/services"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	Router      *chi.Mux
	TaskService services.TaskService
}
