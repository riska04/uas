package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type JabatanService interface {
	All(ctx context.Context) ([]entity.JabatanEntity, error)
	FindByID(ctx context.Context, IdJabatan string) (entity.JabatanEntity, error)
	Create(ctx context.Context, input request.JabatanCreate) (entity.JabatanEntity, error)
	Update(ctx context.Context, input request.JabatanUpdate) error
	Delete(ctx context.Context, IdJabatan string) error
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
	Jabatan := entity.JabatanEntity{
		NamaJabatan: input.NamaJabatan,
	}
	result, err := s.jabatanRepository.Create(ctx, Jabatan)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements JabatanService.
func (s *jabatanService) Delete(ctx context.Context, IdJabatan string) error {
	return s.jabatanRepository.Delete(ctx, IdJabatan)
}

// FindByID implements JabatanService.
func (s *jabatanService) FindByID(ctx context.Context, IdJabatan string) (entity.JabatanEntity, error) {
	result, err := s.jabatanRepository.FindByID(ctx, IdJabatan)
	if err != nil {
		return result, err
	}
	if result.IdJabatan == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements JabatanService.
func (s *jabatanService) Update(ctx context.Context, input request.JabatanUpdate) error {
	result, err := s.jabatanRepository.FindByID(ctx, input.IdJabatan)
	if err != nil {
		return err
	}

	if result.IdJabatan == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.NamaJabatan = input.NamaJabatan
	return s.jabatanRepository.Update(ctx, result)
}

func NewJabatanService(jabatanRepository repository.JabatanRepository) JabatanService {
	return &jabatanService{
		jabatanRepository: jabatanRepository,
	}
}
