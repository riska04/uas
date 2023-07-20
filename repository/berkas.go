package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type BerkasRepository interface {
	All(ctx context.Context) ([]entity.BerkasEntity, error)
	FindByID(ctx context.Context, IdBerkas string) (entity.BerkasEntity, error)
	Create(ctx context.Context, input entity.BerkasEntity) (entity.BerkasEntity, error)
	Update(ctx context.Context, input entity.BerkasEntity) error
	Delete(ctx context.Context, IdBerkas string) error
}
type BerkasConnetion struct {
	tx *sql.DB
}

// All implements BerkasRepository.
func (r *BerkasConnetion) All(ctx context.Context) ([]entity.BerkasEntity, error) {
	SQL := "SELECT * FROM berkas"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	berkas_ := []entity.BerkasEntity{}
	for rows.Next() {
		berkas := entity.BerkasEntity{}
		err := rows.Scan(
			&berkas.IdBerkas,
			&berkas.File,
			&berkas.NamaFile,
			&berkas.TanggalUpload,
			&berkas.TanggalAcara,
			&berkas.StartJam,
			&berkas.SelesaiJam,
			&berkas.IdUnit,
			&berkas.CreateBy,
			&berkas.CreateDate,
			&berkas.UpdateBy,
			&berkas.UpdateDate)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		berkas_ = append(berkas_, berkas)
	}
	return berkas_, nil
}

// Create implements BerkasRepository.
func (r *BerkasConnetion) Create(ctx context.Context, input entity.BerkasEntity) (entity.BerkasEntity, error) {
	SQL := "INSERT INTO berkas (file, nama_file, tanggal_upload , tanggal_acara, start_jam, selesai_jam, id_unit) values (?, ?, ?, ?, ?, ?, ? )"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.File,
		input.NamaFile,
		input.TanggalUpload,
		input.TanggalAcara,
		input.StartJam,
		input.SelesaiJam,
		input.IdUnit)
	if err != nil {
		return entity.BerkasEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.BerkasEntity{}, err
	}
	input.IdBerkas = strconv.Itoa(int(id))
	return input, err
}

// Delete implements BerkasRepository.
func (r *BerkasConnetion) Delete(ctx context.Context, IdBerkas string) error {
	SQL := "DELETE FROM berkas WHERE id_berkas = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdBerkas)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements BerkasRepository.
func (r *BerkasConnetion) FindByID(ctx context.Context, IdBerkas string) (entity.BerkasEntity, error) {
	SQL := "SELECT * FROM berkas WHERE id_berkas = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdBerkas)
	if err != nil {
		return entity.BerkasEntity{}, err
	}
	defer rows.Close()

	berkas := entity.BerkasEntity{}
	if rows.Next() {
		err := rows.Scan(
			&berkas.IdBerkas,
			&berkas.File,
			&berkas.NamaFile,
			&berkas.TanggalUpload,
			&berkas.TanggalAcara,
			&berkas.StartJam,
			&berkas.SelesaiJam,
			&berkas.IdUnit,
			&berkas.CreateBy,
			&berkas.CreateDate,
			&berkas.UpdateBy,
			&berkas.UpdateDate)
		if err != nil {
			return berkas, err
		}
		return berkas, nil
	} else {
		return berkas, err
	}
}

// Update implements BerkasRepository.
func (r *BerkasConnetion) Update(ctx context.Context, input entity.BerkasEntity) error {
	fmt.Print(input)
	SQL := "UPDATE berkas SET file = ?, nama_file = ?, tanggal_upload = ?, tanggal_acara = ?, start_jam = ?, selesai_jam = ? WHERE id_berkas = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.File,
		input.NamaFile,
		input.TanggalUpload,
		input.TanggalAcara,
		input.StartJam,
		input.SelesaiJam,
		input.IdBerkas)
	if err != nil {
		return err
	}
	return nil
}

func NewBerkasRepository(DB *sql.DB) BerkasRepository {
	return &BerkasConnetion{
		tx: DB,
	}
}
