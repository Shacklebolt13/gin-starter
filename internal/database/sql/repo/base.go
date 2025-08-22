package repo

import (
	"context"
	"gin-starter/internal/database/sql/model"
)

type ReadableModelRepository[T model.Model, PK any] interface {
	FindByID(ctx context.Context, id PK) (T, error)
	FindAll(ctx context.Context) ([]T, error)
}

type WritableModelRepository[T model.Model, PK any] interface {
	Create(ctx context.Context, model T) (T, error)
	Delete(ctx context.Context, id PK) error
	CreateOrSave(ctx context.Context, model T) (T, error)
}

type ModelRepository[T model.Model, PK any] interface {
	ReadableModelRepository[T, PK]
	WritableModelRepository[T, PK]
}
