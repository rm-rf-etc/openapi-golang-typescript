package main

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml openapi/api.yaml

import (
	"go_oapi/api"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type UserStore struct {
	Lock sync.Mutex
}

func (u *UserStore) GetCurrentUser(c *gin.Context) {
	u.Lock.Lock()
	defer u.Lock.Unlock()

	id, _ := uuid.Parse("12bd422c-70b5-11ef-8108-962f3e3a692f")
	email := types.Email("jim@email.com")

	c.IndentedJSON(http.StatusOK, api.User{
		Id:    &id,
		Email: &email,
	})
}

func main() {
	r := gin.Default()
	api.RegisterHandlers(r, &UserStore{})
	r.Run("localhost:8080")
}
