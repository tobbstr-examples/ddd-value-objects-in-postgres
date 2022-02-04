package code

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

var (
	personID1 = "1b0cdc83-a8d2-41ed-9592-37cd703fd932"

	stdDB = func(t *testing.T) *pgxpool.Pool {
		return setup(t)
	}
)

func setup(t *testing.T) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig("host=localhost port=5432 dbname=ddd user=tobbstr password=12345")
	require.NoError(t, err)

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	require.NoError(t, err)

	return conn
}

func TestRepository_UpsertPerson(t *testing.T) {
	type fields struct {
		db func(*testing.T) *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		person PersonRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should insert record",
			fields: fields{
				db: stdDB,
			},
			args: args{
				ctx: context.Background(),
				person: PersonRecord{
					ID:   personID1,
					Name: "person-record-1",
					PersonalInfo: PersonalInfoAttr{
						SelfieURLs: []string{"https://example1.com", "https://example2.com"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should insert record2",
			fields: fields{
				db: stdDB,
			},
			args: args{
				ctx: context.Background(),
				person: PersonRecord{
					ID:   personID1,
					Name: "person-record-1",
					PersonalInfo: PersonalInfoAttr{
						SelfieURLs: []string{"https://example1.com", "https://example2.com"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			require := require.New(t)
			r := &Repository{
				db: tt.fields.db(t),
			}

			// When
			err := r.UpsertPerson(tt.args.ctx, tt.args.person)

			// Then
			if tt.wantErr {
				require.Error(err)
				return
			}

			require.NoError(err)
		})
	}
}

func TestRepository_GetPersons(t *testing.T) {
	type fields struct {
		db func(*testing.T) *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []PersonRecord
		wantErr bool
	}{
		{
			name: "should return all persons",
			fields: fields{
				db: stdDB,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			require := require.New(t)
			r := &Repository{
				db: tt.fields.db(t),
			}

			// When
			got, err := r.GetPersons(tt.args.ctx)

			// Then
			if tt.wantErr {
				require.Error(err)
			}

			require.NoError(err)

			for _, person := range got {
				fmt.Printf("\n\nid = %s, name = %s, personalInfo = %#v", person.ID, person.Name, person.PersonalInfo)
			}
		})
	}
}
