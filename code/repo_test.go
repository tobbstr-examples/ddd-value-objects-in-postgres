package code

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

var (
	personID1 = "1b0cdc83-a8d2-41ed-9592-37cd703fd932"
	personID2 = "2b0cdc83-a8d2-41ed-9592-37cd703fd932"

	person1 = PersonRecord{
		ID:   personID1,
		Name: "person-1",
		PersonalInfo: PersonalInfoAttr{
			SelfieURLs: []string{"https://example.com/persons/1/selfies/1", "https://example.com/persons/1/selfies/2"},
		},
	}
	person2 = PersonRecord{
		ID:   personID2,
		Name: "person-2",
		PersonalInfo: PersonalInfoAttr{
			SelfieURLs: []string{"https://example.com/persons/2/selfies/1"},
		},
	}

	teardownQueries = []string{
		"TRUNCATE person;",
	}

	stdDB = func(t *testing.T) *pgxpool.Pool {
		return setup(t)
	}
)

func setup(t *testing.T) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig("host=localhost port=5432 dbname=ddd user=tobbstr password=12345")
	require.NoError(t, err)

	ctx := context.Background()

	conn, err := pgxpool.ConnectConfig(ctx, config)
	require.NoError(t, err)

	for _, query := range teardownQueries {
		if _, err := conn.Exec(ctx, query); err != nil {
			require.NoError(t, err)
		}
	}

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
			name: "should return nil when inserting person1",
			fields: fields{
				db: stdDB,
			},
			args: args{
				ctx:    context.Background(),
				person: person1,
			},
			wantErr: false,
		},
		{
			name: "should return nil when inserting person2",
			fields: fields{
				db: stdDB,
			},
			args: args{
				ctx:    context.Background(),
				person: person2,
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
			defer r.db.Close()

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
		setup   func(context.Context, *Repository)
	}{
		{
			name: "should return 2 persons when 2 exists",
			fields: fields{
				db: stdDB,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    []PersonRecord{person1, person2},
			wantErr: false,
			setup: func(ctx context.Context, r *Repository) {
				r.UpsertPerson(ctx, person1)
				r.UpsertPerson(ctx, person2)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			require := require.New(t)
			r := &Repository{
				db: tt.fields.db(t),
			}
			tt.setup(tt.args.ctx, r)
			defer r.db.Close()

			// When
			got, err := r.GetPersons(tt.args.ctx)

			// Then
			if tt.wantErr {
				require.Error(err)
			}

			require.NoError(err)

			require.Len(got, len(tt.want))
			require.Equal(tt.want[0], got[0])
			require.Equal(tt.want[1], got[1])
		})
	}
}

func TestRepository_GetPersonsUsingScany(t *testing.T) {
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
		setup   func(context.Context, *Repository)
	}{
		{
			name: "should return 2 persons when 2 exists",
			fields: fields{
				db: stdDB,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    []PersonRecord{person1, person2},
			wantErr: false,
			setup: func(ctx context.Context, r *Repository) {
				r.UpsertPerson(ctx, person1)
				r.UpsertPerson(ctx, person2)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			require := require.New(t)
			r := &Repository{
				db: tt.fields.db(t),
			}
			tt.setup(tt.args.ctx, r)
			defer r.db.Close()

			// When
			got, err := r.GetPersonsUsingScany(tt.args.ctx)

			// Then
			if tt.wantErr {
				require.Error(err)
			}

			require.NoError(err)

			require.Len(got, len(tt.want))
			require.Equal(tt.want[0], got[0])
			require.Equal(tt.want[1], got[1])
		})
	}
}
