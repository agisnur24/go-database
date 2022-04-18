package kacamata

import (
	"context"
	"go_database/entity"
)

type KacamataRepository interface {
	Insert(ctx context.Context, kacamata entity.Kacamata) (entity.Kacamata, error)
	FindById(ctx context.Context, id int32) (entity.Kacamata, error)
	FindAll(ctx context.Context) ([]entity.Kacamata, error)
	Update(ctx context.Context, id int32, kacamata entity.Kacamata) (entity.Kacamata, error)
	Delete(ctx context.Context, kacamata entity.Kacamata) (entity.Kacamata, error)
}
