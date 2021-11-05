create database messaging;

create extension pgcrypto;

create table users(
	user_id serial not null primary key,
	email varchar(64) not null,
	password varchar(60) not null
);

