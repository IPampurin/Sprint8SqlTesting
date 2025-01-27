package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	clientID := 1
	client, err := selectClient(db, clientID)

	require.NoError(t, err)
	assert.Equal(t, clientID, client.ID)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Login)
	assert.NotEmpty(t, client.Birthday)
	assert.NotEmpty(t, client.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	clientID := -1
	client, err := selectClient(db, clientID)

	require.ErrorIs(t, err, sql.ErrNoRows)

	assert.NotEmpty(t, client.ID)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Login)
	assert.NotEmpty(t, client.Birthday)
	assert.NotEmpty(t, client.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)

	require.NotEmpty(t, id)
	require.NoError(t, err)

	cl.ID = id

	client, err := selectClient(db, cl.ID)

	require.NoError(t, err)
	assert.Equal(t, cl, client)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)

	require.NotEmpty(t, id)
	require.NoError(t, err)

	cl.ID = id

	_, err = selectClient(db, cl.ID)
	require.NoError(t, err)

	err = deleteClient(db, cl.ID)
	require.NoError(t, err)

	_, err = selectClient(db, cl.ID)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}
