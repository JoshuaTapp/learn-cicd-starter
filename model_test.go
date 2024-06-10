package main

import (
	"testing"
	"time"

	"github.com/bootdotdev/learn-cicd-starter/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseUserToUser(t *testing.T) {
	dbUser := database.User{
		ID:        "user123",
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		Name:      "John Doe",
		ApiKey:    "apikey123",
	}

	user, err := databaseUserToUser(dbUser)
	assert.NoError(t, err)
	assert.Equal(t, dbUser.ID, user.ID)
	assert.Equal(t, dbUser.Name, user.Name)
	assert.Equal(t, dbUser.ApiKey, user.ApiKey)

	// Check time parsing
	expectedCreatedAt, _ := time.Parse(time.RFC3339, dbUser.CreatedAt)
	expectedUpdatedAt, _ := time.Parse(time.RFC3339, dbUser.UpdatedAt)
	assert.Equal(t, expectedCreatedAt, user.CreatedAt)
	assert.Equal(t, expectedUpdatedAt, user.UpdatedAt)
}

func TestDatabaseUserToUser_Error(t *testing.T) {
	dbUser := database.User{
		ID:        "user123",
		CreatedAt: "invalid-time",
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		Name:      "John Doe",
		ApiKey:    "apikey123",
	}

	user, err := databaseUserToUser(dbUser)
	assert.Error(t, err)
	assert.Equal(t, User{}, user)
}

func TestDatabaseNoteToNote(t *testing.T) {
	dbNote := database.Note{
		ID:        "note123",
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		Note:      "Sample Note",
		UserID:    "user123",
	}

	note, err := databaseNoteToNote(dbNote)
	assert.NoError(t, err)
	assert.Equal(t, dbNote.ID, note.ID)
	assert.Equal(t, dbNote.Note, note.Note)
	assert.Equal(t, dbNote.UserID, note.UserID)

	// Check time parsing
	expectedCreatedAt, _ := time.Parse(time.RFC3339, dbNote.CreatedAt)
	expectedUpdatedAt, _ := time.Parse(time.RFC3339, dbNote.UpdatedAt)
	assert.Equal(t, expectedCreatedAt, note.CreatedAt)
	assert.Equal(t, expectedUpdatedAt, note.UpdatedAt)
}

func TestDatabaseNoteToNote_Error(t *testing.T) {
	dbNote := database.Note{
		ID:        "note123",
		CreatedAt: "invalid-time",
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		Note:      "Sample Note",
		UserID:    "user123",
	}

	note, err := databaseNoteToNote(dbNote)
	assert.Error(t, err)
	assert.Equal(t, Note{}, note)
}

func TestDatabasePostsToPosts(t *testing.T) {
	dbNotes := []database.Note{
		{
			ID:        "note123",
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
			Note:      "Sample Note 1",
			UserID:    "user123",
		},
		{
			ID:        "note124",
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
			Note:      "Sample Note 2",
			UserID:    "user123",
		},
	}

	notes, err := databasePostsToPosts(dbNotes)
	assert.NoError(t, err)
	assert.Len(t, notes, len(dbNotes))

	for i := range dbNotes {
		assert.Equal(t, dbNotes[i].ID, notes[i].ID)
		assert.Equal(t, dbNotes[i].Note, notes[i].Note)
		assert.Equal(t, dbNotes[i].UserID, notes[i].UserID)

		expectedCreatedAt, _ := time.Parse(time.RFC3339, dbNotes[i].CreatedAt)
		expectedUpdatedAt, _ := time.Parse(time.RFC3339, dbNotes[i].UpdatedAt)
		assert.Equal(t, expectedCreatedAt, notes[i].CreatedAt)
		assert.Equal(t, expectedUpdatedAt, notes[i].UpdatedAt)
	}
}

func TestDatabasePostsToPosts_Error(t *testing.T) {
	dbNotes := []database.Note{
		{
			ID:        "note123",
			CreatedAt: "invalid-time",
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
			Note:      "Sample Note 1",
			UserID:    "user123",
		},
		{
			ID:        "note124",
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
			Note:      "Sample Note 2",
			UserID:    "user123",
		},
	}

	notes, err := databasePostsToPosts(dbNotes)
	assert.Error(t, err)
	assert.Nil(t, notes)
}
