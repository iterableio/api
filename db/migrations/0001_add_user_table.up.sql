CREATE TABLE users (
  id          serial NOT NULL,
  email       text NOT NULL,
  inserted_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at  timestamp without time zone NOT NULL DEFAULT now(),
  token       text NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_email_idx UNIQUE (email);
ALTER TABLE users ADD CONSTRAINT users_token_idx UNIQUE (token);
