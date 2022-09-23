package http

import (
	"Uvarenko2022/restaurant/internal/entity"
	"Uvarenko2022/restaurant/internal/usecase"
	mock_usecase "Uvarenko2022/restaurant/internal/usecase/mock"
	"Uvarenko2022/restaurant/internal/validate"
	"bytes"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockIOrderUC, order *entity.Order, food []entity.Food)

	testTable := []struct {
		name                 string
		inputBody            string
		inputOrder           *entity.Order
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponceBody string
	}{
		{
			name: "OK",
			inputOrder: &entity.Order{
				FoodIds: []uint{1, 2},
				State:   1,
			},
			inputBody: `{"foodids":[1,2], "state":1}`,
			mockBehavior: func(s *mock_usecase.MockIOrderUC, order *entity.Order, food []entity.Food) {
				s.EXPECT().CreateOrder(order, food).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponceBody: "[]\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			ouc := mock_usecase.NewMockIOrderUC(c)
			fuc := mock_usecase.NewMockIFoodUC(c)
			testCase.mockBehavior(ouc, testCase.inputOrder, []entity.Food{
				{
					Type: 1,
					Name: "niggaaaaa",
					Cost: 40,
				},
				{
					Type: 1,
					Name: "niggaaaaa",
					Cost: 40,
				},
				{
					Type: 1,
					Name: "niggaaaaa",
					Cost: 40,
				},
				{
					Type: 1,
					Name: "niggaaaaa",
					Cost: 40,
				},
			})
			usecases := &usecase.PiizzaUseCase{
				IOrderUC: ouc,
				IFoodUC:  fuc,
			}

			//translator
			translator := en.New()
			uni := ut.New(translator, translator)
			trans, found := uni.GetTranslator("en")

			if !found {
				log.Fatal("translation not found")
			}

			//validations
			v := validator.New()
			cv := validate.New(v, trans)
			validate.RegisterValidations(v)
			validate.RegisterMessages(v, trans)

			handler := New(usecases, cv)

			r := chi.NewRouter()
			r.Post("/create-order", handler.CreateOrder)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create-order", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponceBody, w.Body.String())
		})
	}
}
