package db

import (
	"CarFix/internal/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
	Ctx  context.Context
}

func (db *Database) GetAllClients(ctx context.Context) (clients []*models.Client, err error) {
	err = pgxscan.Select(ctx, db.Conn, &clients,
		`SELECT id, name, last_name, phone_number FROM clients`)

	return
}

func (db *Database) GetClientById(ctx context.Context, id int) (client []*models.Client, err error) {
	err = pgxscan.Select(ctx, db.Conn, &client,
		`SELECT id, name, last_name, phone_number FROM clients WHERE id=$1`, id)

	return
}

func (db *Database) CreateNewClient(ctx context.Context, client *models.Client) (err error) {
	_, err = db.Conn.Exec(ctx,
		`INSERT INTO clients (name, last_name, phone_number) VALUES ($1, $2, $3)`,
		client.Name, client.LastName, client.PhoneNumber)

	return
}

func (db *Database) EditClient(ctx context.Context, client *models.Client) (err error) {
	_, err = db.Conn.Exec(ctx,
		"UPDATE clients SET name=$1, last_name=$2, phone_number=$3 WHERE id=$4",
		client.Name, client.LastName, client.PhoneNumber, client.ID)

	return
}

func (db *Database) DeleteClient(ctx context.Context, id int) (err error) {
	_, err = db.Conn.Exec(ctx, "DELETE FROM clients WHERE id=$1", id)

	return
}
