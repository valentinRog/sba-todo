CREATE TABLE todos (
  id      INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  name    text    NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);