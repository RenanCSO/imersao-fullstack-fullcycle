package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/RenanCSO/imersao-fullstack-fullcycle/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {
	name := "Renan Campos Silverio de Oliveira"
	email := "renan.c.s.o@gmail.com"
	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.name, name)
	require.Equal(t, user.email, email)
}