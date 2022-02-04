CREATE TABLE person (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    personal_info JSONB NOT NULL
);
