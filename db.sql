CREATE TABLE users (
  ID serial PRIMARY KEY,
  Name varchar(255),
  Surname varchar(255),
  Patronymic varchar(255),
  Age int,
  Gender varchar(10),
  Nationality varchar(255)
);