CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS user;

CREATE TABLE users(
    id int auto_increment primary key,
    user_name varchar(50) not null,
    username varchar(50) not null unique,
    email varchar(50) not null unique,
    user_password varchar(20) not null,
    createdAt timestamp default current_timestamp(),
    updatedAt timestamp default current_timestamp()
) ENGINE = INNODB;