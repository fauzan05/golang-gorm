ALTER TABLE users CHANGE COLUMN `name` first_name VARCHAR(100);

ALTER TABLE users ADD COLUMN middle_name VARCHAR(100) NULL AFTER first_name;

ALTER TABLE users ADD COLUMN last_name VARCHAR(100) NULL AFTER middle_name;