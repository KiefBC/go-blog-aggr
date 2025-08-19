-- +goose up
CREATE TABLE posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  title TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  description TEXT NOT NULL,
  published_at TIMESTAMP NOT NULL,
  feed_id UUID NOT NULL,
  CONSTRAINT posts_feed_id_fkey
    FOREIGN KEY (feed_id)
    REFERENCES feeds(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose down
DROP TABLE posts;
