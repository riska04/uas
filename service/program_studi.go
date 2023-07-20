package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type ProgramStudiService interface {
	All(ctx context.Context) ([]entity.ProgramStudiEntity, error)
	FindByID(ctx context.Context, IdProgramStudi string) (entity.ProgramStudiEntity, error)
	Create(ctx context.Context, input request.ProgramStudiCreate) (entity.ProgramStudiEntity, error)
	Update(ctx context.Context, input request.ProgramStudiUpdate) error
	Delete(ctx context.Context, IdProgramStudi string) error
}

type programstudiService struct {
	programstudiRepository repository.ProgramStudiRepository
}

// All implements ProgramStudiService.
func (s *programstudiService) All(ctx context.Context) ([]entity.ProgramStudiEntity, error) {
	result, err := s.programstudiRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements ProgramStudiService.
func (s *programstudiService) Create(ctx context.Context, input request.ProgramStudiCreate) (entity.ProgramStudiEntity, error) {
	programstudi := entity.ProgramStudiEntity{
		ProgramStudi: input.ProgramStudi,
	}
	result, err := s.programstudiRepository.Create(ctx, programstudi)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements ProgramStudiService.
func (s *programstudiService) Delete(ctx context.Context, IdProgramStudi string) error {
	return s.programstudiRepository.Delete(ctx, IdProgramStudi)
}

// FindByID implements ProgramStudiService.
func (s *programstudiService) FindByID(ctx context.Context, IdProgramStudi string) (entity.ProgramStudiEntity, error) {
	result, err := s.programstudiRepository.FindByID(ctx, IdProgramStudi)
	if err != nil {
		return result, err
	}
	if result.ProdiId == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements ProgramStudiService.
func (s *programstudiService) Update(ctx context.Context, input request.ProgramStudiUpdate) error {
	result, err := s.programstudiRepository.FindByID(ctx, input.ProdiId)
	if err != nil {
		return err
	}

	if result.ProdiId == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.ProdiId = input.ProdiId
	result.IdFakultas = input.IdFakultas
	result.ProgramStudi = input.ProgramStudi
	return s.programstudiRepository.Update(ctx, result)
}

func NewProgramStudiService(programstudiRepository repository.ProgramStudiRepository) ProgramStudiService {
	return &programstudiService{
		programstudiRepository: programstudiRepository,
	}
}
