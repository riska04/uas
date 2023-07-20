package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/entity"
	"strconv"
)

type ProgramStudiRepository interface {
	All(ctx context.Context) ([]entity.ProgramStudiEntity, error)
	FindByID(ctx context.Context, IdProgramStudi string) (entity.ProgramStudiEntity, error)
	Create(ctx context.Context, input entity.ProgramStudiEntity) (entity.ProgramStudiEntity, error)
	Update(ctx context.Context, input entity.ProgramStudiEntity) error
	Delete(ctx context.Context, IdProgramStudi string) error
}

type ProgramStudiConnetion struct {
	tx *sql.DB
}

// All implements ProgramStudiRepository.
func (r *ProgramStudiConnetion) All(ctx context.Context) ([]entity.ProgramStudiEntity, error) {
	SQL := "SELECT * FROM program_studi"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	programstudi_ := []entity.ProgramStudiEntity{}
	for rows.Next() {
		programstudi := entity.ProgramStudiEntity{}
		err := rows.Scan(
			&programstudi.ProdiId,
			&programstudi.IdFakultas,
			&programstudi.ProgramStudi,
			&programstudi.Singkatan)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		programstudi_ = append(programstudi_, programstudi)
	}
	return programstudi_, nil
}

// Create implements ProgramStudiRepository.
func (r *ProgramStudiConnetion) Create(ctx context.Context, input entity.ProgramStudiEntity) (entity.ProgramStudiEntity, error) {
	SQL := "INSERT INTO program_studi (id_fakultas, program_studi, singkatan) values (?, ?, ?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.IdFakultas,
		input.ProgramStudi,
		input.Singkatan)
	if err != nil {
		return entity.ProgramStudiEntity{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.ProgramStudiEntity{}, err
	}
	input.ProdiId = strconv.Itoa(int(id))
	return input, err
}

// Delete implements ProgramStudiRepository.
func (r *ProgramStudiConnetion) Delete(ctx context.Context, IdProgramStudi string) error {
	SQL := "DELETE FROM berkas WHERE id_program_studi = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdProgramStudi)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements ProgramStudiRepository.
func (r *ProgramStudiConnetion) FindByID(ctx context.Context, IdProgramStudi string) (entity.ProgramStudiEntity, error) {
	SQL := "SELECT * FROM program_studi WHERE prodi_id = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdProgramStudi)
	if err != nil {
		return entity.ProgramStudiEntity{}, err
	}
	defer rows.Close()

	ProgramStudi := entity.ProgramStudiEntity{}
	if rows.Next() {
		err := rows.Scan(
			&ProgramStudi.ProdiId,
			&ProgramStudi.IdFakultas,
			&ProgramStudi.ProgramStudi,
			&ProgramStudi.Singkatan)

		if err != nil {
			return ProgramStudi, err
		}
		return ProgramStudi, nil
	} else {
		return ProgramStudi, err
	}
}

// Update implements ProgramStudiRepository.
func (r *ProgramStudiConnetion) Update(ctx context.Context, input entity.ProgramStudiEntity) error {
	fmt.Print(input)
	SQL := "UPDATE program_studi SET id_fakultas = ?, program_studi = ?, singkatan = ? WHERE prodi_id = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.ProdiId,
		input.IdFakultas,
		input.ProgramStudi)
	if err != nil {
		return err
	}
	return nil
}

func NewProgramStudiRepository(DB *sql.DB) ProgramStudiRepository {
	return &ProgramStudiConnetion{
		tx: DB,
	}
}
