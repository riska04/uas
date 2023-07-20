package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type StrukturService interface {
	All(ctx context.Context) ([]entity.StrukturEntity, error)
	FindByID(ctx context.Context, Struktur string) (entity.StrukturEntity, error)
	Create(ctx context.Context, input request.StrukturCreate) (entity.StrukturEntity, error)
	Update(ctx context.Context, input request.StrukturUpdate) error
	Delete(ctx context.Context, Struktur string) error
}

type strukturService struct {
	strukturRepository repository.StrukturRepository
}

// All implements StrukturService.
func (s *strukturService) All(ctx context.Context) ([]entity.StrukturEntity, error) {
	result, err := s.strukturRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements StrukturService.
func (s *strukturService) Create(ctx context.Context, input request.StrukturCreate) (entity.StrukturEntity, error) {
	Struktur := entity.StrukturEntity{
		NamaStruktur: input.NamaStruktur,
	}
	result, err := s.strukturRepository.Create(ctx, Struktur)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements StrukturService.
func (s *strukturService) Delete(ctx context.Context, IdStruktur string) error {
	return s.strukturRepository.Delete(ctx, IdStruktur)
}

// FindByID implements StrukturService.
func (s *strukturService) FindByID(ctx context.Context, IdStruktur string) (entity.StrukturEntity, error) {
	result, err := s.strukturRepository.FindByID(ctx, IdStruktur)
	if err != nil {
		return result, err
	}
	if result.IdStruktur == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements StrukturService.
func (s *strukturService) Update(ctx context.Context, input request.StrukturUpdate) error {
	result, err := s.strukturRepository.FindByID(ctx, input.IdStruktur)
	if err != nil {
		return err
	}

	if result.IdStruktur == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.IdStruktur = input.IdStruktur
	result.NamaStruktur = input.NamaStruktur
	return s.strukturRepository.Update(ctx, result)
}

func NewStrukturService(strukturRepository repository.StrukturRepository) StrukturService {
	return &strukturService{
		strukturRepository: strukturRepository,
	}
}
