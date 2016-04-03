CREATE TABLE frames (
  id          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  content     text,
  editor      text NOT NULL,
  frame_taken timestamp without time zone NOT NULL,
  file_id     uuid NOT NULL,
  inserted_at timestamp without time zone NOT NULL DEFAULT now()
);

ALTER TABLE frames ADD CONSTRAINT frames_file_id_fkey FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE;
