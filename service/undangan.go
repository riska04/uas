package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type UndanganService interface {
	All(ctx context.Context) ([]entity.UndanganEntity, error)
	FindByID(ctx context.Context, Undangan string) (entity.UndanganEntity, error)
	Create(ctx context.Context, input request.UndanganCreate) (entity.UndanganEntity, error)
	Update(ctx context.Context, input request.UndanganUpdate) error
	Delete(ctx context.Context, Undangan string) error
}

type undanganService struct {
	undanganRepository repository.UndanganRepository
}

// All implements UndanganService.
func (s *undanganService) All(ctx context.Context) ([]entity.UndanganEntity, error) {
	result, err := s.undanganRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements UndanganService.
func (s *undanganService) Create(ctx context.Context, input request.UndanganCreate) (entity.UndanganEntity, error) {
	Struktur := entity.UndanganEntity{
		IdUnit:    input.IdUnit,
		IdJabatan: input.IdJabatan,
		IdUser:    input.IdUser,
		IdBerkas:  input.IdBerkas,
	}
	result, err := s.undanganRepository.Create(ctx, Struktur)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements UndanganService.
func (s *undanganService) Delete(ctx context.Context, IdUndangan string) error {
	return s.undanganRepository.Delete(ctx, IdUndangan)
}

// FindByID implements UndanganService.
func (s *undanganService) FindByID(ctx context.Context, IdUndangan string) (entity.UndanganEntity, error) {
	result, err := s.undanganRepository.FindByID(ctx, IdUndangan)
	if err != nil {
		return result, err
	}
	if result.IdUndangan == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements UndanganService.
func (s *undanganService) Update(ctx context.Context, input request.UndanganUpdate) error {
	result, err := s.undanganRepository.FindByID(ctx, input.IdUndangan)
	if err != nil {
		return err
	}

	if result.IdUndangan == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.IdUnit = input.IdUnit
	result.IdJabatan = input.IdJabatan
	result.IdUser = input.IdUser
	result.IdBerkas = input.IdBerkas
	return s.undanganRepository.Update(ctx, result)
}

func NewUndanganService(undanganRepository repository.UndanganRepository) UndanganService {
	return &undanganService{
		undanganRepository: undanganRepository,
	}
}
