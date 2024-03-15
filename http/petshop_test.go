package http_test

import (
	"errors"
	"testing"

	"github.com/fogfish/blueprint-serverless-golang/http"
	"github.com/fogfish/blueprint-serverless-golang/http/api"
	"github.com/fogfish/blueprint-serverless-golang/internal/mock"
	µ "github.com/fogfish/gouldian/v2"
	µmock "github.com/fogfish/gouldian/v2/mock"
	ø "github.com/fogfish/gouldian/v2/output"
	"github.com/fogfish/guid/v2"
	"github.com/fogfish/it/v2"
	"github.com/golang/mock/gomock"
)

func init() {
	µ.Sequence = guid.NewClockMock()
}

func TestPetShopList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mpf := mock.NewMockPetFetcher(ctrl)
	mpf.EXPECT().LookupPetsAfterKey(
		gomock.Any(),
		gomock.Eq(""),
		gomock.Any(),
	).Return(mock.Pets[0:1], nil)

	service := http.NewPetShopAPI(mpf, nil)
	httpd := µmock.Endpoint(service.List())
	yield := httpd(µmock.Input(
		µmock.Method("GET"),
		µmock.URL("/petshop/pets"),
		µmock.Header("Accept", "application/json"),
	))

	it.Then(t).Should(
		it.Equiv(yield,
			ø.Status.OK(
				ø.ContentType.ApplicationJSON,
				ø.Send(api.NewPets(2, mock.Pets[0:1])),
			),
		),
	)
}

func TestPetShopListWithCursor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mpf := mock.NewMockPetFetcher(ctrl)
	mpf.EXPECT().LookupPetsAfterKey(
		gomock.Any(),
		gomock.Eq("A02"),
		gomock.Any(),
	).Return(mock.Pets[0:1], nil)

	service := http.NewPetShopAPI(mpf, nil)
	httpd := µmock.Endpoint(service.List())
	yield := httpd(µmock.Input(
		µmock.Method("GET"),
		µmock.URL("/petshop/pets?cursor=A02"),
		µmock.Header("Accept", "application/json"),
	))

	it.Then(t).Should(
		it.Equiv(yield,
			ø.Status.OK(
				ø.ContentType.ApplicationJSON,
				ø.Send(api.NewPets(2, mock.Pets[0:1])),
			),
		),
	)
}

func TestPetShopListFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mpf := mock.NewMockPetFetcher(ctrl)
	mpf.EXPECT().LookupPetsAfterKey(
		gomock.Any(),
		gomock.Eq(""),
		gomock.Any(),
	).Return(nil, errors.New("fault"))

	service := http.NewPetShopAPI(mpf, nil)
	httpd := µmock.Endpoint(service.List())
	yield := httpd(µmock.Input(
		µmock.Method("GET"),
		µmock.URL("/petshop/pets"),
		µmock.Header("Accept", "application/json"),
	))

	it.Then(t).Should(
		it.Equiv(yield,
			ø.Status.InternalServerError(
				ø.Error(errors.New("fault")),
			),
		),
	)
}

func TestPetShopLookup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mpf := mock.NewMockPetFetcher(ctrl)
	mpf.EXPECT().LookupPet(
		gomock.Any(),
		gomock.Eq("A01"),
	).Return(mock.Pets[0], nil)

	service := http.NewPetShopAPI(mpf, nil)
	httpd := µmock.Endpoint(service.Lookup())
	yield := httpd(µmock.Input(
		µmock.Method("GET"),
		µmock.URL("/petshop/pets/A01"),
		µmock.Header("Accept", "application/json"),
	))

	it.Then(t).Should(
		it.Equiv(yield,
			ø.Status.OK(
				ø.ContentType.ApplicationJSON,
				ø.Send(api.NewPet(mock.Pets[0])),
			),
		),
	)
}

func TestPetShopCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mpc := mock.NewMockPetCreator(ctrl)
	mpc.EXPECT().CreatePet(
		gomock.Any(),
		gomock.Eq(mock.Pets[0]),
	).Return(nil)

	service := http.NewPetShopAPI(nil, mpc)
	httpd := µmock.Endpoint(service.Create())
	yield := httpd(µmock.Input(
		µmock.Method("POST"),
		µmock.URL("/petshop/pets"),
		µmock.Header("Accept", "application/json"),
		µmock.Header("X-Secret-Code", "cGV0c3RvcmU6b3duZXIK"),
		µmock.JSON(mock.Pets[0]),
	))

	it.Then(t).Should(
		it.Equiv(yield,
			ø.Status.Created(),
		),
	)
}
