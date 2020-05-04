// Info sobre como realizar connection com Postgresql + GORM
// https://gist.github.com/adigunhammedolalekan/d65145512cb1de55e40d74a37fe34f5a

package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// Aqui estamos criando uma variável para a base de dados:
var database *gorm.DB

func initDatabase() {

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

	conn, err := gorm.Open("postgress", dbConnectionString)
	if err != nil {
		fmt.Println("Erro ao realizar a conexão com a Base de Dados..: ", err)
	}

	database = conn

	defer conn.Close()
}
