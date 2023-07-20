package repository

import (
	"context"
	"database/sql"
	"materi/model/entity"
)

type LogActivityRepository interface {
	All(ctx context.Context) ([]entity.LogActivityEntity, error)
	FindByID(ctx context.Context, LogActivity string) (entity.LogActivityEntity, error)
	Create(ctx context.Context, input entity.LogActivityEntity) (entity.LogActivityEntity, error)
	Update(ctx context.Context, input entity.LogActivityEntity) error
	Delete(ctx context.Context, IdLogActivity string) error
}
type LogActivityConnetion struct {
	tx *sql.DB
}

// All implements LogActivityRepository.
func (*LogActivityConnetion) All(ctx context.Context) ([]entity.LogActivityEntity, error) {
	panic("unimplemented")
}

// Create implements LogActivityRepository.
func (*LogActivityConnetion) Create(ctx context.Context, input entity.LogActivityEntity) (entity.LogActivityEntity, error) {
	panic("unimplemented")
}

// Delete implements LogActivityRepository.
func (*LogActivityConnetion) Delete(ctx context.Context, IdLogActivity string) error {
	panic("unimplemented")
}

// FindByID implements LogActivityRepository.
func (*LogActivityConnetion) FindByID(ctx context.Context, LogActivity string) (entity.LogActivityEntity, error) {
	panic("unimplemented")
}

// Update implements LogActivityRepository.
func (*LogActivityConnetion) Update(ctx context.Context, input entity.LogActivityEntity) error {
	panic("unimplemented")
}

func NewLogActivityRepository(DB *sql.DB) LogActivityRepository {
	return &LogActivityConnetion{
		tx: DB,
	}
}
