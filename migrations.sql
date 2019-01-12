create table users
(
	id int auto_increment,
	email varchar(255) not null,
	password varchar(255) not null,
	constraint users_pk
		primary key (id)
);

create unique index users_email_uindex
	on users (email);

create table scopes
(
	id int auto_increment,
	user_id int not null,
	scope varchar(255) not null,
	constraint scopes_pk
		primary key (id)
);

create index scopes_scope_index
	on scopes (scope);

create index scopes_user_id_index
	on scopes (user_id);

