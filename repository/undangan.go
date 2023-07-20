package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type UndanganRepository interface {
	All(ctx context.Context) ([]entity.UndanganEntity, error)
	FindByID(ctx context.Context, Undangan string) (entity.UndanganEntity, error)
	Create(ctx context.Context, input entity.UndanganEntity) (entity.UndanganEntity, error)
	Update(ctx context.Context, input entity.UndanganEntity) error
	Delete(ctx context.Context, Undangan string) error
}

type UndanganConnetion struct {
	tx *sql.DB
}

// All implements UndanganRepository.
func (r *UndanganConnetion) All(ctx context.Context) ([]entity.UndanganEntity, error) {
	SQL := "SELECT * FROM undangan"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	undangan_ := []entity.UndanganEntity{}
	for rows.Next() {
		undangan := entity.UndanganEntity{}
		err := rows.Scan(
			&undangan.IdUndangan,
			&undangan.IdUnit,
			&undangan.IdJabatan,
			&undangan.IdUser,
			&undangan.IdBerkas)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		undangan_ = append(undangan_, undangan)
	}
	return undangan_, nil
}

// Create implements UndanganRepository.
func (r *UndanganConnetion) Create(ctx context.Context, input entity.UndanganEntity) (entity.UndanganEntity, error) {
	SQL := "INSERT INTO undangan (id_unit, id_jabatan, id_user,id_berkas) values (?, ?, ?, ?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.IdUnit,
		input.IdJabatan,
		input.IdUser,
		input.IdBerkas)
	if err != nil {
		return entity.UndanganEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.UndanganEntity{}, err
	}
	input.IdUndangan = strconv.Itoa(int(id))
	return input, err
}

// Delete implements UndanganRepository.
func (r *UndanganConnetion) Delete(ctx context.Context, IdUndangan string) error {
	SQL := "DELETE FROM berkas WHERE id_undangan = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdUndangan)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements UndanganRepository.
func (r *UndanganConnetion) FindByID(ctx context.Context, IdUndangan string) (entity.UndanganEntity, error) {
	SQL := "SELECT * FROM undangan WHERE id_undangan = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdUndangan)
	if err != nil {
		return entity.UndanganEntity{}, err
	}
	defer rows.Close()

	undangan := entity.UndanganEntity{}
	if rows.Next() {
		err := rows.Scan(
			&undangan.IdUndangan,
			&undangan.IdUnit,
			&undangan.IdJabatan,
			&undangan.IdUser,
			&undangan.IdBerkas)
		if err != nil {
			return undangan, err
		}
		return undangan, nil
	} else {
		return undangan, err
	}
}

// Update implements UndanganRepository.
func (r *UndanganConnetion) Update(ctx context.Context, input entity.UndanganEntity) error {
	fmt.Print(input)
	SQL := "UPDATE undangan SET id_unit = ?, id_jabatan = ?, id_user = ?, id_berkas = ? WHERE id_undangan = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.IdUnit,
		input.IdJabatan,
		input.IdUser,
		input.IdBerkas,
		input.IdUndangan)
	if err != nil {
		return err
	}
	return nil
}

func NewUndanganRepository(DB *sql.DB) UndanganRepository {
	return &UndanganConnetion{
		tx: DB,
	}
}
