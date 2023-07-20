package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type UnitService interface {
	All(ctx context.Context) ([]entity.UnitEntity, error)
	FindByID(ctx context.Context, Unit string) (entity.UnitEntity, error)
	Create(ctx context.Context, input request.UnitCreate) (entity.UnitEntity, error)
	Update(ctx context.Context, input request.UnitUpdate) error
	Delete(ctx context.Context, Unit string) error
}

type unitService struct {
	unitRepository repository.UnitRepository
}

// All implements UnitService.
func (s *unitService) All(ctx context.Context) ([]entity.UnitEntity, error) {
	result, err := s.unitRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements UnitService.
func (s *unitService) Create(ctx context.Context, input request.UnitCreate) (entity.UnitEntity, error) {
	Unit := entity.UnitEntity{}
	Unit.IdStruktur = input.IdStruktur
	result, err := s.unitRepository.Create(ctx, Unit)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements UnitService.
func (s *unitService) Delete(ctx context.Context, IdUnit string) error {
	return s.unitRepository.Delete(ctx, IdUnit)
}

// FindByID implements UnitService.
func (s *unitService) FindByID(ctx context.Context, IdUnit string) (entity.UnitEntity, error) {
	result, err := s.unitRepository.FindByID(ctx, IdUnit)
	if err != nil {
		return result, err
	}
	if result.IdUnit == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements UnitService.
func (s *unitService) Update(ctx context.Context, input request.UnitUpdate) error {
	unit, err := s.unitRepository.FindByID(ctx, input.IdUnit)
	if err != nil {
		return err
	}

	if unit.IdUnit == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	unit.IdStruktur = input.IdStruktur
	unit.Status = input.Status
	return s.unitRepository.Update(ctx, unit)
}

func NewUnitService(unitRepository repository.UnitRepository) UnitService {
	return &unitService{
		unitRepository: unitRepository,
	}
}
