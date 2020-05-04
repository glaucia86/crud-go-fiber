package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //lib postgres para fazer a conexão com a base de dados
	"github.com/joho/godotenv"
)

// Aqui estamos criando uma variável para a base de dados:
var db *gorm.DB

func InitDatabase() {

	// Aqui estamos carregando o arquivo '.env':
	envFile := godotenv.Load()
	// Se o arquivo '.env' for diferente de nulo, então...
	if envFile != nil {
		fmt.Print(envFile)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	// Aqui estamos criando a nossa 'connectionString' com o PostgreSQL
	dbConnectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println("Conexão realizada com sucesso!")

	conn, err := gorm.Open("postgres", dbConnectionString)
	if err != nil {
		fmt.Println("Erro ao realizar a conexão com a Base de Dados..: ", err)
	}

	db = conn

	defer conn.Close()
}

//ReturnDatabase ... Aqui vamos retornar toda a conexão com a base de dados!
func ReturnDatabase() *gorm.DB {
	return db
}
