CREATE TABLE TASKS (
  id SERIAL PRIMARY KEY,
  Title VARCHAR(255) NOT NULL,
  description TEXT,
  status VARCHAR(50) DEFAULT 'pending',
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
)
