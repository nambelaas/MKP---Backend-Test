CREATE SCHEMA IF NOT EXISTS tiket_bioskop_2;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TYPE IF EXISTS tiket_bioskop_2.user_role CASCADE;
DROP TYPE IF EXISTS tiket_bioskop_2.seat_status CASCADE;
DROP TYPE IF EXISTS tiket_bioskop_2.order_status CASCADE;

CREATE TYPE tiket_bioskop_2.user_role AS ENUM ('Admin', 'User');
CREATE TYPE tiket_bioskop_2.seat_status AS ENUM ('Available', 'Hold', 'Paid', 'Cancelled', 'Expired');
CREATE TYPE tiket_bioskop_2.order_status AS ENUM ('Pending', 'Paid', 'Completed', 'Cancelled', 'Refunded');


-- ===================
-- Table: users
-- ===================
create table tiket_bioskop_2.users(
	id uuid primary key default gen_random_uuid(),
	email text not null unique,
	password_hash text not null,
	full_name text not null,
	role tiket_bioskop_2.user_role default 'User',
	created_at timestamp default current_timestamp
);

-- ===================
-- Table: movies
-- ===================
create table tiket_bioskop_2.movies(
	id uuid primary key default gen_random_uuid(),
	title text not null,
	duration_minutes integer,
	created_at timestamp default current_timestamp
);


-- ===================
-- Table: theater
-- ===================
create table tiket_bioskop_2.theaters(
	id uuid primary key default gen_random_uuid(),
	name text not null,
	city text not null,
	created_at timestamp default current_timestamp
);


-- ===================
-- Table: showtimes
-- ===================
create table tiket_bioskop_2.showtimes(
	id uuid primary key default gen_random_uuid(),
	movie_id uuid not null,
	theater_id uuid not null,
	start_at timestamp not null,
	base_price numeric(12,2) not null,
	created_at timestamp default current_timestamp,
	constraint fk_movie foreign key (movie_id) references tiket_bioskop_2.movies (id) on delete cascade,
	constraint fk_theater foreign key (theater_id) references tiket_bioskop_2.theaters (id) on delete cascade
);

-- ===================
-- Table: seats
-- ===================
create table tiket_bioskop_2.seats(
	id uuid primary key default gen_random_uuid(),
	showtime_id uuid not null,
	seat_code text not null,
	status tiket_bioskop_2.seat_status default 'Available',
	hold_until timestamp,
	order_id uuid,
	created_at timestamp default current_timestamp,
	constraint fk_showtime foreign key (showtime_id) references tiket_bioskop_2.showtimes (id) on delete cascade
);



