package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type BerkasService interface {
	All(ctx context.Context) ([]entity.BerkasEntity, error)
	FindByID(ctx context.Context, IdBerkas string) (entity.BerkasEntity, error)
	Create(ctx context.Context, input request.BerkasCreate) (entity.BerkasEntity, error)
	Update(ctx context.Context, input request.BerkasUpdate) error
	Delete(ctx context.Context, IdBerkas string) error
}

type berkasService struct {
	berkasRepository repository.BerkasRepository
}

// All implements BerkasService.
func (s *berkasService) All(ctx context.Context) ([]entity.BerkasEntity, error) {
	result, err := s.berkasRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements BerkasService.
func (s *berkasService) Create(ctx context.Context, input request.BerkasCreate) (entity.BerkasEntity, error) {
	Berkas := entity.BerkasEntity{}
	Berkas.File = input.File
	Berkas.NamaFile = input.NamaFile
	Berkas.TanggalUpload = input.TanggalUpload
	Berkas.TanggalAcara = input.TanggalAcara
	Berkas.StartJam = input.StartJam
	Berkas.SelesaiJam = input.SelesaiJam
	Berkas.IdUnit = input.IdUnit
	result, err := s.berkasRepository.Create(ctx, Berkas)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements BerkasService.
func (s *berkasService) Delete(ctx context.Context, IdBerkas string) error {
	return s.berkasRepository.Delete(ctx, IdBerkas)
}

// FindByID implements BerkasService.
func (s *berkasService) FindByID(ctx context.Context, IdBerkas string) (entity.BerkasEntity, error) {
	result, err := s.berkasRepository.FindByID(ctx, IdBerkas)
	if err != nil {
		return result, err
	}
	if result.IdBerkas == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements BerkasService.
func (s *berkasService) Update(ctx context.Context, input request.BerkasUpdate) error {
	result, err := s.berkasRepository.FindByID(ctx, input.IdBerkas)
	if err != nil {
		return err
	}

	if result.IdBerkas == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.File = input.File
	result.NamaFile = input.NamaFile
	result.TanggalUpload = input.TanggalUpload
	result.TanggalAcara = input.TanggalAcara
	result.StartJam = input.StartJam
	result.SelesaiJam = input.SelesaiJam
	return s.berkasRepository.Update(ctx, result)
}

func NewBerkasService(berkasRepository repository.BerkasRepository) BerkasService {
	return &berkasService{
		berkasRepository: berkasRepository,
	}
}
