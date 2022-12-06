package processor

import (
	"context"
	"errors"
	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	indexer "ic-indexer-service/app/model/request"
	dasMock "ic-service/app/api/dataaccessor/mocks"
	ibMock "ic-service/app/builder/mocks"
	"ic-service/app/config"
	"ic-service/app/model/bo"
	"ic-service/app/model/request"
	kMock "ic-service/app/processor/handler/mocks"
	"testing"
)

func TestNewIcecreamServiceConstructor(t *testing.T) {

	config.InitializeTestConfig()

	daMock := new(dasMock.IcecreamDataAccessor)
	kaMock := new(kMock.SendMessageToKafkaHandler)

	is := GetNewIcecreamService(nil, daMock)
	is.kafkaHandler = kaMock
	assert.NotNil(t, is.idas)

}

func TestNewIcecereamService_IcecreamExists(t *testing.T) {

	config.InitializeTestConfig()

	daMock := new(dasMock.IcecreamDataAccessor)
	builderMock := new(ibMock.IcecreamBuilderInterface)

	daMock.On("GetByProductId", mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil).Times(1)
	builderMock.On("IcecreamBuilder", mock.Anything).Return(&bo.Icecream{}).Times(1)
	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock

	_, err := is.Save(context.TODO(), request.IcecreamRequest{ProductId: "p1", Ingredients: []string{"i1"}})

	assert.NotNil(t, err)

}

func TestIcecreamSuccess(t *testing.T) {

	config.InitializeTestConfig()

	builderMock := new(ibMock.IcecreamBuilderInterface)
	kaMock := new(kMock.SendMessageToKafkaHandler)
	daMock := new(dasMock.IcecreamDataAccessor)

	daMock.On("GetByProductId", mock.Anything, mock.Anything).Return(&bo.Icecream{}, pg.ErrNoRows).Times(1)
	daMock.On("TxInsert", mock.Anything, mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil, nil).Times(1)
	kaMock.On("SendMessage", mock.Anything, mock.Anything, mock.Anything).Return(nil).Times(1)
	builderMock.On("IcecreamBuilder", mock.Anything).Return(&bo.Icecream{}, nil, nil).Times(1)
	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock
	is.kafkaHandler = kaMock

	_, err := is.Save(context.TODO(), request.IcecreamRequest{ProductId: "p1", Ingredients: []string{"i1"}})

	assert.Nil(t, err)

}

func TestIcecreamInsertFailure(t *testing.T) {

	config.InitializeTestConfig()

	builderMock := new(ibMock.IcecreamBuilderInterface)
	kaMock := new(kMock.SendMessageToKafkaHandler)
	daMock := new(dasMock.IcecreamDataAccessor)

	daMock.On("GetByProductId", mock.Anything, mock.Anything).Return(&bo.Icecream{}, pg.ErrNoRows).Times(1)
	daMock.On("TxInsert", mock.Anything, mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil, errors.New("generic error")).Times(1)
	kaMock.On("SendMessage", mock.Anything, mock.Anything, mock.Anything).Return(nil).Times(1)
	builderMock.On("IcecreamBuilder", mock.Anything).Return(&bo.Icecream{}, nil, nil).Times(1)
	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock
	is.kafkaHandler = kaMock

	_, err := is.Save(context.TODO(), request.IcecreamRequest{ProductId: "p1", Ingredients: []string{"i1"}})

	assert.NotNil(t, err)

}

func TestIcecreamDelete(t *testing.T) {

	config.InitializeTestConfig()

	builderMock := new(ibMock.IcecreamBuilderInterface)
	kaMock := new(kMock.SendMessageToKafkaHandler)
	daMock := new(dasMock.IcecreamDataAccessor)

	daMock.On("GetById", mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil).Times(1)
	daMock.On("DeleteByProductId", mock.Anything, mock.Anything).Return(nil).Times(1)
	kaMock.On("SendMessage", mock.Anything, mock.Anything, mock.Anything).Return(nil).Times(1)
	builderMock.On("IcecreamBuilder", mock.Anything).Return(&bo.Icecream{}).Times(1)
	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock
	is.kafkaHandler = kaMock

	_, err := is.Delete(context.TODO(), indexer.IcecreamDelete{})

	assert.Nil(t, err)

}

