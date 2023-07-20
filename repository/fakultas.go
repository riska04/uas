package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type FakultasRepository interface {
	All(ctx context.Context) ([]entity.FakultasEntity, error)
	FindByID(ctx context.Context, Fakultas string) (entity.FakultasEntity, error)
	Create(ctx context.Context, input entity.FakultasEntity) (entity.FakultasEntity, error)
	Update(ctx context.Context, input entity.FakultasEntity) error
	Delete(ctx context.Context, IdFakultas string) error
}

type FakultasConnetion struct {
	tx *sql.DB
}

// All implements FakultasRepository.
func (r *FakultasConnetion) All(ctx context.Context) ([]entity.FakultasEntity, error) {
	SQL := "SELECT * FROM fakultas"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fakultas_ := []entity.FakultasEntity{}
	for rows.Next() {
		fakultas := entity.FakultasEntity{}
		err := rows.Scan(
			&fakultas.IdFakultas,
			&fakultas.NamaFakultas,
			&fakultas.SingkatanFakultas)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		fakultas_ = append(fakultas_, fakultas)
	}
	return fakultas_, nil
}

// Create implements FakultasRepository.
func (r *FakultasConnetion) Create(ctx context.Context, input entity.FakultasEntity) (entity.FakultasEntity, error) {
	SQL := "INSERT INTO fakultas (nama_fakultas,singkatan_fakultas) values (?, ?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.NamaFakultas,
		input.SingkatanFakultas,
	)
	if err != nil {
		return entity.FakultasEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.FakultasEntity{}, err
	}
	input.IdFakultas = strconv.Itoa(int(id))
	return input, err
}

// Delete implements FakultasRepository.
func (r *FakultasConnetion) Delete(ctx context.Context, IdFakultas string) error {
	SQL := "DELETE FROM fakultas WHERE id_fakultas = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdFakultas)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements FakultasRepository.
func (r *FakultasConnetion) FindByID(ctx context.Context, IdFakultas string) (entity.FakultasEntity, error) {
	SQL := "SELECT * FROM fakultas WHERE id_fakultas = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdFakultas)
	if err != nil {
		return entity.FakultasEntity{}, err
	}
	defer rows.Close()

	fakultas := entity.FakultasEntity{}
	if rows.Next() {
		err := rows.Scan(
			&fakultas.IdFakultas,
			&fakultas.NamaFakultas,
			&fakultas.SingkatanFakultas)
		if err != nil {
			return fakultas, err
		}
		return fakultas, nil
	} else {
		return fakultas, err
	}
}

// Update implements FakultasRepository.
func (r *FakultasConnetion) Update(ctx context.Context, input entity.FakultasEntity) error {
	fmt.Print(input)
	SQL := "UPDATE fakultas SET nama_fakultas = ?, singkatan_fakultas = ? WHERE id_fakultas = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.NamaFakultas,
		input.SingkatanFakultas,
		input.IdFakultas)
	if err != nil {
		return err
	}
	return nil
}

func NewFakultasRepository(DB *sql.DB) FakultasRepository {
	return &FakultasConnetion{
		tx: DB,
	}
}
