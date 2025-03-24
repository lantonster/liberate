package main

import (
	"github.com/lantonster/liberate/internal/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/orm",
		Mode:    gen.WithDefaultQuery, // generate mode
	})

	g.ApplyBasic(
		model.User{},
	)

	g.Execute()
}
