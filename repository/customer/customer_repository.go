package customer

import (
	"context"
	"go_database/entity"
)

type CustomerRepository interface {
	Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	FindById(ctx context.Context, id int32) (entity.Customer, error)
	FindAll(ctx context.Context) ([]entity.Customer, error)
	Update(ctx context.Context, id int32, customer entity.Customer) (entity.Customer, error)
	Delete(ctx context.Context, customer entity.Customer) (entity.Customer, error)
}
