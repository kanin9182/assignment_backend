package helper

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateModels(db *gorm.DB) error {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internals/repositories/models",
		ModelPkgPath: "models",
	})
	g.UseDB(db)
	g.GenerateAllTable()
	g.Execute()
	return nil
}
