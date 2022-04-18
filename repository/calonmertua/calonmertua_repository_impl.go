package calonmertua

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type CalonmertuaRepositoryImpl struct {
	DB *sql.DB
}

func NewCalonmertuaRepository(db *sql.DB) CalonmertuaRepository {
	return &CalonmertuaRepositoryImpl{db}
}

func (repository *CalonmertuaRepositoryImpl) Insert(ctx context.Context, calonmertua entity.Calonmertua) (entity.Calonmertua, error) {
	script := "INSERT INTO list_calonmertua(nama, alamat, suku) VALUES (?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, calonmertua.Nama, calonmertua.Alamat, calonmertua.Suku)
	if err != nil {
		return calonmertua, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return calonmertua, err
	}
	calonmertua.Id = int32(id)
	return calonmertua, nil
}

func (repository *CalonmertuaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Calonmertua, error) {
	script := "SELECT id FROM list_calonmertua WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	calonmertua := entity.Calonmertua{}
	if err != nil {
		return calonmertua, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&calonmertua.Id)
		return calonmertua, nil
	} else {
		//tidak ada
		return calonmertua, errors.New("Id" + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *CalonmertuaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Calonmertua, error) {
	script := "SELECT id, nama, alamat, suku FROM list_calonmertua"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var kandidat []entity.Calonmertua
	for rows.Next() {
		calonmertua := entity.Calonmertua{}
		rows.Scan(&calonmertua.Id, &calonmertua.Nama, &calonmertua.Alamat, &calonmertua.Suku)
		kandidat = append(kandidat, calonmertua)
	}
	return kandidat, nil
}

func (repository *CalonmertuaRepositoryImpl) Delete(ctx context.Context, calonmertua entity.Calonmertua) (entity.Calonmertua, error) {
	script := "DELETE FROM list_calonmertua WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, calonmertua.Id)
	if err != nil {
		return calonmertua, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return calonmertua, err
	}
	if rows == 0 {
		return calonmertua, err
	}
	return calonmertua, nil
}

func (repository *CalonmertuaRepositoryImpl) Update(ctx context.Context, id int32, calonmertua entity.Calonmertua) (entity.Calonmertua, error) {
	script := "SELECT id, nama, alamat, suku FROM list_calonmertua WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return calonmertua, err
	}

	if rows.Next() {
		script := "UPDATE list_calonmertua SET nama = ?, alamat = ?, suku = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, calonmertua.Nama, calonmertua.Alamat, calonmertua.Suku, id)
		if err != nil {
			return calonmertua, err
		}
		calonmertua.Id = id
		return calonmertua, nil
	} else {
		return calonmertua, errors.New(("Id" + strconv.Itoa(int(id)) + " Update Failed"))
	}
}
