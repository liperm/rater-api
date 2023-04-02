create schema rater;
	
create table rater.user(
    id serial primary key,
    name varchar not null,
    email varchar not null unique,
    password varchar not null,
    advertiser boolean default false,
    active boolean default true,
	average_rating float default 0.0,
	profile_picture varchar default 'https://drive.google.com/file/d/1QX9Zhj1eFTjDOSqQIRpBCm8IHjwPO8om/view?usp=share_link',
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);

create type rater.item_category as enum(
	'eletronic',
	'book',
	'furniture',
	'video_games',
	'board_games',
	'clothes',
	'vehicles'
);

create table rater.item(
	id serial primary key,
	user_id integer not null,
	name varchar not null,
	category rater.item_category not null,
	average_rating float default 0.0,
	active boolean default true,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default current_timestamp not null,
	constraint fk_user_id
		foreign key (user_id)
			references rater.user(id) 
);

