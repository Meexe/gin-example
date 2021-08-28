create table workers(
	id integer generated always as identity primary key,
 	department_id integer not null references departments(id),
  	name varchar not null default '',
    phone varchar not null default '',
    email varchar not null default '',
  	is_supervisor boolean not null default false,
  	position varchar not null default ''
);