package motor

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type MotorRepositoryImpl struct {
	DB *sql.DB
}

func NewMotorRepository(db *sql.DB) MotorRepository {
	return &MotorRepositoryImpl{db}
}

func (repository *MotorRepositoryImpl) Insert(ctx context.Context, motor entity.Motor) (entity.Motor, error) {
	script := "INSERT INTO harga_motor(merk, transmition, price) VALUES (?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, motor.Merk, motor.Transmition, motor.Price)
	if err != nil {
		return motor, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return motor, err
	}
	motor.Id = int32(id)
	return motor, nil
}

func (repository *MotorRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Motor, error) {
	script := "SELECT id FROM harga_motor WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	motor := entity.Motor{}
	if err != nil {
		return motor, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&motor.Id)
		return motor, nil
	} else {
		//tidak ada
		return motor, errors.New("Id" + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *MotorRepositoryImpl) FindAll(ctx context.Context) ([]entity.Motor, error) {
	script := "SELECT id, merk, transmition, price FROM harga_motor"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var motors []entity.Motor
	for rows.Next() {
		motor := entity.Motor{}
		rows.Scan(&motor.Id, &motor.Merk, &motor.Transmition, &motor.Price)
		motors = append(motors, motor)
	}
	return motors, nil
}

func (repository *MotorRepositoryImpl) Delete(ctx context.Context, motor entity.Motor) (entity.Motor, error) {
	script := "DELETE FROM harga_motor WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, motor.Id)
	if err != nil {
		return motor, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return motor, err
	}
	if rows == 0 {
		return motor, err
	}
	return motor, nil
}

func (repository *MotorRepositoryImpl) Update(ctx context.Context, id int32, motor entity.Motor) (entity.Motor, error) {
	script := "SELECT id, merk, transmition, price FROM harga_motor WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return motor, err
	}

	if rows.Next() {
		script := "UPDATE harga_motor SET merk = ?, transmition = ?, price = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, motor.Merk, motor.Transmition, motor.Price, id)
		if err != nil {
			return motor, err
		}
		motor.Id = id
		return motor, nil
	} else {
		return motor, errors.New(("Id" + strconv.Itoa(int(id)) + " Update Failed"))
	}
}
