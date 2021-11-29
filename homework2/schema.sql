--Добавить ограничения foreign keys для всех имеющихся связей между таблицами в БД, созданной в первом занятии. 
--Ограничения можно добавлять как в существующую таблицу (используя alter), так и изменив команду создания таблиц из прошлого урока. 
--Команды на добавление ограничений описать в файле schema.sql (редактировать файл из прошлого урока).
--Выявить необходимые ограничения (constraints) и добавить их в структуру базы данных. 
--Примером может быть ограничение на неотрицательность з/п. Как и в пункте 1, делайте это в файле schema.sql.

create table users (
	id bigint primary key generated always as identity,
	name varchar(200) not null,
	surname varchar(200) not null,
	active boolean not null default true
);

create table authors (
	id bigint primary key generated always as identity,
	name varchar(200) not null,
	surname varchar(200) not null
);

create table books (
	id bigint primary key generated always as identity,
	title varchar(999) not null,
	author_id integer,
	constraint author_id_fkey foreign key (author_id) references authors (id)
);

create table users_books_rates (
    date timestamp with time zone default current_timestamp,
    user_id integer not null,
    book_id integer not null,
    rate integer not null,
    constraint user_book_pkey primary key(user_id, book_id),
    constraint user_id_fkey foreign key (user_id) references users (id),
    constraint book_id_fkey foreign key (book_id) references books (id),
    constraint users_books_rate_between_0_and_10 check (rate > 0 and rate <= 10)
);
