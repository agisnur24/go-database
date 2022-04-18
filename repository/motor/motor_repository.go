package motor

import (
	"context"
	"go_database/entity"
)

type MotorRepository interface {
	Insert(ctx context.Context, harga_motor entity.Motor) (entity.Motor, error)
	FindById(ctx context.Context, id int32) (entity.Motor, error)
	FindAll(ctx context.Context) ([]entity.Motor, error)
	Update(ctx context.Context, id int32, harga_motor entity.Motor) (entity.Motor, error)
	Delete(ctx context.Context, harga_motor entity.Motor) (entity.Motor, error)
}
