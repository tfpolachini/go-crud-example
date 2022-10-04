create table products(
	id varchar(40) not null,
	name varchar(255) not null,
	status varchar(40) not null,
	created_at datetime not null,
	updated_at datetime not null,
	primary key (id)
);

create unique index name_index on crud.products(name);