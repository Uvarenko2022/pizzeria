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
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"github.com/go-playground/locales/en"
	"github.com/golang/mock/gomock"
)

func TestAddFood(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockIFoodUC, food *entity.Food)

	testTable := []struct {
		name                string
		inputBody           string
		inputFood           *entity.Food
		mockBehavior        mockBehavior
		expectedStatusCode  int
		exectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"type":1, "name":"foodaa", "cost":40}`,
			inputFood: &entity.Food{
				Type: 1,
				Name: "foodaa",
				Cost: 40,
			},
			mockBehavior: func(s *mock_usecase.MockIFoodUC, food *entity.Food) {
				s.EXPECT().AddFood(food).Return(nil)
			},
			expectedStatusCode:  200,
			exectedResponseBody: "",
		},
		{
			name:      "Type Error",
			inputBody: `{"type":-1, "name":"foooda", "cost":40}`,
			inputFood: &entity.Food{
				Type: -1,
				Name: "foooda",
				Cost: 40,
			},
			mockBehavior:        func(s *mock_usecase.MockIFoodUC, food *entity.Food) {},
			expectedStatusCode:  400,
			exectedResponseBody: "bad type, type should be 0 or 1\n",
		},
		{
			name:      "Name Erorr",
			inputBody: `{"type":1, "name":"fo", "cost":40}`,
			inputFood: &entity.Food{
				Type: 1,
				Name: "fo",
				Cost: 40,
			},
			mockBehavior:        func(s *mock_usecase.MockIFoodUC, food *entity.Food) {},
			expectedStatusCode:  400,
			exectedResponseBody: "bad name, name should be more than 5 characters long and less than 30\n",
		},
		{
			name:      "Cost Error",
			inputBody: `{"type":1, "name":"foooda", "cost":10}`,
			inputFood: &entity.Food{
				Type: 1,
				Name: "foooda",
				Cost: 10,
			},
			mockBehavior:        func(s *mock_usecase.MockIFoodUC, food *entity.Food) {},
			expectedStatusCode:  400,
			exectedResponseBody: "bad cost, cost should be more than 15\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			fduc := mock_usecase.NewMockIFoodUC(c)
			testCase.mockBehavior(fduc, testCase.inputFood)

			usecases := &usecase.PiizzaUseCase{
				IFoodUC: fduc,
			}

			translator := en.New()
			uni := ut.New(translator, translator)
			trans, found := uni.GetTranslator("en")

			if !found {
				log.Fatal("translation not found")
			}

			v := validator.New()
			cv := validate.New(v, trans)
			validate.RegisterValidations(v)
			validate.RegisterMessages(v, trans)

			handler := New(usecases, cv)

			//Test Server
			r := chi.NewRouter()
			r.Post("/add-food", handler.AddFood)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/add-food", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.exectedResponseBody, w.Body.String())
		})
	}
}

func TestFoodGet(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockIFoodUC, ids []uint)

	testTable := []struct {
		name                 string
		inputBody            string
		inputIds             []uint
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponceBody string
	}{
		{
			name:      "OK",
			inputBody: `ids:[]`,
			inputIds:  []uint{},
			mockBehavior: func(s *mock_usecase.MockIFoodUC, ids []uint) {
				s.EXPECT().GetFood(ids).Return([]entity.Food{}, nil)
			},
			expectedStatusCode:   200,
			expectedResponceBody: "[]\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			fuc := mock_usecase.NewMockIFoodUC(c)
			testCase.mockBehavior(fuc, testCase.inputIds)

			usecases := &usecase.PiizzaUseCase{
				IFoodUC: fuc,
			}

			translator := en.New()
			uni := ut.New(translator, translator)
			trans, found := uni.GetTranslator("en")

			if !found {
				log.Fatal("translation not found")
			}

			v := validator.New()
			cv := validate.New(v, trans)
			validate.RegisterValidations(v)
			validate.RegisterMessages(v, trans)

			handler := New(usecases, cv)

			//Test Server
			r := chi.NewRouter()
			r.Get("/getfood", handler.GetFood)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/getfood", bytes.NewBufferString(testCase.inputBody))

			//Perform Request
			r.ServeHTTP(w, req)

			//Asserts
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponceBody, w.Body.String())
		})
	}
}

func TestUpdateFood(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockIFoodUC, food *entity.Food)

	testTable := []struct {
		name                string
		inputBody           string
		inputFood           *entity.Food
		mockBehavior        mockBehavior
		expectedStatusCode  int
		exectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"type":1, "name":"foodaa", "cost":40}`,
			inputFood: &entity.Food{
				Type: 1,
				Name: "foodaa",
				Cost: 40,
			},
			mockBehavior: func(s *mock_usecase.MockIFoodUC, food *entity.Food) {
				s.EXPECT().UpdateFood(food).Return(nil)
			},
			expectedStatusCode:  200,
			exectedResponseBody: "",
		},
		{
			name:      "Type Error",
			inputBody: `{"type":-1, "name":"foooda", "cost":40}`,
			inputFood: &entity.Food{
				Type: -1,
				Name: "foooda",
				Cost: 40,
			},
			mockBehavior:        func(s *mock_usecase.MockIFoodUC, food *entity.Food) {},
			expectedStatusCode:  400,
			exectedResponseBody: "bad type, type should be 0 or 1\n",
		},
		{
			name:      "Name Erorr",
			inputBody: `{"type":1, "name":"fo", "cost":40}`,
			inputFood: &entity.Food{
				Type: 1,
				Name: "fo",
				Cost: 40,
			},
			mockBehavior:        func(s *mock_usecase.MockIFoodUC, food *entity.Food) {},
			expectedStatusCode:  400,
			exectedResponseBody: "bad name, name should be more than 5 characters long and less than 30\n",
		},
		{
			name:      "Cost Error",
			inputBody: `{"type":1, "name":"foooda", "cost":10}`,
			inputFood: &entity.Food{
				Type: 1,
				Name: "foooda",
				Cost: 10,
			},
			mockBehavior:        func(s *mock_usecase.MockIFoodUC, food *entity.Food) {},
			expectedStatusCode:  400,
			exectedResponseBody: "bad cost, cost should be more than 15\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			fduc := mock_usecase.NewMockIFoodUC(c)
			testCase.mockBehavior(fduc, testCase.inputFood)

			usecases := &usecase.PiizzaUseCase{
				IFoodUC: fduc,
			}

			translator := en.New()
			uni := ut.New(translator, translator)
			trans, found := uni.GetTranslator("en")

			if !found {
				log.Fatal("translation not found")
			}

			v := validator.New()
			cv := validate.New(v, trans)
			validate.RegisterValidations(v)
			validate.RegisterMessages(v, trans)

			handler := New(usecases, cv)

			//Test Server
			r := chi.NewRouter()
			r.Post("/add-food", handler.UpdateFood)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/add-food", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.exectedResponseBody, w.Body.String())
		})
	}
}
