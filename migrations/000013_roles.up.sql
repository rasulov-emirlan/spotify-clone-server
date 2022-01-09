CREATE TABLE IF NOT EXISTS roles (
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	name varchar(100) NOT NULL UNIQUE
);

INSERT INTO roles(name)
VALUES('admin');
INSERT INTO roles(name)
VALUES('singer');
INSERT INTO roles(name)
VALUES('label');