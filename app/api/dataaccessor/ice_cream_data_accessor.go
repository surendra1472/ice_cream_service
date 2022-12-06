package dataaccessor

import (
	"context"
	"github.com/go-pg/pg"
	"ic-service/app/config"
	"ic-service/app/model/bo"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=IcecreamDataAccessor"
type IcecreamDataAccessor interface {
	Insert(context.Context, *bo.Icecream) (*bo.Icecream, error)
	GetById(context.Context, *bo.Icecream) (*bo.Icecream, error)
	GetByProductId(context.Context, *bo.Icecream) (*bo.Icecream, error)
	DeleteByProductId(context.Context, *bo.Icecream) error
	TxInsert(context.Context, *pg.Tx, *bo.Icecream) (*bo.Icecream, *pg.Tx, error)
	TxUpdate(context.Context, *pg.Tx, *bo.Icecream) (*bo.Icecream, *pg.Tx, error)
}

func NewIcecreamDataAccessor() IcecreamDataAccessor {
	return &icecreamDataAccessorService{db: config.GetDBConnection()}
}

type icecreamDataAccessorService struct {
	db *pg.DB
}

func (idas icecreamDataAccessorService) GetByProductId(ctx context.Context, icecream *bo.Icecream) (*bo.Icecream, error) {
	err := idas.db.Model(icecream).
		Where("product_id = ?", icecream.ProductId).Select()
	return icecream, err
}

func (idas icecreamDataAccessorService) GetById(ctx context.Context, icecream *bo.Icecream) (*bo.Icecream, error) {
	err := idas.db.Model(icecream).
		Where("id = ?", icecream.Id).Select()
	return icecream, err
}

func (idas icecreamDataAccessorService) DeleteByProductId(ctx context.Context, icecream *bo.Icecream) error {
	_, err := idas.db.Model(icecream).
		Where("product_id = ?", icecream.ProductId).Delete()
	return err
}

func (idas icecreamDataAccessorService) Insert(ctx context.Context, icecream *bo.Icecream) (*bo.Icecream, error) {
	err := idas.db.Insert(icecream)
	return icecream, err
}

func (idas icecreamDataAccessorService) TxInsert(ctx context.Context, tx *pg.Tx, icecream *bo.Icecream) (*bo.Icecream, *pg.Tx, error) {
	return icecream, tx, tx.Insert(icecream)
}

func (idas icecreamDataAccessorService) TxUpdate(ctx context.Context, tx *pg.Tx, icecream *bo.Icecream) (*bo.Icecream, *pg.Tx, error) {
	return icecream, tx, tx.Update(icecream)
}
