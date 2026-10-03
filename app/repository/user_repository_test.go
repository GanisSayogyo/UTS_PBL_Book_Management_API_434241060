package repository

import (
	"context"
	"testing"
	"time"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/model"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/config"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/database"
	"github.com/joho/godotenv"
)

func TestUserRepository(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("failed to load .env: %v", err)
	}

	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	username := "testuser_" + time.Now().Format("150405")

	user := &model.User{
		Username: username,
		Email:    username + "@example.com",
		Password: "hashed-password",
		Role:     "user",
	}

	ctx := context.Background()

	err = repo.Create(ctx, user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	if user.ID == 0 {
		t.Fatal("expected user ID to be generated")
	}

	foundUser, err := repo.FindByUsername(ctx, username)
	if err != nil {
		t.Fatalf("failed to find user: %v", err)
	}

	if foundUser.Username != username {
		t.Fatalf("expected username %s, got %s", username, foundUser.Username)
	}

	if foundUser.Email != user.Email {
		t.Fatalf("expected email %s, got %s", user.Email, foundUser.Email)
	}

	_, err = db.Exec(
		ctx,
		"DELETE FROM users WHERE id = $1",
		user.ID,
	)
	if err != nil {
		t.Fatalf("failed to clean up test user: %v", err)
	}
}
