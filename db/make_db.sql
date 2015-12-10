DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  username VARCHAR(20) PRIMARY KEY NOT NULL,
  pw VARCHAR(20) NOT NULL
);

CREATE TABLE items (
  username VARCHAR(20) NOT NULL,
  link VARCHAR(64) NOT NULL,
  title VARCHAR(64) NOT NULL, 
  postWhen TIMESTAMP NOT NULL
);


INSERT INTO users VALUES("adam_chal", "pw");
INSERT INTO users VALUES("maddydell", "pw");

INSERT INTO items VALUES("adam_chal", "https://adamchalmers.github.io", "My site", '2010-08-28T13:40:02.200');
INSERT INTO items VALUES("adam_chal", "http://www.theverge.com", "The Verge", '2011-07-28T13:40:02.200');
INSERT INTO items VALUES("maddydell", "http://en.wikipedia.org/wiki/Lena_Heady", "Queen", '1999-11-28T13:40:02.200');