package service

const CreateSchema = `
	CREATE TABLE IF NOT EXISTS beers (
	id SERIAL PRIMARY KEY ,
	beer_name TEXT NOT NULL UNIQUE,
	creator TEXT NOT NULL,
	origin_country TEXT,
	current_country TEXT,
	alcohol NUMERIC(4, 2)
 );`