func TestIcecreamDeleteFail(t *testing.T) {

	config.InitializeTestConfig()

	builderMock := new(ibMock.IcecreamBuilderInterface)
	daMock := new(dasMock.IcecreamDataAccessor)

	daMock.On("GetById", mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil).Times(1)
	daMock.On("DeleteByProductId", mock.Anything, mock.Anything).Return(errors.New("generic error")).Times(1)
	builderMock.On("IcecreamBuilder", mock.Anything).Return(&bo.Icecream{}).Times(1)
	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock

	_, err := is.Delete(context.TODO(), indexer.IcecreamDelete{})

	assert.NotNil(t, err)

}

func TestPartialUpdate(t *testing.T) {

	config.InitializeTestConfig()

	builderMock := new(ibMock.IcecreamBuilderInterface)
	kaMock := new(kMock.SendMessageToKafkaHandler)
	daMock := new(dasMock.IcecreamDataAccessor)

	daMock.On("GetById", mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil).Times(1)
	daMock.On("TxUpdate", mock.Anything, mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil, nil).Times(1)
	kaMock.On("SendMessage", mock.Anything, mock.Anything, mock.Anything).Return(nil).Times(1)
	builderMock.On("IcecreamPartialBuilder", mock.Anything, mock.Anything).Return(&bo.Icecream{}).Times(1)
	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock
	is.kafkaHandler = kaMock

	var id int64
	id = 1
	_, err := is.PartialUpdate(context.TODO(), request.IcecreamUpdateRequest{Id: *request.NewCusInt64(&id)})

	assert.Nil(t, err)

}

func TestPartialUpdate_OldIcecreamNotPresent(t *testing.T) {

	config.InitializeTestConfig()

	builderMock := new(ibMock.IcecreamBuilderInterface)
	kaMock := new(kMock.SendMessageToKafkaHandler)
	daMock := new(dasMock.IcecreamDataAccessor)

	daMock.On("GetById", mock.Anything, mock.Anything).Return(&bo.Icecream{}, errors.New("generic error")).Times(1)

	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock
	is.kafkaHandler = kaMock

	var id int64
	id = 1
	_, err := is.PartialUpdate(context.TODO(), request.IcecreamUpdateRequest{Id: *request.NewCusInt64(&id)})

	assert.NotNil(t, err)

}

func TestPartialUpdate_Failure(t *testing.T) {

	config.InitializeTestConfig()

	builderMock := new(ibMock.IcecreamBuilderInterface)
	kaMock := new(kMock.SendMessageToKafkaHandler)
	daMock := new(dasMock.IcecreamDataAccessor)

	daMock.On("GetById", mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil).Times(1)
	builderMock.On("IcecreamPartialBuilder", mock.Anything, mock.Anything).Return(&bo.Icecream{}).Times(1)
	daMock.On("TxUpdate", mock.Anything, mock.Anything, mock.Anything).Return(&bo.Icecream{}, nil, errors.New("generic error")).Times(1)

	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock
	is.kafkaHandler = kaMock

	var id int64
	id = 1
	_, err := is.PartialUpdate(context.TODO(), request.IcecreamUpdateRequest{Id: *request.NewCusInt64(&id)})

	assert.NotNil(t, err)

}

func TestPartialUpdate_IdNotPresent(t *testing.T) {

	config.InitializeTestConfig()

	builderMock := new(ibMock.IcecreamBuilderInterface)
	kaMock := new(kMock.SendMessageToKafkaHandler)
	daMock := new(dasMock.IcecreamDataAccessor)

	is := GetNewIcecreamService(config.GetDBConnection(), daMock)
	is.ib = builderMock
	is.kafkaHandler = kaMock

	_, err := is.PartialUpdate(context.TODO(), request.IcecreamUpdateRequest{})

	assert.NotNil(t, err)

}
