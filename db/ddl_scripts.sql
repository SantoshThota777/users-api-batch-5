CREATE TABLE users
(
	user_id SERIAL PRIMARY KEY,
	first_name varchar(50),
	last_name varchar(50),
	email_id varchar(50),
	date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
