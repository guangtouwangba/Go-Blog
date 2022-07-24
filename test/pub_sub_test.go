package test

import (
	"Go-Blog/internal/domain/entity"
	"log"
	"testing"
)

func Test_should_test_pub_sub_successfully(t *testing.T) {
	user := entity.User{}
	comments := &entity.Comment{}
	comments.Register(&user)
	comments.Notify()
	log.Printf("Test_should_test_pub_sub_successfully")
}
