package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Credenciales de la base de datos  ver en el readme.md
const driver = "mysql"
const dataSourceName = "user:pass@tcp(localhost:3306)/db"

// Variable global para la coneccion a la base de datos del tipo puntero *sql.DB
var db *sql.DB

// Realiza la coneccion
func Connect() {
	//Crear coneccion a la base de datos con libreria de sql Open()
	coneccion, err := sql.Open(driver, dataSourceName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado")
	db = coneccion

}

// Cierra la coneccion
func Close() {
	db.Close()
}

// Verifica la coneccion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Verificar si existe tabla
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return rows.Next() //retorna true si existe la tabla
}

// Crear tabla usuarios

func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return result, err
}

// Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return rows, err
}

// Reiniciar el registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	Exec(sql)
}
