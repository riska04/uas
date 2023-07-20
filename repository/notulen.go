package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type NotulenRepository interface {
	All(ctx context.Context) ([]entity.NotulenEntity, error)
	FindByID(ctx context.Context, Notulen string) (entity.NotulenEntity, error)
	Create(ctx context.Context, input entity.NotulenEntity) (entity.NotulenEntity, error)
	Update(ctx context.Context, input entity.NotulenEntity) error
	Delete(ctx context.Context, IdNotulen string) error
}

type NotulenConnetion struct {
	tx *sql.DB
}

// All implements NotulenRepository.
func (r *NotulenConnetion) All(ctx context.Context) ([]entity.NotulenEntity, error) {
	SQL := "SELECT * FROM notulen"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	notulen_ := []entity.NotulenEntity{}
	for rows.Next() {
		notulen := entity.NotulenEntity{}
		err := rows.Scan(
			&notulen.IdNotulen,
			&notulen.IdBerkas,
			&notulen.Catatan,
			&notulen.CreateDate,
			&notulen.CreateBy,
			&notulen.UpdateDate,
			&notulen.UpdateBy)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		notulen_ = append(notulen_, notulen)
	}
	return notulen_, nil
}

// Create implements NotulenRepository.
func (r *NotulenConnetion) Create(ctx context.Context, input entity.NotulenEntity) (entity.NotulenEntity, error) {
	SQL := "INSERT INTO notulen (id_berkas, catatan) values (?, ?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.IdBerkas,
		input.Catatan)
	if err != nil {
		return entity.NotulenEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.NotulenEntity{}, err
	}
	input.IdNotulen = strconv.Itoa(int(id))
	return input, err
}

// Delete implements NotulenRepository.
func (r *NotulenConnetion) Delete(ctx context.Context, IdNotulen string) error {
	SQL := "DELETE FROM notulen WHERE id_notulen = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdNotulen)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements NotulenRepository.
func (r *NotulenConnetion) FindByID(ctx context.Context, IdNotulen string) (entity.NotulenEntity, error) {
	SQL := "SELECT * FROM notulen WHERE id_notulen = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdNotulen)
	if err != nil {
		return entity.NotulenEntity{}, err
	}
	defer rows.Close()

	notulen := entity.NotulenEntity{}
	if rows.Next() {
		err := rows.Scan(
			&notulen.IdNotulen,
			&notulen.IdBerkas,
			&notulen.Catatan,
			&notulen.CreateDate,
			&notulen.CreateBy,
			&notulen.UpdateDate,
			&notulen.UpdateBy)
		if err != nil {
			return notulen, err
		}
		return notulen, nil
	} else {
		return notulen, err
	}
}

// Update implements NotulenRepository.
func (r *NotulenConnetion) Update(ctx context.Context, input entity.NotulenEntity) error {
	fmt.Print(input)
	SQL := "UPDATE notulen SET id_berkas = ?, catatan = ? WHERE id_notulen = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.IdNotulen,
		input.IdBerkas,
		input.Catatan)
	if err != nil {
		return err
	}
	return nil
}

func NewNotulenRepository(DB *sql.DB) NotulenRepository {
	return &NotulenConnetion{
		tx: DB,
	}
}
