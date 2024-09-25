-- +migrate Up
CREATE TYPE user_role AS ENUM ('admin', 'general');

CREATE TABLE profiles (
    user_id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS profiles;
DROP TYPE IF EXISTS user_role;
