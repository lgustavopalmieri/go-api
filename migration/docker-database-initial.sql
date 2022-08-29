CREATE TABLE personalities(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  history VARCHAR(255)
);

INSERT INTO personalities(name, history)
VALUES('Tenent Dan', 'One of the major captains on earth')