CREATE TABLE IF NOT EXISTS account (
	id SERIAL NOT NULL,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS pasta (
	id SERIAL NOT NULL,
	title VARCHAR(255),
	description TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	account_id INT references account(id),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS noodle (
	id SERIAL NOT NULL,
	content TEXT NOT NULL,
	filename varchar(255) NOT NULL,
	language varchar(255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	pasta_id INT references pasta(id) ON DELETE CASCADE,
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS redirect (
	id SERIAL NOT NULL,
	url VARCHAR,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	pasta_id INT references pasta(id),
	PRIMARY KEY(id)
);
