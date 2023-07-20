package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type LembagaService interface {
	All(ctx context.Context) ([]entity.LembagaEntity, error)
	FindByID(ctx context.Context, IdLembaga string) (entity.LembagaEntity, error)
	Create(ctx context.Context, input request.LembagaCreate) (entity.LembagaEntity, error)
	Update(ctx context.Context, input request.LembagaUpdate) error
	Delete(ctx context.Context, IdLembaga string) error
}

type lembagaService struct {
	lembagaRepository repository.LembagaRepository
}

// All implements LembagaService.
func (s *lembagaService) All(ctx context.Context) ([]entity.LembagaEntity, error) {
	result, err := s.lembagaRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements LembagaService.
func (s *lembagaService) Create(ctx context.Context, input request.LembagaCreate) (entity.LembagaEntity, error) {
	Lembaga := entity.LembagaEntity{
		NamaLembaga: input.NamaLembaga,
	}
	result, err := s.lembagaRepository.Create(ctx, Lembaga)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements LembagaService.
func (s *lembagaService) Delete(ctx context.Context, IdLembaga string) error {
	return s.lembagaRepository.Delete(ctx, IdLembaga)
}

// FindByID implements LembagaService.
func (s *lembagaService) FindByID(ctx context.Context, IdLembaga string) (entity.LembagaEntity, error) {
	result, err := s.lembagaRepository.FindByID(ctx, IdLembaga)
	if err != nil {
		return result, err
	}
	if result.IdLembaga == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements LembagaService.
func (s *lembagaService) Update(ctx context.Context, input request.LembagaUpdate) error {
	result, err := s.lembagaRepository.FindByID(ctx, input.IdLembaga)
	if err != nil {
		return err
	}

	if result.IdLembaga == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.NamaLembaga = input.NamaLembaga
	return s.lembagaRepository.Update(ctx, result)
}

func NewLembagaService(lembagaRepository repository.LembagaRepository) LembagaService {
	return &lembagaService{
		lembagaRepository: lembagaRepository,
	}
}
