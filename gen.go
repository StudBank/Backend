package main

import (
	"gitea.repetitra.ru/StudBank/Backend/models"
	"gorm.io/gen"
)

func gendb() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./db",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.ApplyBasic(models.User{})
	g.ApplyInterface(func(models.UserQuerier) {}, models.User{})
	g.Execute()
}
