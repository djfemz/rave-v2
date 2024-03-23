package repositories

import (
	"errors"
	"github.com/djfemz/rave/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
)

type Repository[T, U any] interface {
	crudRepository[T, U]
}

type crudRepository[T, U any] interface {
	Save(t *T) *T
	FindById(id U) *T
	FindAll() []*T
	FindAllBy(pageable Pageable) []*T
	DeleteById(id U)
}

type repositoryImpl[T, U any] struct {
}

var db *gorm.DB

func (r *repositoryImpl[T, U]) Save(t *T) *T {
	db = connect()
	db = db.Save(t)
	var id, _ = GetId(*t)
	db.First(t, id)
	return t
}

func (r *repositoryImpl[T, U]) FindById(id U) *T {
	db = connect()
	var t = new(T)
	db.First(t, id)
	return t
}

func (r *repositoryImpl[T, U]) FindAll() []*T {
	db = connect()
	var organizations []*T
	db.Find(&organizations)
	return organizations
}

func (r *repositoryImpl[T, U]) FindAllBy(pageable Pageable) []*T {
	db = connect()
	page := getPage[T](db, pageable)
	return page.GetElements()
}

func (r *repositoryImpl[T, U]) DeleteById(id U) {
	db = connect()
	db.Delete(new(T), id)
}

func connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=rave port=5432 sslmode=disable TimeZone=Africa/Lagos"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	_ = addEntities(models.Entities, db)
	return db
}

func addEntities(m map[string]any, db *gorm.DB) error {
	for _, v := range m {
		err := db.AutoMigrate(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func GetId(T any) (any, error) {
	obj := reflect.ValueOf(T)
	numberOfFields := obj.NumField()
	for index := ZERO; index < numberOfFields; index++ {
		tag := obj.Type().Field(index).Tag
		isPrimaryKeyField := strings.Contains(string(tag), "id")
		if isPrimaryKeyField {
			id, err, done := getPrimaryKey(obj, index)
			if done {
				return id, err
			}
		}
	}
	return nil, errors.New("could not retrieve id")
}

func getPrimaryKey(obj reflect.Value, index int) (any, error, bool) {
	idField := obj.Field(index)
	if idField.CanConvert(reflect.TypeOf(uint64(5))) {
		v := idField.Uint()
		return v, nil, true
	} else if idField.CanConvert(reflect.TypeOf(5)) {
		v := idField.Int()
		return v, nil, true
	} else if idField.CanConvert(reflect.TypeOf("")) {
		return idField.String(), nil, true
	}
	return nil, nil, false
}
