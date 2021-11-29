create table users
(
    id      bigint primary key generated always as identity,
    name    varchar(200) not null,
    surname varchar(200) not null,
    active  boolean      not null default true
);

create table authors
(
    id      bigint primary key generated always as identity,
    name    varchar(200) not null,
    surname varchar(200) not null
);

create table books
(
    id        bigint primary key generated always as identity,
    title     varchar(999) not null,
    author_id integer,
    constraint author_id_fkey foreign key (author_id) references authors (id)
);

create table users_books_rates
(
    date    timestamp with time zone default current_timestamp,
    user_id integer not null,
    book_id integer not null,
    rate    integer not null,
    constraint user_book_pkey primary key (user_id, book_id),
    constraint user_id_fkey foreign key (user_id) references users (id),
    constraint book_id_fkey foreign key (book_id) references books (id),
    constraint users_books_rate_between_0_and_10 check (rate > 0 and rate <= 10)
);

insert into users (name, surname)
values ('Alex', 'Tonkonogov'),
       ('John', 'Doe'),
       ('Jane', 'Doe');

insert into authors (name, surname)
values ('Федор', 'Достоевский'),
       ('Айн', 'Рэнд'),
       ('Джек', 'Лондон'),
       ('Стивен', 'Кинг');

insert into books (title, author_id)
values ('Преступление и наказание', 1),
       ('Источник', 2),
       ('Атлант расправил плечи', 2),
       ('Мартин Иден', 3),
       ('Мертвая зона', 4),
       ('11.22.63', 4),
       ('Темная башня', 4);

insert into users_books_rates (user_id, book_id, rate)
values (1, 1, 10),
       (1, 2, 10),
       (1, 3, 9),
       (1, 4, 10),
       (1, 5, 8),
       (1, 6, 9),
       (1, 7, 7),
       (2, 1, 10),
       (2, 2, 9),
       (2, 3, 10),
       (2, 4, 8),
       (2, 5, 8),
       (2, 6, 9),
       (2, 7, 8),
       (3, 1, 9),
       (3, 2, 10),
       (3, 3, 9),
       (3, 4, 9),
       (3, 5, 9),
       (3, 6, 9),
       (3, 7, 7);

create index users_surname_idx on users (surname);
create index authors_surname_idx on authors (surname);
create index books_title_idx on books using btree (title text_pattern_ops);