package profile

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/jackc/pgx/v5"
	"testing"
)

func TestGetProfile_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании sqlmock: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	mock.ExpectQuery("SELECT id, email, username, age, gender, avatar_url FROM users WHERE users.id = \\$1;").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "age", "gender", "avatar_url"}).
			AddRow(1, "test@example.com", "testuser", 25, "male", "http://example.com/avatar.jpg"))

	mock.ExpectQuery("SELECT .* FROM addresses WHERE profile_id = \\$1").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"street", "city", "country"}).
			AddRow("Main St", "Metropolis", "Wonderland"))

	pr := &ProfilesStore{
		db:  db,
		log: nil, // Замените на ваш логгер, если необходимо
	}

	profile, err := pr.GetProfile(ctx, 1)
	if err != nil {
		t.Errorf("ожидалась успешная операция, но получена ошибка: %v", err)
	}

	if profile.ID != 1 || profile.Email != "test@example.com" || profile.Username != "testuser" {
		t.Errorf("неожиданный профиль: %+v", profile)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все ожидания были выполнены: %v", err)
	}
}

func TestGetProfile_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании sqlmock: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	mock.ExpectQuery("SELECT id, email, username, age, gender, avatar_url FROM users WHERE users.id = \\$1;").
		WithArgs(2).
		WillReturnError(pgx.ErrNoRows)

	pr := &ProfilesStore{
		Db:  db,
		Log: nil,
	}

	_, err = pr.GetProfile(ctx, 2)
	if !errors.Is(err, errs.UserDoesNotExist) {
		t.Errorf("ожидалась ошибка UserDoesNotExist, но получена: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все ожидания были выполнены: %v", err)
	}
}

func TestGetProfile_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании sqlmock: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	mock.ExpectQuery("SELECT id, email, username, age, gender, avatar_url FROM users WHERE users.id = \\$1;").
		WithArgs(3).
		WillReturnError(errors.New("database error"))

	pr := &ProfilesStore{
		Db:  db,
		Log: nil,
	}

	_, err = pr.GetProfile(ctx, 3)
	if err == nil || err.Error() != "database error" {
		t.Errorf("ожидалась ошибка 'database error', но получена: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все ожидания были выполнены: %v", err)
	}
}
