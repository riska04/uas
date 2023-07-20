package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type LembagaRepository interface {
	All(ctx context.Context) ([]entity.LembagaEntity, error)
	FindByID(ctx context.Context, Lembaga string) (entity.LembagaEntity, error)
	Create(ctx context.Context, input entity.LembagaEntity) (entity.LembagaEntity, error)
	Update(ctx context.Context, input entity.LembagaEntity) error
	Delete(ctx context.Context, Lembaga string) error
}

type LembagaConnetion struct {
	tx *sql.DB
}

// All implements LembagaRepository.
func (r *LembagaConnetion) All(ctx context.Context) ([]entity.LembagaEntity, error) {
	SQL := "SELECT * FROM lembaga"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lembaga_ := []entity.LembagaEntity{}
	for rows.Next() {
		lembaga := entity.LembagaEntity{}
		err := rows.Scan(
			&lembaga.IdLembaga,
			&lembaga.NamaLembaga,
			&lembaga.SingkatanLembaga)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		lembaga_ = append(lembaga_, lembaga)
	}
	return lembaga_, nil
}

// Create implements LembagaRepository.
func (r *LembagaConnetion) Create(ctx context.Context, input entity.LembagaEntity) (entity.LembagaEntity, error) {
	SQL := "INSERT INTO lembaga (nama_lembaga) values (?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.NamaLembaga)
	if err != nil {
		return entity.LembagaEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.LembagaEntity{}, err
	}
	input.IdLembaga = strconv.Itoa(int(id))
	return input, err
}

// Delete implements LembagaRepository.
func (r *LembagaConnetion) Delete(ctx context.Context, IdLembaga string) error {
	SQL := "DELETE FROM lembaga WHERE id_lembaga = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdLembaga)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements LembagaRepository.
func (r *LembagaConnetion) FindByID(ctx context.Context, IdLembaga string) (entity.LembagaEntity, error) {
	SQL := "SELECT * FROM lembaga WHERE id_lembaga = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdLembaga)
	if err != nil {
		return entity.LembagaEntity{}, err
	}
	defer rows.Close()

	lembaga := entity.LembagaEntity{}
	if rows.Next() {
		err := rows.Scan(
			&lembaga.IdLembaga,
			&lembaga.NamaLembaga,
			&lembaga.SingkatanLembaga)
		if err != nil {
			return lembaga, err
		}
		return lembaga, nil
	} else {
		return lembaga, err
	}
}

// Update implements LembagaRepository.
func (r *LembagaConnetion) Update(ctx context.Context, input entity.LembagaEntity) error {
	fmt.Print(input)
	SQL := "UPDATE lembaga SET nama_lembaga = ?, singkatan_lembaga = ? WHERE id_lembaga = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.NamaLembaga,
		input.SingkatanLembaga,
		input.IdLembaga)
	if err != nil {
		return err
	}
	return nil
}

func NewLembagaRepository(DB *sql.DB) LembagaRepository {
	return &LembagaConnetion{
		tx: DB,
	}
}
