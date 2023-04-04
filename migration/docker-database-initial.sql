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
	'video_games',
	'board_games',
	'clothes',
	'vehicles'
);

create table rater.item(
	id serial primary key,
	user_id integer not null,
	brand_id integer not null,
	name varchar not null,
	category rater.item_category not null,
	average_rating float default 0.0,
	active boolean default true,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default current_timestamp not null,
	constraint fk_user_id
		foreign key (user_id)
			references rater.user(id),
	constraint fk_brand_id
		foreign key (brand_id)
			references rater.brand(id) 
);

create table rater.review(
	id serial primary key,
	user_id integer not null,
	item_id integer not null,
	stars float default 0.0,
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

create table rater.brand(
    id serial primary key,
    name varchar not null,
	average_rating float default 0.0,
    active boolean default true,
	profile_picture varchar default 'https://drive.google.com/file/d/10dg-hhQm8jhagSjcgcXbeyo49IHbum5F/view?usp=share_link',
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);
