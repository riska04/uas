package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type AbsenRepository interface {
	All(ctx context.Context) ([]entity.AbsenEntity, error)
	FindByID(ctx context.Context, IdAbsen string) (entity.AbsenEntity, error)
	Create(ctx context.Context, input entity.AbsenEntity) (entity.AbsenEntity, error)
	Update(ctx context.Context, input entity.AbsenEntity) error
	Delete(ctx context.Context, IdAbsen string) error
}
type AbsenConnetion struct {
	tx *sql.DB
}

// All implements AbsenRepository.
func (r *AbsenConnetion) All(ctx context.Context) ([]entity.AbsenEntity, error) {
	SQL := "SELECT * FROM absen"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	absen_ := []entity.AbsenEntity{}
	for rows.Next() {
		absen := entity.AbsenEntity{}
		err := rows.Scan(
			&absen.IdAbsen,
			&absen.IdBerkas,
			&absen.IdUnit,
			&absen.IdJabatan,
			&absen.IdUser,
			&absen.CreatedAt,
			&absen.UpdatedAt)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		absen_ = append(absen_, absen)
	}
	return absen_, nil
}

// Create implements AbsenRepository.
func (r *AbsenConnetion) Create(ctx context.Context, input entity.AbsenEntity) (entity.AbsenEntity, error) {
	SQL := "INSERT INTO absen (id_berkas,id_unit,id_jabatan, id_user) values (?, ?, ?, ?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.IdBerkas,
		input.IdUnit,
		input.IdJabatan,
		input.IdUser)
	if err != nil {
		return entity.AbsenEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.AbsenEntity{}, err
	}
	input.IdAbsen = strconv.Itoa(int(id))
	return input, err
}

// Delete implements AbsenRepository.
func (r *AbsenConnetion) Delete(ctx context.Context, IdAbsen string) error {
	SQL := "DELETE FROM absen WHERE id_absen = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdAbsen)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements AbsenRepository.
func (r *AbsenConnetion) FindByID(ctx context.Context, IdAbsen string) (entity.AbsenEntity, error) {
	SQL := "SELECT * FROM absen WHERE id_absen = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdAbsen)
	if err != nil {
		return entity.AbsenEntity{}, err
	}
	defer rows.Close()

	absen := entity.AbsenEntity{}
	if rows.Next() {
		err := rows.Scan(
			&absen.IdAbsen,
			&absen.IdBerkas,
			&absen.IdUnit,
			&absen.IdJabatan,
			&absen.IdUser,
			&absen.CreatedAt,
			&absen.UpdatedAt)
		if err != nil {
			return absen, err
		}
		return absen, nil
	} else {
		return absen, err
	}
}

// Update implements AbsenRepository.
func (r *AbsenConnetion) Update(ctx context.Context, input entity.AbsenEntity) error {
	fmt.Print(input)
	SQL := "UPDATE absen  SET id_berkas = ?, id_unit =?, id_jabatan = ?, id_user =?  WHERE id_absen = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.IdBerkas,
		input.IdUnit,
		input.IdJabatan,
		input.IdUser,
		input.IdAbsen)
	if err != nil {
		return err
	}
	return nil
}

func NewAbsenRepository(DB *sql.DB) AbsenRepository {
	return &AbsenConnetion{
		tx: DB,
	}
}
