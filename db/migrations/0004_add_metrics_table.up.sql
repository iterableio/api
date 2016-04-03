CREATE TABLE metrics (
  id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  type          text NOT NULL,
  context       jsonb,
  assets        jsonb,
  user_id       int NOT NULL,
  previous_hash text,
  inserted_at   timestamp without time zone NOT NULL DEFAULT now()
);

ALTER TABLE metrics ADD CONSTRAINT metrics_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
