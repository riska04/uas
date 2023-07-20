package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type StrukturRepository interface {
	All(ctx context.Context) ([]entity.StrukturEntity, error)
	FindByID(ctx context.Context, Struktur string) (entity.StrukturEntity, error)
	Create(ctx context.Context, input entity.StrukturEntity) (entity.StrukturEntity, error)
	Update(ctx context.Context, input entity.StrukturEntity) error
	Delete(ctx context.Context, Struktur string) error
}

type StrukturConnetion struct {
	tx *sql.DB
}

// All implements StrukturRepository.
func (r *StrukturConnetion) All(ctx context.Context) ([]entity.StrukturEntity, error) {
	SQL := "SELECT * FROM struktur"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	struktur_ := []entity.StrukturEntity{}
	for rows.Next() {
		struktur := entity.StrukturEntity{}
		err := rows.Scan(
			&struktur.IdStruktur,
			&struktur.NamaStruktur)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		struktur_ = append(struktur_, struktur)
	}
	return struktur_, nil
}

// Create implements StrukturRepository.
func (r *StrukturConnetion) Create(ctx context.Context, input entity.StrukturEntity) (entity.StrukturEntity, error) {
	SQL := "INSERT INTO struktur (nama_struktur) values (?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.NamaStruktur)
	if err != nil {
		return entity.StrukturEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.StrukturEntity{}, err
	}
	input.IdStruktur = strconv.Itoa(int(id))
	return input, err
}

// Delete implements StrukturRepository.
func (r *StrukturConnetion) Delete(ctx context.Context, IdStruktur string) error {
	SQL := "DELETE FROM berkas WHERE id_struktur = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdStruktur)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements StrukturRepository.
func (r *StrukturConnetion) FindByID(ctx context.Context, IdStruktur string) (entity.StrukturEntity, error) {
	SQL := "SELECT * FROM struktur WHERE id_struktur = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdStruktur)
	if err != nil {
		return entity.StrukturEntity{}, err
	}
	defer rows.Close()

	struktur := entity.StrukturEntity{}
	if rows.Next() {
		err := rows.Scan(
			&struktur.IdStruktur,
			&struktur.NamaStruktur)
		if err != nil {
			return struktur, err
		}
		return struktur, nil
	} else {
		return struktur, err
	}
}

// Update implements StrukturRepository.
func (r *StrukturConnetion) Update(ctx context.Context, input entity.StrukturEntity) error {
	fmt.Print(input)
	SQL := "UPDATE struktur SET nama_struktur = ? WHERE id_struktur = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.IdStruktur,
		input.NamaStruktur)
	if err != nil {
		return err
	}
	return nil
}

func NewStrukturRepository(DB *sql.DB) StrukturRepository {
	return &StrukturConnetion{
		tx: DB,
	}
}
