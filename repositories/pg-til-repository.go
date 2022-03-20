package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/cbodonnell/til-api/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PGTilRepository struct {
	db *pgxpool.Pool
}

func NewPSQLTilRepository(connStr string) TilRepository {
	return &PGTilRepository{
		db: connectDb(connStr),
	}
}

func connectDb(connStr string) *pgxpool.Pool {
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	log.Printf("Connected to %s as %s\n", db.Config().ConnConfig.Database, db.Config().ConnConfig.User)
	return db
}

func (r *PGTilRepository) Close() {
	r.db.Close()
}

func (r *PGTilRepository) GetAllByUserID(user_uuid string) ([]models.Til, error) {
	sql := `SELECT id, text FROM tils
	WHERE user_uuid = $1`

	rows, err := r.db.Query(context.Background(), sql, user_uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tils := make([]models.Til, 0)
	for rows.Next() {
		var til models.Til
		err = rows.Scan(
			&til.ID,
			&til.Text,
		)
		if err != nil {
			return tils, err
		}
		tils = append(tils, til)
	}
	err = rows.Err()
	if err != nil {
		return tils, err
	}
	return tils, nil
}

func (r *PGTilRepository) Create(user_uuid string, til models.Til) (models.Til, error) {
	sql := `INSERT INTO tils (text, user_uuid)
	VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRow(context.Background(), sql, til.Text, user_uuid).Scan(&til.ID)
	if err != nil {
		return til, err
	}
	return til, nil
}
