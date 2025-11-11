package store


import (
"gorm.io/driver/postgres"
"gorm.io/gorm"


"github.com/ShadowNos1/qa-api/internal/model"
)


func NewPostgres(dsn string) (*gorm.DB, error) {
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
return nil, err
}
// Не использовать AutoMigrate как единственный источник правды.
// Здесь для dev удобства можно включить, но миграции должны быть в migrations/.
_ = db.AutoMigrate(&model.Question{}, &model.Answer{})
return db, nil
}