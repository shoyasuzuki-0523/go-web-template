CREATE TABLE posts (
  id           SERIAL       NOT NULL PRIMARY KEY,
  name         VARCHAR(255) NOT NULL,
  created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
  updated_at   TIMESTAMPTZ  NOT NULL DEFAULT now()
);
