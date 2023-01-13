drop table if exists "server";
create table "server" (
	id serial primary key,
	discord_id varchar(255) not null,
	name varchar(255) not null,
	created timestamp not null default now(),
	updated timestamp not null default now()
);

drop table if exists "user";
create table "user" (
	id serial primary key,
	discord_id varchar(255) not null,
	nickname varchar(255) not null,
	created timestamp not null default now(),
	updated timestamp not null default now()
);
	
drop table if exists guild_member;
create table guild_member (
	user_id integer references "user" (id) on delete cascade,
	guild_id integer references "server" (id) on delete cascade,
	primary key (user_id, guild_id)
);
	

drop table if exists movie;
create table movie (
	id serial primary key,
	imdb_id varchar(255),
	title varchar(255) not null,
	runtime_min integer not null default 0,
	description varchar (2000) not null default '',
	released date not null,
	created timestamp not null default now()
);

drop table if exists guild_movie;
create table guild_movie (
	guild_id int references "server" (id) on delete cascade,
	movie_id int references movie (id) on delete set null,
	num_viewings int not null default 0,
	suggested_by int references "user" (id) on delete set null,
	created timestamp not null default now(),
	updated timestamp not null default now(),
	primary key (guild_id, movie_id)
);

drop table if exists event_loc;
create table event_loc (
	id serial primary key,
	guild_id int references "server" (id),
	is_vc boolean not null default false,
	loc_description varchar(255) not null
);
	
drop table if exists event;
create table event (
	id serial primary key,
	
	guild_id int references "server" (id),
	topic varchar(255) not null default '',
	location references event_loc (id),
	description varchar(1000) not null default '',
	
	start_time timestamp not null,
	end_time timestamp default null,
	
	repeating boolean not null default false,
	repeat_days int default null,
	
	completed boolean not null default false,
	
	created timestamp not null default now(),
	updated timestamp not null default now()
	
);