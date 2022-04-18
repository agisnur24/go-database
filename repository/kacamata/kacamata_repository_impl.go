package kacamata

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type KacamataRepositoryImpl struct {
	DB *sql.DB
}

func NewKacamataRepository(db *sql.DB) KacamataRepository {
	return &KacamataRepositoryImpl{db}
}

func (repository *KacamataRepositoryImpl) Insert(ctx context.Context, kacamata entity.Kacamata) (entity.Kacamata, error) {
	script := "INSERT INTO kacamata(merk, price) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, kacamata.Merk, kacamata.Price)
	if err != nil {
		return kacamata, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return kacamata, err
	}
	kacamata.Id = int32(id)
	return kacamata, nil
}

func (repository *KacamataRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Kacamata, error) {
	script := "SELECT id FROM kacamata WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	kacamata := entity.Kacamata{}
	if err != nil {
		return kacamata, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&kacamata.Id)
		return kacamata, nil
	} else {
		//tidak ada
		return kacamata, errors.New("Id" + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *KacamataRepositoryImpl) FindAll(ctx context.Context) ([]entity.Kacamata, error) {
	script := "SELECT id, merk, price FROM kacamata"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var glasses []entity.Kacamata
	for rows.Next() {
		kacamata := entity.Kacamata{}
		rows.Scan(&kacamata.Id, &kacamata.Merk, &kacamata.Price)
		glasses = append(glasses, kacamata)
	}
	return glasses, nil
}

func (repository *KacamataRepositoryImpl) Delete(ctx context.Context, kacamata entity.Kacamata) (entity.Kacamata, error) {
	script := "DELETE FROM kacamata WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, kacamata.Id)
	if err != nil {
		return kacamata, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return kacamata, err
	}
	if rows == 0 {
		return kacamata, err
	}
	return kacamata, nil
}

func (repository *KacamataRepositoryImpl) Update(ctx context.Context, id int32, kacamata entity.Kacamata) (entity.Kacamata, error) {
	script := "SELECT id, merk, price FROM kacamata WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return kacamata, err
	}

	if rows.Next() {
		script := "UPDATE kacamata SET price = ?, merk = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, kacamata.Merk, kacamata.Price, id)
		if err != nil {
			return kacamata, err
		}
		kacamata.Id = id
		return kacamata, nil
	} else {
		return kacamata, errors.New(("Id" + strconv.Itoa(int(id)) + " Update Failed"))
	}
}
