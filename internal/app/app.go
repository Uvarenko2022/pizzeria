package app

import (
	"Uvarenko2022/restaurant/database"
	"Uvarenko2022/restaurant/internal/controller/http"
	"Uvarenko2022/restaurant/internal/repo"
	"Uvarenko2022/restaurant/internal/usecase"
	"Uvarenko2022/restaurant/internal/validate"
	"context"
	"fmt"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func Run() {
	godotenv.Load()

	ctx := context.Background()

	//databases
	db := database.NewPostgre()
	cache := database.NewRedis(ctx)

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

	repos := repo.NewRepository(db, cache)

	usecase := usecase.New(repos)

	food, err := usecase.IFoodUC.GetFood([]uint{})

	if err != nil {
		fmt.Println("FUCKING APP")
		fmt.Println(err)
	}

	database.SetUpHookCache(db, ctx, usecase.IFoodUC.GetFood, usecase.ICacheUC.UpdateCache)
	usecase.UpdateCache(context.Background(), food)

	rout := http.New(usecase, cv)
	http.RegisterRouts(rout)
}
