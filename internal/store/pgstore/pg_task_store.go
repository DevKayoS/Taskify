package pgstore

import (
	"context"

	"github.com/DevKayoS/taskify/internal/store"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgTaskStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPgTaskStore(pool *pgxpool.Pool) PgTaskStore {
	return PgTaskStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (p *PgTaskStore) CreateTask(ctx context.Context, title, description string, priority int32) (store.Task, error) {
	task, err := p.Queries.CreateTask(ctx, CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}
	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}, nil
}

func (p *PgTaskStore) GetTaskById(ctx context.Context, id int32) (store.Task, error) {
	task, err := p.Queries.GetTaskById(ctx, id)
	if err != nil {
		return store.Task{}, err
	}

	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}, nil
}

func (p *PgTaskStore) ListTasks(ctx context.Context) ([]store.Task, error) {
	tasks, err := p.Queries.ListTasks(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]store.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, store.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Priority:    task.Priority,
			CreatedAt:   task.CreatedAt.Time,
			UpdatedAt:   task.UpdatedAt.Time,
		})
	}
	return result, nil
}

func (p *PgTaskStore) UpdateTask(ctx context.Context, id int32, title, description string, priority int32) (store.Task, error) {
	task, err := p.Queries.UpdateTask(ctx, UpdateTaskParams{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
	})

	if err != nil {
		return store.Task{}, nil
	}
	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}, nil
}

func (p *PgTaskStore) DeleteTask(ctx context.Context, id int32) error {
	err := p.Queries.DeleteTask(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
