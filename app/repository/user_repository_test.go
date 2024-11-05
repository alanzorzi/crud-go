package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alanzorzi/crud-go/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	// Cria uma conexão de banco de dados mockada
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Instancia o repositório com o banco de dados mockado
	repo := &userRepository{db: db}

	// Cria um usuário de teste
	user := &model.User{
		ID:       "1",
		Name:     "John Doe",
		Age:      30,
		Email:    "johndoe@example.com",
		Password: "password123",
	}

	// Define o comportamento esperado para o banco de dados mockado
	mock.ExpectExec("INSERT INTO user_data").
		WithArgs(user.ID, user.Name, user.Age, user.Email, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Sucesso na inserção

	// Executa o método a ser testado
	err = repo.CreateUser(user)

	// Verifica se não houve erro
	assert.NoError(t, err)

	// Verifica se todas as expectativas foram atendidas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateUser_Error(t *testing.T) {
	// Cria uma conexão de banco de dados mockada
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Instancia o repositório com o banco de dados mockado
	repo := &userRepository{db: db}

	// Cria um usuário de teste
	user := &model.User{
		ID:       "2",
		Name:     "Jane Doe",
		Age:      25,
		Email:    "janedoe@example.com",
		Password: "password456",
	}

	// Define o comportamento esperado para o banco de dados mockado (erro na inserção)
	mock.ExpectExec("INSERT INTO user_data").
		WithArgs(user.ID, user.Name, user.Age, user.Email, user.Password).
		WillReturnError(assert.AnError) // Erro durante a inserção

	// Executa o método a ser testado
	err = repo.CreateUser(user)

	// Verifica se houve erro conforme esperado
	assert.Error(t, err)

	// Verifica se todas as expectativas foram atendidas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
