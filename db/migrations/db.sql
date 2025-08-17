CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  nickname VARCHAR UNIQUE,
  fullname VARCHAR,
  about VARCHAR,
  email VARCHAR UNIQUE,
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE forums (
  id SERIAL PRIMARY KEY,
  title VARCHAR,
  "user" VARCHAR NOT NULL REFERENCES users(nickname) ON DELETE CASCADE,
  slug VARCHAR UNIQUE,
  posts INTEGER DEFAULT 0,
  threads INTEGER DEFAULT 0,
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE threads (
  id SERIAL PRIMARY KEY,
  title VARCHAR UNIQUE,
  author VARCHAR NOT NULL REFERENCES users(nickname) ON DELETE CASCADE,
  forum VARCHAR NOT NULL REFERENCES forums(slug) ON DELETE CASCADE,
  "message" VARCHAR,
  votes INTEGER,
  slug VARCHAR UNIQUE
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 

CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  parent INTEGER,
  author VARCHAR NOT NULL REFERENCES users(nickname) ON DELETE CASCADE,
  "message" TEXT,
  is_edited BOOL,
  forum VARCHAR NOT NULL REFERENCES forums(slug) ON DELETE CASCADE,
  thread INTEGER NOT NULL REFERENCES threads(id) ON DELETE CASCADE,
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_forum_threads_count()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        UPDATE forums SET threads = threads + 1 WHERE slug = NEW.forum;
        RETURN NEW;
    ELSIF (TG_OP = 'DELETE') THEN
        UPDATE forums SET threads = threads - 1 WHERE slug = OLD.forum;
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER forum_threads_count_trigger
AFTER INSERT OR DELETE ON threads
FOR EACH ROW
EXECUTE FUNCTION update_forum_threads_count();

CREATE OR REPLACE FUNCTION update_forum_posts_count()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        UPDATE forums SET posts = posts + 1 WHERE slug = NEW.forum;
        RETURN NEW;
    ELSIF (TG_OP = 'DELETE') THEN
        UPDATE forums SET posts = posts - 1 WHERE slug = OLD.forum;
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER forum_posts_count_trigger
AFTER INSERT OR DELETE ON posts
FOR EACH ROW
EXECUTE FUNCTION update_forum_posts_count();
