DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS follows;
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

CREATE TABLE follows (
  username VARCHAR(20) NOT NULL,
  following VARCHAR(20) NOT NULL
);


INSERT INTO users VALUES("adam_chal", "pw");
INSERT INTO users VALUES("maddydell", "pw");
INSERT INTO users VALUES("sarahj_berry", "pw");
INSERT INTO users VALUES("justipen", "pw");
INSERT INTO users VALUES("wint", "pw");
INSERT INTO users VALUES("cashbonez", "pw");

INSERT INTO items VALUES("adam_chal", "https://adamchalmers.github.io", "My site", '2010-08-28T13:40:02.200');
INSERT INTO items VALUES("adam_chal", "http://www.theverge.com", "The Verge", '2011-07-28T13:40:02.200');
INSERT INTO items VALUES("maddydell", "http://orphanblack.com", "Queen", '1999-11-28T13:40:02.200');
INSERT INTO items VALUES("maddydell", "http://dogs.org", "I need this", '2001-11-28T13:40:02.200');
INSERT INTO items VALUES("sarahj_berry", "http://getup.org", "Booooo", '2002-11-28T13:40:02.200');
INSERT INTO items VALUES("sarahj_berry", "http://buzzfeed.com/25-horror-films", "Gotta watch these", '2003-11-28T13:40:02.200');
INSERT INTO items VALUES("sarahj_berry", "http://en.wikipedia.org/wiki/The Muppets", "My ultimate fears", '2004-11-28T13:40:02.200');
INSERT INTO items VALUES("justipen", "http://grouches.wikia.com", "cool site about naaman zhou", '2005-11-28T13:40:02.200');
INSERT INTO items VALUES("justipen", "http://alj.org", "Hey I got published again", '2006-11-28T13:40:02.200');
INSERT INTO items VALUES("justipen", "http://sydney.edu.au", "Trash City", '2007-11-28T13:40:02.200');
INSERT INTO items VALUES("cashbonez", "http://watch-anime-online.net", "Love to watch anime", '2008-11-28T13:40:02.200');
INSERT INTO items VALUES("cashbonez", "http://en.wikipedia.org/wiki/Chess", "My gf", '2009-11-28T13:40:02.200');

INSERT INTO follows VALUES("adam_chal", "maddydell");
INSERT INTO follows VALUES("adam_chal", "sarahj_berry");
INSERT INTO follows VALUES("adam_chal", "justipen");
INSERT INTO follows VALUES("adam_chal", "wint");
INSERT INTO follows VALUES("justipen", "cashbonez");
INSERT INTO follows VALUES("justipen", "wint");
