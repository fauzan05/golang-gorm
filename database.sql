create database golang_gorm;
create database golang_gorm_test;

use golang_gorm;
use golang_gorm_test;

show tables;
select * from todos;
select * from users;
select * from wallets;
select * from products;
select * from user_like_product;
select * from schema_migrations;
drop table schema_migrations;
drop table todos;
drop table users;
drop table wallets;
drop table products;
drop table user_like_product;

show create table users;
show create table todos;
show create table wallets;
select * from users where id = "1" for update;