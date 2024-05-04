CREATE TABLE IF NOT EXISTS "time_tables" (
  id UUID PRIMARY KEY,
  teacher_id UUID NOT NULL,
  student_id UUID NOT NULL,
  subject_id UUID NOT NULL,
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL
);


ALTER TABLE time_tables ADD FOREIGN KEY (teacher_id) REFERENCES teachers (id);
ALTER TABLE time_tables ADD FOREIGN KEY (student_id) REFERENCES students (id);
-- ALTER TABLE time_tables ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);
