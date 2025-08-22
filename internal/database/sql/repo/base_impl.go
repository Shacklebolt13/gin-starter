package repo

import (
	"context"
	"fmt"
	"gin-starter/internal/database/sql/model"

	"gorm.io/gorm"
)

type readableModelRepositoryImpl[T model.Model, PK any] struct {
	db *gorm.DB
}

func (r *readableModelRepositoryImpl[T, PK]) FindByID(ctx context.Context, id PK) (T, error) {
	queryString := fmt.Sprintf("id = %s", model.ID_FIELD)

	result, err := gorm.G[T](r.db).Where(queryString, id).First(ctx)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *readableModelRepositoryImpl[T, PK]) FindAll(ctx context.Context) ([]T, error) {
	rows, err := gorm.G[T](r.db).Find(ctx)
	if err != nil {
		return nil, err
	}

	return rows, err
}

type writableModelRepositoryImpl[T model.Model, PK any] struct {
	db *gorm.DB
}

func (w *writableModelRepositoryImpl[T, PK]) Create(ctx context.Context, model T) (T, error) {
	if err := gorm.G[T](w.db).Create(ctx, &model); err != nil {
		return model, err
	}

	return model, nil
}

func (w *writableModelRepositoryImpl[T, PK]) Delete(ctx context.Context, id PK) error {
	_, err := gorm.G[T](w.db).Where("id = ?", id).Delete(ctx)
	return err
}

func (w *writableModelRepositoryImpl[T, PK]) CreateOrSave(ctx context.Context, model T) (T, error) {
	if err := w.db.Save(&model).Error; err != nil {
		return model, err
	}

	return model, nil
}
