create schema rater;
	
create table rater.user(
    id serial primary key,
    name varchar not null,
    email varchar not null unique,
    password varchar not null,
    active boolean default true,
	profile_picture varchar default 'https://drive.google.com/file/d/1QX9Zhj1eFTjDOSqQIRpBCm8IHjwPO8om/view?usp=share_link',
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);

create type rater.item_category as enum(
	'eletronic',
	'book',
	'furniture',
	'video_game',
	'board_game',
	'clothe',
	'vehicle'
);

create table rater.item(
	id serial primary key,
	user_id integer not null,
	name varchar not null,
	category rater.item_category not null,
	brand_name varchar not null,
	average_rating numeric(2, 1) default 0.0,
	price numeric(11, 2) default 0.0,
	active boolean default true,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default current_timestamp not null,
	constraint fk_user_id
		foreign key (user_id)
			references rater.user(id)
);

create table rater.review(
	id serial primary key,
	user_id integer not null,
	item_id integer not null,
	stars numeric(2, 1) default 0.0,
	"text" text,
	active boolean default true,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default current_timestamp not null,	
	constraint fk_user_id
		foreign key (user_id)
			references rater.user(id),
	constraint fk_item_id
		foreign key (item_id)
			references rater.item(id)
);

