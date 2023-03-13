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
