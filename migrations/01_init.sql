
drop table tasks;
drop table categories;
drop type task_status;

create table departments(
	id integer generated always as identity primary key,
	department varchar,
	parent_id integer references departments(id)
);

create type task_status as enum(
	'backlog',
	'in_progress',
	'done'
);

create table tasks(
	id integer generated always as identity primary key,
	department_id integer references departments(id),
	text varchar not null default '',
	email varchar not null default '',
	phone varchar not null default '',
	status task_status not null default 'backlog'
);
