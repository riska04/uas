package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type UnitRepository interface {
	All(ctx context.Context) ([]entity.UnitEntity, error)
	FindByID(ctx context.Context, Unit string) (entity.UnitEntity, error)
	Create(ctx context.Context, input entity.UnitEntity) (entity.UnitEntity, error)
	Update(ctx context.Context, input entity.UnitEntity) error
	Delete(ctx context.Context, Unit string) error
}

type UnitConnetion struct {
	tx *sql.DB
}

// All implements UnitRepository.
func (r *UnitConnetion) All(ctx context.Context) ([]entity.UnitEntity, error) {
	SQL := "SELECT * FROM unit"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	unit_ := []entity.UnitEntity{}
	for rows.Next() {
		unit := entity.UnitEntity{}
		err := rows.Scan(
			&unit.IdUnit,
			&unit.IdLembaga,
			&unit.IdFakultas,
			&unit.IdProdi,
			&unit.IdStruktur,
			&unit.Status)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		unit_ = append(unit_, unit)
	}
	return unit_, nil
}

// Create implements UnitRepository.
func (r *UnitConnetion) Create(ctx context.Context, input entity.UnitEntity) (entity.UnitEntity, error) {
	SQL := "INSERT INTO unit (id_lembaga, id_struktur, status) values (?, ?, ?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.IdLembaga,
		input.IdStruktur,
		input.Status)
	if err != nil {
		return entity.UnitEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.UnitEntity{}, err
	}
	input.IdUnit = strconv.Itoa(int(id))
	return input, err
}

// Delete implements UnitRepository.
func (r *UnitConnetion) Delete(ctx context.Context, IdUnit string) error {
	SQL := "DELETE FROM unit WHERE id_unit = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdUnit)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements UnitRepository.
func (r *UnitConnetion) FindByID(ctx context.Context, IdUnit string) (entity.UnitEntity, error) {
	SQL := "SELECT * FROM unit WHERE id_unit = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdUnit)
	if err != nil {
		return entity.UnitEntity{}, err
	}
	defer rows.Close()

	unit := entity.UnitEntity{}
	if rows.Next() {
		err := rows.Scan(
			&unit.IdUnit,
			&unit.IdLembaga,
			&unit.IdFakultas,
			&unit.IdProdi,
			&unit.IdStruktur,
			&unit.Status)
		if err != nil {
			return unit, err
		}
		return unit, nil
	} else {
		return unit, err
	}
}

// Update implements UnitRepository.
func (r *UnitConnetion) Update(ctx context.Context, input entity.UnitEntity) error {
	fmt.Print(input)
	SQL := "UPDATE unit SET id_lembaga = ?, id_fakultas = ?, id_prodi = ?, id_struktur = ?, status = ? WHERE id_unit = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.IdUnit,
		input.IdLembaga,
		input.IdFakultas,
		input.IdProdi,
		input.IdStruktur,
		input.Status)
	if err != nil {
		return err
	}
	return nil
}

func NewUnitRepository(DB *sql.DB) UnitRepository {
	return &UnitConnetion{
		tx: DB,
	}
}
