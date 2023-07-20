package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type NotulenService interface {
	All(ctx context.Context) ([]entity.NotulenEntity, error)
	FindByID(ctx context.Context, IdNotulen string) (entity.NotulenEntity, error)
	Create(ctx context.Context, input request.NotulenCreate) (entity.NotulenEntity, error)
	Update(ctx context.Context, input request.NotulenUpdate) error
	Delete(ctx context.Context, IdNotulen string) error
}

type notulenService struct {
	notulenRepository repository.NotulenRepository
}

// All implements NotulenService.
func (s *notulenService) All(ctx context.Context) ([]entity.NotulenEntity, error) {
	result, err := s.notulenRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements NotulenService.
func (s *notulenService) Create(ctx context.Context, input request.NotulenCreate) (entity.NotulenEntity, error) {
	Berkas := entity.NotulenEntity{
		IdBerkas: input.IdBerkas,
		Catatan:  input.Catatan,
	}
	result, err := s.notulenRepository.Create(ctx, Berkas)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements NotulenService.
func (s *notulenService) Delete(ctx context.Context, IdNotulen string) error {
	return s.notulenRepository.Delete(ctx, IdNotulen)
}

// FindByID implements NotulenService.
func (s *notulenService) FindByID(ctx context.Context, IdNotulen string) (entity.NotulenEntity, error) {
	result, err := s.notulenRepository.FindByID(ctx, IdNotulen)
	if err != nil {
		return result, err
	}
	if result.IdNotulen == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements NotulenService.
func (s *notulenService) Update(ctx context.Context, input request.NotulenUpdate) error {
	result, err := s.notulenRepository.FindByID(ctx, input.IdNotulen)
	if err != nil {
		return err
	}

	if result.IdNotulen == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.IdBerkas = input.IdBerkas
	result.Catatan = input.Catatan

	return s.notulenRepository.Update(ctx, result)
}

func NewNotulenService(notulenRepository repository.NotulenRepository) NotulenService {
	return &notulenService{
		notulenRepository: notulenRepository,
	}
}
