package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/entity"
	"materi/model/request"
	"materi/repository"
)

type UserService interface {
	All(ctx context.Context) ([]entity.UserEntity, error)
	FindByID(ctx context.Context, Unit string) (entity.UserEntity, error)
	Create(ctx context.Context, input request.UserCreate) (entity.UserEntity, error)
	Update(ctx context.Context, input request.UserUpdate) error
	Delete(ctx context.Context, User string) error
}

type userService struct {
	userRepository repository.UserRepository
}

// All implements UserService.
func (s *userService) All(ctx context.Context) ([]entity.UserEntity, error) {
	result, err := s.userRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements UserService.
func (s *userService) Create(ctx context.Context, input request.UserCreate) (entity.UserEntity, error) {
	Struktur := entity.UserEntity{
		IdUnit:      input.IdUnit,
		IdJabatan:   input.IdJabatan,
		NamaLengkap: input.NamaLengkap,
		Alamat:      input.Alamat,
		NomorHp:     input.NomorHp,
		UserName:    input.UserName,
		Password:    input.Password,
		Status:      input.Status,
	}
	result, err := s.userRepository.Create(ctx, Struktur)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements UserService.
func (s *userService) Delete(ctx context.Context, IdUser string) error {
	return s.userRepository.Delete(ctx, IdUser)
}

// FindByID implements UserService.
func (s *userService) FindByID(ctx context.Context, IdUser string) (entity.UserEntity, error) {
	result, err := s.userRepository.FindByID(ctx, IdUser)
	if err != nil {
		return result, err
	}
	if result.IdUser == "" {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements UserService.
func (s *userService) Update(ctx context.Context, input request.UserUpdate) error {
	result, err := s.userRepository.FindByID(ctx, input.IdUser)
	if err != nil {
		return err
	}

	if result.IdUser == "" {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.IdUnit = input.IdUnit
	result.IdJabatan = input.IdJabatan
	result.NamaLengkap = input.NamaLengkap
	result.Alamat = input.Alamat
	result.NomorHp = input.NomorHp
	result.UserName = input.Username
	result.Password = input.Password

	return s.userRepository.Update(ctx, result)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
