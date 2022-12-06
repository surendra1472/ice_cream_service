package processor

import (
	"context"
	"errors"
	"github.com/go-pg/pg"
	is "ic-indexer-service/app/model/request"
	"ic-service/app/api/dataaccessor"
	"ic-service/app/builder"
	"ic-service/app/model/bo"
	"ic-service/app/model/request"
	"ic-service/app/processor/handler"
	"log"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=IcecreamService"
type IcecreamService interface {
	Save(ctx context.Context, icecreamRequest request.IcecreamRequest) (*bo.Icecream, error)
	PartialUpdate(ctx context.Context, icecreamUpdateRequest request.IcecreamUpdateRequest) (*bo.Icecream, error)
	Delete(ctx context.Context, icecreamRequest request.IcecreamRequest) (*bo.Icecream, error)
}

type icecreamService struct {
	db           *pg.DB
	idas         dataaccessor.IcecreamDataAccessor
	ib           builder.IcecreamBuilderInterface
	kafkaHandler handler.SendMessageToKafkaHandler
}

func GetNewIcecreamService(db *pg.DB, idas dataaccessor.IcecreamDataAccessor) *icecreamService {
	return &icecreamService{

		db:           db,
		idas:         idas,
		ib:           builder.NewIcecreamBuilder(),
		kafkaHandler: handler.NewSendMessageToKafkaHandler(),
	}
}

func (is icecreamService) Save(ctx context.Context, icecreamRequest request.IcecreamRequest) (*bo.Icecream, error) {

	var err error

	icecream := &bo.Icecream{}
	err = is.db.RunInTransaction(func(tx *pg.Tx) error {

		icecream = is.ib.IcecreamBuilder(icecreamRequest)
		icecream, err = is.idas.GetByProductId(ctx, icecream)

		if err != nil && err == pg.ErrNoRows {
			icecream, tx, err = is.idas.TxInsert(ctx, tx, icecream)
			if err != nil {
				log.Print(ctx, "error creating icecream")
				return err
			}
		} else {
			log.Print(ctx, "icecream already exists")
			return errors.New("icecream already exists")
		}

		err = is.sendMessageToKafka(ctx, icecream, false)

		return err
	})

	return icecream, err
}

func (is icecreamService) PartialUpdate(ctx context.Context, icecreamUpdateRequest request.IcecreamUpdateRequest) (*bo.Icecream, error) {

	var err error

	updatedIcecream := &bo.Icecream{}
	err = is.db.RunInTransaction(func(tx *pg.Tx) error {

		if !icecreamUpdateRequest.Id.Set || icecreamUpdateRequest.Id.Value == nil {
			return errors.New("cannot update!")
		}

		oldIcecream, err := is.idas.GetById(ctx, &bo.Icecream{Id: *icecreamUpdateRequest.Id.Value})
		if err != nil {
			return err
		}

		updatedIcecream = is.ib.IcecreamPartialBuilder(icecreamUpdateRequest, oldIcecream)
		_, _, err = is.idas.TxUpdate(ctx, tx, updatedIcecream)

		if err != nil {
			log.Print(ctx, "icecream updation failed")
			return err
		}

		err = is.sendMessageToKafka(ctx, updatedIcecream, false)

		return err
	})

	return updatedIcecream, err
}

func (is icecreamService) Delete(ctx context.Context, icecreamDelete is.IcecreamDelete) (*bo.Icecream, error) {

	var err error

	icecream := &bo.Icecream{}
	err = is.db.RunInTransaction(func(tx *pg.Tx) error {

		icecream = is.ib.IcecreamBuilder(request.IcecreamRequest{ProductId: icecreamDelete.ProductId})
		icecream, err = is.idas.GetById(ctx, icecream)
		if err == nil {
			err = is.idas.DeleteByProductId(ctx, icecream)
		}
		if err != nil {
			log.Print(ctx, "icecream deletion failed")
			return err
		}

		err = is.sendMessageToKafka(ctx, icecream, true)

		return err
	})

	return icecream, err
}

func (is icecreamService) sendMessageToKafka(ctx context.Context, icecream *bo.Icecream, isDeleted bool) error {

	return is.kafkaHandler.SendMessage(ctx, icecream, isDeleted)
}
