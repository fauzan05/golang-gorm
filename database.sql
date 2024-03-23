create database golang_gorm;
create database golang_gorm_test;

use golang_gorm;
use golang_gorm_test;

show tables;
select * from todos;
select * from users;
select * from schema_migrations;
drop table schema_migrations;
drop table todos;
drop table users;

select * from users where id = "1" for update;