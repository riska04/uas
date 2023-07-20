package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type FakultasService interface {
	All(ctx context.Context) ([]entity.FakultasEntity, error)
	FindByID(ctx context.Context, Fakultas string) (entity.FakultasEntity, error)
	Create(ctx context.Context, input request.FakultasCreate) (entity.FakultasEntity, error)
	Update(ctx context.Context, input request.FakultasUpdate) error
	Delete(ctx context.Context, Fakultas string) error
}

type fakultasService struct {
	fakultasRepository repository.FakultasRepository
}

// All implements FakultasService.
func (s *fakultasService) All(ctx context.Context) ([]entity.FakultasEntity, error) {
	result, err := s.fakultasRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements FakultasService.
func (s *fakultasService) Create(ctx context.Context, input request.FakultasCreate) (entity.FakultasEntity, error) {
	Fakultas := entity.FakultasEntity{}
	Fakultas.NamaFakultas = input.NamaFakultas
	Fakultas.SingkatanFakultas = input.SingkatanFakultas
	result, err := s.fakultasRepository.Create(ctx, Fakultas)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements FakultasService.
func (s *fakultasService) Delete(ctx context.Context, IdFakultas string) error {
	return s.fakultasRepository.Delete(ctx, IdFakultas)
}

// FindByID implements FakultasService.
func (s *fakultasService) FindByID(ctx context.Context, IdFakultas string) (entity.FakultasEntity, error) {
	result, err := s.fakultasRepository.FindByID(ctx, IdFakultas)
	if err != nil {
		return result, err
	}
	if result.IdFakultas == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements FakultasService.
func (s *fakultasService) Update(ctx context.Context, input request.FakultasUpdate) error {
	result, err := s.fakultasRepository.FindByID(ctx, input.IdFakultas)
	if err != nil {
		return err
	}

	if result.IdFakultas == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.NamaFakultas = input.NamaFakultas
	result.SingkatanFakultas = input.SingkatanFakultas
	return s.fakultasRepository.Update(ctx, result)
}

func NewFakultasService(fakultasRepository repository.FakultasRepository) FakultasService {
	return &fakultasService{
		fakultasRepository: fakultasRepository,
	}
}
