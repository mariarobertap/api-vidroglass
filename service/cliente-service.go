package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
	_ "github.com/mattn/go-sqlite3"
)

type clienteService struct {
	Client     models.Cliente
	ClientList []models.Cliente
}

func NewClienteService() interfaces.ClienteService {
	return &clienteService{}
}

func (c *clienteService) Save(cliente models.Cliente) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("INSERT INTO cliente(id_endereco, nome, cpf, telefone) values(?,?,?, ?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(cliente.Id_endereco, cliente.Nome, cliente.Cpf, cliente.Telefone)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println(id)

	db.Close()

	return int(id), nil

}

func (c *clienteService) FindAll() ([]models.Cliente, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	rows, err := db.Query("SELECT * FROM cliente")

	if err != nil {
		fmt.Println(err)
		return c.ClientList, err
	}

	c.ClientList = nil

	for rows.Next() {
		err = rows.Scan(&c.Client.Id_cliente,
			&c.Client.Id_endereco,
			&c.Client.Nome,
			&c.Client.Cpf,
			&c.Client.Telefone)
		if err != nil {
			fmt.Println(err)
			return c.ClientList, err
		}
		c.ClientList = append(c.ClientList, c.Client)
	}

	rows.Close()
	db.Close()
	return c.ClientList, nil
}

func (c *clienteService) GetClientById(id_cliente int) (models.Cliente, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	row := db.QueryRow("SELECT * FROM cliente WHERE id_cliente = ?", id_cliente)

	err = row.Scan(&c.Client.Id_cliente,
		&c.Client.Id_endereco,
		&c.Client.Nome,
		&c.Client.Cpf,
		&c.Client.Telefone)

	if err != nil {
		fmt.Println(err)
		return c.Client, err
	}

	return c.Client, nil
}

func (c *clienteService) UpdateClientById(cliente models.Cliente) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("UPDATE cliente SET id_endereco = ?, nome = ?, cpf = ?, telefone = ? WHERE id_cliente = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(cliente.Id_endereco, cliente.Nome, cliente.Cpf, cliente.Telefone, cliente.Id_cliente)

	fmt.Println(res)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
