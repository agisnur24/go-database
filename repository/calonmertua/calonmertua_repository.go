package calonmertua

import (
	"context"
	"go_database/entity"
)

type CalonmertuaRepository interface {
	Insert(ctx context.Context, calonmertua entity.Calonmertua) (entity.Calonmertua, error)
	FindById(ctx context.Context, id int32) (entity.Calonmertua, error)
	FindAll(ctx context.Context) ([]entity.Calonmertua, error)
	Update(ctx context.Context, id int32, calonmertua entity.Calonmertua) (entity.Calonmertua, error)
	Delete(ctx context.Context, calonmertua entity.Calonmertua) (entity.Calonmertua, error)
}
