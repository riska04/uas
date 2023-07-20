package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type JabatanRepository interface {
	All(ctx context.Context) ([]entity.JabatanEntity, error)
	FindByID(ctx context.Context, IdJabatan string) (entity.JabatanEntity, error)
	Create(ctx context.Context, input entity.JabatanEntity) (entity.JabatanEntity, error)
	Update(ctx context.Context, input entity.JabatanEntity) error
	Delete(ctx context.Context, IdJabatan string) error
}

type JabatanConnetion struct {
	tx *sql.DB
}

// All implements JabatanRepository.
func (r *JabatanConnetion) All(ctx context.Context) ([]entity.JabatanEntity, error) {

	SQL := "SELECT * FROM jabatan"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	jabatan_ := []entity.JabatanEntity{}
	for rows.Next() {
		jabatan := entity.JabatanEntity{}
		err := rows.Scan(
			&jabatan.IdJabatan,
			&jabatan.NamaJabatan,
		)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		jabatan_ = append(jabatan_, jabatan)
	}
	return jabatan_, nil
}

// Create implements JabatanRepository.
func (r *JabatanConnetion) Create(ctx context.Context, input entity.JabatanEntity) (entity.JabatanEntity, error) {
	SQL := "INSERT INTO jabatan (nama_jabatan) values (?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.NamaJabatan)
	if err != nil {
		return entity.JabatanEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.JabatanEntity{}, err
	}
	input.IdJabatan = strconv.Itoa(int(id))
	return input, err
}

// Delete implements JabatanRepository.
func (r *JabatanConnetion) Delete(ctx context.Context, IdJabatan string) error {
	SQL := "DELETE FROM jabatan WHERE id_jabatan = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdJabatan)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements JabatanRepository.
func (r *JabatanConnetion) FindByID(ctx context.Context, IdJabatan string) (entity.JabatanEntity, error) {
	SQL := "SELECT * FROM jabatan WHERE id_jabatan = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdJabatan)
	if err != nil {
		return entity.JabatanEntity{}, err
	}
	defer rows.Close()

	jabatan := entity.JabatanEntity{}
	if rows.Next() {
		err := rows.Scan(
			&jabatan.IdJabatan,
			&jabatan.NamaJabatan)
		if err != nil {
			return jabatan, err
		}
		return jabatan, nil
	} else {
		return jabatan, err
	}
}

// Update implements JabatanRepository.
func (r *JabatanConnetion) Update(ctx context.Context, input entity.JabatanEntity) error {
	fmt.Print(input)
	SQL := "UPDATE jabatan SET nama_jabatan = ? WHERE id_jabatan = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.NamaJabatan,
		input.IdJabatan)
	if err != nil {
		return err
	}
	return nil
}

func NewJabatanRepository(DB *sql.DB) JabatanRepository {
	return &JabatanConnetion{
		tx: DB,
	}
}
