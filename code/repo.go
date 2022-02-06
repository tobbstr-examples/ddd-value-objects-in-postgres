package code

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func (r *Repository) UpsertPerson(ctx context.Context, person PersonRecord) error {
	query := `
	INSERT INTO person (id, name, personal_info)
	VALUES('%s','%s','%s')
	ON CONFLICT (id) DO
	UPDATE SET name = '%s', personal_info = '%s';`

	bytes, err := json.Marshal(person.PersonalInfo)
	if err != nil {
		return err
	}
	personalInfoAsJSON := string(bytes)

	query = fmt.Sprintf(query, person.ID, person.Name, personalInfoAsJSON, person.Name, personalInfoAsJSON)

	if _, err := r.db.Exec(ctx, query); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetPersons(ctx context.Context) ([]PersonRecord, error) {
	query := "SELECT id, name, personal_info FROM person;"

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var personRecords []PersonRecord

	for rows.Next() {
		var id string
		var name string
		var personalInfo PersonalInfoAttr
		if err = rows.Scan(&id, &name, &personalInfo); err != nil {
			return nil, err
		}

		personRecord := PersonRecord{
			ID:           id,
			Name:         name,
			PersonalInfo: personalInfo,
		}

		personRecords = append(personRecords, personRecord)
	}

	return personRecords, nil
}

func (r *Repository) GetPersonsUsingScany(ctx context.Context) ([]PersonRecord, error) {
	query := "SELECT id, name, personal_info FROM person;"

	var personRecords []PersonRecord
	if err := pgxscan.Select(ctx, r.db, &personRecords, query); err != nil {
		return nil, err
	}

	return personRecords, nil
}
