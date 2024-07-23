package service

const CreateSchema = `
  CREATE TABLE IF NOT EXISTS countries (
	id SERIAL PRIMARY KEY ,
	country_name TEXT NOT NULL UNIQUE
 );
	CREATE TABLE IF NOT EXISTS beers (
	id SERIAL PRIMARY KEY ,
	beer_name TEXT NOT NULL,
	creator TEXT NOT NULL,
	origin_country_id INTEGER references countries (id) NOT NULL,
	current_country_id INTEGER references countries (id) NOT NULL,
	alcohol NUMERIC(4, 2)
 );
`
