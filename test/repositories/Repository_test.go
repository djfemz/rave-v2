package test

import (
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var repository repositories.Repository[models.Organizer, uint64] = &repositories.RepositoryImpl[models.Organizer, uint64]{}

func TestRepositoryImpl_Save(t *testing.T) {
	var savedOrg = repository.Save(&models.Organizer{
		Name:      "John",
		CreatedAt: time.Now(),
	})
	log.Println(savedOrg)
	assert.NotNil(t, savedOrg)
}

func TestRepositoryImpl_FindById(t *testing.T) {
	foundOrg := repository.FindById(3)
	log.Println(foundOrg)
	assert.NotNil(t, foundOrg)
}

func TestRepositoryImpl_FindAll(t *testing.T) {

}

func TestRepositoryImpl_DeleteById(t *testing.T) {

}

func TestGetId(t *testing.T) {
	var event = models.Event{
		ID:   24,
		Name: "John",
		Date: time.Now(),
	}
	var id, _ = repositories.GetId(event)
	assert.Equal(t, uint64(24), id)
}
