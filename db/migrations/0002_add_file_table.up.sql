CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE files (
  id          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  path        text NOT NULL,
  name        text NOT NULL,
  user_id     int NOT NULL,
  inserted_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at  timestamp without time zone NOT NULL DEFAULT now()
);

ALTER TABLE files ADD CONSTRAINT files_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

