create schema rater;

create table rater.customer(
    id serial primary key,
    name varchar not null,
    email varchar not null,
    password varchar not null,
    advertiser boolean default false,
    active boolean default true,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);

create type rater.item_category as enum(
	'eletronic',
	'book',
	'furniture',
	'video_games',
	'board_games'
);

create table rater.item(
	id serial primary key,
	customer_id integer not null,
	name varchar not null,
	category rater.item_category not null,
	average_rating float default 0.0,
	active boolean default true,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default current_timestamp not null,
	constraint fk_customer_id
		foreign key (customer_id)
			references rater.customer(id) 
);

