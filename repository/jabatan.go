package repository

import (
	"context"
	"database/sql"
	"materi/model/entity"
)

type JabatanRepository interface {
	All(ctx context.Context) ([]entity.JabatanEntity, error)
	FindByID(ctx context.Context, IdUser string) (entity.JabatanEntity, error)
	Create(ctx context.Context, input entity.JabatanEntity) (entity.JabatanEntity, error)
	Update(ctx context.Context, input entity.JabatanEntity) error
	Delete(ctx context.Context, IdUser string) error
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
	panic("unimplemented")
}

// Delete implements JabatanRepository.
func (r *JabatanConnetion) Delete(ctx context.Context, IdUser string) error {
	panic("unimplemented")
}

// FindByID implements JabatanRepository.
func (r *JabatanConnetion) FindByID(ctx context.Context, IdUser string) (entity.JabatanEntity, error) {
	panic("unimplemented")
}

// Update implements JabatanRepository.
func (r *JabatanConnetion) Update(ctx context.Context, input entity.JabatanEntity) error {
	panic("unimplemented")
}

func NewJabatanRepository(DB *sql.DB) JabatanRepository {
	return &JabatanConnetion{
		tx: DB,
	}
}
