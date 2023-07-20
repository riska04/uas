package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type UserRepository interface {
	All(ctx context.Context) ([]entity.UserEntity, error)
	FindByID(ctx context.Context, User string) (entity.UserEntity, error)
	Create(ctx context.Context, input entity.UserEntity) (entity.UserEntity, error)
	Update(ctx context.Context, input entity.UserEntity) error
	Delete(ctx context.Context, User string) error
}

type UserConnetion struct {
	tx *sql.DB
}

// All implements UserRepository.
func (r *UserConnetion) All(ctx context.Context) ([]entity.UserEntity, error) {
	SQL := "SELECT * FROM user"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user_ := []entity.UserEntity{}
	for rows.Next() {
		user := entity.UserEntity{}
		err := rows.Scan(
			&user.IdUser,
			&user.IdUnit,
			&user.IdJabatan,
			&user.NamaLengkap,
			&user.Alamat,
			&user.NomorHp,
			&user.UserName,
			&user.Password,
			&user.CreateDate,
			&user.CreateBy,
			&user.UpdateDate,
			&user.UpdateBy,
			&user.Status)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		user_ = append(user_, user)
	}
	return user_, nil
}

// Create implements UserRepository.
func (r *UserConnetion) Create(ctx context.Context, input entity.UserEntity) (entity.UserEntity, error) {
	SQL := "INSERT INTO user (id_unit, id_jabatan, nama_lengkap, alamat, nomor_hp, username, password, status) values (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.IdUnit,
		input.IdJabatan,
		input.NamaLengkap,
		input.Alamat,
		input.NomorHp,
		input.UserName,
		input.Password,
		input.Status)
	if err != nil {
		return entity.UserEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.UserEntity{}, err
	}
	input.IdUser = strconv.Itoa(int(id))
	return input, err
}

// Delete implements UserRepository.
func (r *UserConnetion) Delete(ctx context.Context, IdUser string) error {
	SQL := "DELETE FROM berkas WHERE id_user = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdUser)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements UserRepository.
func (r *UserConnetion) FindByID(ctx context.Context, IdUser string) (entity.UserEntity, error) {
	SQL := "SELECT * FROM user WHERE id_user = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdUser)
	if err != nil {
		return entity.UserEntity{}, err
	}
	defer rows.Close()

	user := entity.UserEntity{}
	if rows.Next() {
		err := rows.Scan(
			&user.IdUser,
			&user.IdUnit,
			&user.IdJabatan,
			&user.NamaLengkap,
			&user.Alamat,
			&user.NomorHp,
			&user.UserName,
			&user.Password,
			&user.CreateDate,
			&user.CreateBy,
			&user.UpdateDate,
			&user.UpdateBy,
			&user.Status)
		if err != nil {
			return user, err
		}
		return user, nil
	} else {
		return user, err
	}
}

// Update implements UserRepository.
func (r *UserConnetion) Update(ctx context.Context, input entity.UserEntity) error {
	fmt.Print(input)
	SQL := "UPDATE user SET id_unit = ?, id_jabatan = ?, nama_lengkap = ?, alamat = ?, nomor_hp = ?, username = ?, password = ? WHERE id_user = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.IdUnit,
		input.IdJabatan,
		input.NamaLengkap,
		input.Alamat,
		input.NomorHp,
		input.UserName,
		input.Password,
		input.IdUser)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &UserConnetion{
		tx: DB,
	}
}
