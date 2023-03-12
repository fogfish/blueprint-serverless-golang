package http_test

import (
	"errors"
	"fmt"
	"testing"

	core "github.com/fogfish/blueprint-serverless-golang"
	"github.com/fogfish/blueprint-serverless-golang/http"
	"github.com/fogfish/blueprint-serverless-golang/http/api"
	"github.com/fogfish/blueprint-serverless-golang/mock"
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
	request := µmock.Input(
		µmock.Method("GET"),
		µmock.URL("/petshop/pets"),
		µmock.Header("Accept", "application/json"),
	)

	for _, tt := range []struct {
		Input *µ.Context
		Setup func(*mock.MockPetFetcher)
		Yield error
	}{
		{
			request,
			func(mpf *mock.MockPetFetcher) {
				mpf.EXPECT().LookupPetsAfterKey(
					gomock.Any(),
					gomock.Eq(""),
					gomock.Any(),
				).Return(mock.Pets[0:1], nil)
			},
			ø.Status.OK(
				ø.ContentType.ApplicationJSON,
				ø.Send(api.NewPets(2, mock.Pets[0:1])),
			),
		},
		{
			request,
			func(mpf *mock.MockPetFetcher) {
				mpf.EXPECT().LookupPetsAfterKey(
					gomock.Any(),
					gomock.Eq(""),
					gomock.Any(),
				).Return(nil, errors.New("fault"))
			},
			ø.Status.InternalServerError(
				ø.Error(errors.New("fault")),
			),
		},
		{
			µmock.Input(
				µmock.Method("GET"),
				µmock.URL("/petshop/pets?cursor=A02"),
				µmock.Header("Accept", "application/json"),
			),
			func(mpf *mock.MockPetFetcher) {
				mpf.EXPECT().LookupPetsAfterKey(
					gomock.Any(),
					gomock.Eq("A02"),
					gomock.Any(),
				).Return(mock.Pets[0:1], nil)
			},
			ø.Status.OK(
				ø.ContentType.ApplicationJSON,
				ø.Send(api.NewPets(2, mock.Pets[0:1])),
			),
		},
	} {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mpf := mock.NewMockPetFetcher(ctrl)
		tt.Setup(mpf)

		api := http.NewPetShopAPI(mpf, nil)
		service := µmock.Endpoint(api.List())

		it.Then(t).Should(
			it.Equiv(service(tt.Input), tt.Yield),
		)
	}
}

func TestPetShopLookup(t *testing.T) {
	request := µmock.Input(
		µmock.Method("GET"),
		µmock.URL("/petshop/pets/A01"),
		µmock.Header("Accept", "application/json"),
	)

	for _, tt := range []struct {
		Input *µ.Context
		Setup func(*mock.MockPetFetcher)
		Yield error
	}{
		{
			request,
			func(mpf *mock.MockPetFetcher) {
				mpf.EXPECT().LookupPet(
					gomock.Any(),
					gomock.Eq("A01"),
				).Return(mock.Pets[0], nil)
			},
			ø.Status.OK(
				ø.ContentType.ApplicationJSON,
				ø.Send(api.NewPet(mock.Pets[0])),
			),
		},
		{
			request,
			func(mpf *mock.MockPetFetcher) {
				mpf.EXPECT().LookupPet(
					gomock.Any(),
					gomock.Eq("A01"),
				).Return(core.Pet{}, errors.New("fault"))
			},
			ø.Status.InternalServerError(
				ø.Error(errors.New("fault")),
			),
		},
	} {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mpf := mock.NewMockPetFetcher(ctrl)
		tt.Setup(mpf)

		api := http.NewPetShopAPI(mpf, nil)
		service := µmock.Endpoint(api.Lookup())

		it.Then(t).Should(
			it.Equiv(service(tt.Input), tt.Yield),
		)
	}
}

func TestPetShopCreate(t *testing.T) {
	for _, tt := range []struct {
		Input *µ.Context
		Setup func(*mock.MockPetCreator)
		Yield error
	}{
		{
			µmock.Input(
				µmock.Method("POST"),
				µmock.URL("/petshop/pets"),
				µmock.Header("Content-Type", "application/json"),
				µmock.Header("Authorization", "Basic cGV0c3RvcmU6b3duZXIK"),
				µmock.JSON(mock.Pets[0]),
			),
			func(mpf *mock.MockPetCreator) {
				mpf.EXPECT().CreatePet(
					gomock.Any(),
					gomock.Eq(mock.Pets[0]),
				).Return(nil)
			},
			ø.Status.Created(),
		},
		{
			µmock.Input(
				µmock.Method("POST"),
				µmock.URL("/petshop/pets"),
				µmock.Header("Content-Type", "application/json"),
				µmock.Header("Authorization", "Basic cGV0c3RvcmU6b3duZXI_"),
				µmock.JSON(mock.Pets[0]),
			),
			func(mpf *mock.MockPetCreator) {},
			ø.Status.Unauthorized(
				ø.Error(µ.ErrNoMatch),
			),
		},
		{
			µmock.Input(
				µmock.Method("POST"),
				µmock.URL("/petshop/pets"),
				µmock.Header("Content-Type", "application/json"),
				µmock.JSON(mock.Pets[0]),
			),
			func(mpf *mock.MockPetCreator) {},
			ø.Status.Unauthorized(
				ø.Error(fmt.Errorf("unauthorized /petshop/pets")),
			),
		},
	} {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mpc := mock.NewMockPetCreator(ctrl)
		tt.Setup(mpc)

		api := http.NewPetShopAPI(nil, mpc)
		service := µmock.Endpoint(api.Create())

		it.Then(t).Should(
			it.Equiv(service(tt.Input), tt.Yield),
		)
	}
}
