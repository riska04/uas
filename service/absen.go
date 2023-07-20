package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type AbsenService interface {
	All(ctx context.Context) ([]entity.AbsenEntity, error)
	FindByID(ctx context.Context, IdAbsen string) (entity.AbsenEntity, error)
	Create(ctx context.Context, input request.AbsenCreate) (entity.AbsenEntity, error)
	Update(ctx context.Context, input request.AbsenUpdate) error
	Delete(ctx context.Context, IdAbsen string) error
}

type absenService struct {
	absenRepository repository.AbsenRepository
}

// All implements AbsenService.
func (s *absenService) All(ctx context.Context) ([]entity.AbsenEntity, error) {
	result, err := s.absenRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements AbsenService.
func (s *absenService) Create(ctx context.Context, input request.AbsenCreate) (entity.AbsenEntity, error) {
	Berkas := entity.AbsenEntity{
		IdBerkas:  input.IdBerkas,
		IdUnit:    input.IdUnit,
		IdJabatan: input.IdJabatan,
		IdUser:    input.IdUser,
	}
	result, err := s.absenRepository.Create(ctx, Berkas)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements AbsenService.
func (s *absenService) Delete(ctx context.Context, IdAbsen string) error {
	return s.absenRepository.Delete(ctx, IdAbsen)
}

// FindByID implements AbsenService.
func (s *absenService) FindByID(ctx context.Context, IdAbsen string) (entity.AbsenEntity, error) {
	result, err := s.absenRepository.FindByID(ctx, IdAbsen)
	if err != nil {
		return result, err
	}
	if result.IdAbsen == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements AbsenService.
func (s *absenService) Update(ctx context.Context, input request.AbsenUpdate) error {
	result, err := s.absenRepository.FindByID(ctx, input.IdAbsen)
	if err != nil {
		return err
	}

	if result.IdAbsen == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.IdBerkas = input.IdBerkas
	result.IdUnit = input.IdUnit
	result.IdJabatan = input.IdJabatan
	result.IdUser = input.IdUser
	return s.absenRepository.Update(ctx, result)
}

// Upload implements absenService.
func (*absenService) Upload(ctx context.Context, input request.AbsenProfile) error {
	panic("unimplemented")
}

func NewAbsenService(absenRepository repository.AbsenRepository) AbsenService {
	return &absenService{
		absenRepository: absenRepository,
	}
}
