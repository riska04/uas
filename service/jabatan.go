package service

import (
	"context"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type JabatanService interface {
	All(ctx context.Context) ([]entity.JabatanEntity, error)
	FindByID(ctx context.Context, IdUser string) (entity.JabatanEntity, error)
	Create(ctx context.Context, input request.JabatanCreate) (entity.JabatanEntity, error)
	Update(ctx context.Context, input request.JabatanUpdate) error
	Delete(ctx context.Context, IdUser string) error
}

type jabatanService struct {
	jabatanRepository repository.JabatanRepository
}

// All implements JabatanService.
func (s *jabatanService) All(ctx context.Context) ([]entity.JabatanEntity, error) {
	result, err := s.jabatanRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements JabatanService.
func (s *jabatanService) Create(ctx context.Context, input request.JabatanCreate) (entity.JabatanEntity, error) {
	panic("unimplemented")
}

// Delete implements JabatanService.
func (s *jabatanService) Delete(ctx context.Context, IdUser string) error {
	panic("unimplemented")
}

// FindByID implements JabatanService.
func (s *jabatanService) FindByID(ctx context.Context, IdUser string) (entity.JabatanEntity, error) {
	panic("unimplemented")
}

// Update implements JabatanService.
func (s *jabatanService) Update(ctx context.Context, input request.JabatanUpdate) error {
	panic("unimplemented")
}

func NewJabatanService(jabatanRepository repository.JabatanRepository) JabatanService {
	return &jabatanService{
		jabatanRepository: jabatanRepository,
	}
}
