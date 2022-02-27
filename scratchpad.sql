CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);

select * from users;

insert into users (age, first_name, last_name, email) values (45,'michael', 'jordan', 'michael.jordan@nba.com')
insert into users (age, first_name, last_name, email) values (60,'hakeem', 'olajuwon', 'hakeem.olajuwon@nba.com')