package service

import (
	"context"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type LogActivityService interface {
	All(ctx context.Context) ([]entity.LogActivityEntity, error)
	FindByID(ctx context.Context, LogActivity string) (entity.LogActivityEntity, error)
	Create(ctx context.Context, input request.LogActivityCreate) (entity.LogActivityEntity, error)
	Update(ctx context.Context, input request.LogActivityUpdate) error
	Delete(ctx context.Context, IdLogActivity string) error
}

type logactivityService struct {
	logActivityRepository repository.LogActivityRepository
}
