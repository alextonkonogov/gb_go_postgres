--Подготовить набор данных для вашей базы. 
--Не обязательно много, главное покрыть возможные примеры использования вашей базы. 
--Описать запросы на добавление данных в data.sql файле.

insert into users (name, surname) values 
('Alex','Tonkonogov'), 
('John','Doe'), 
('Jane','Doe');

insert into authors (name, surname) values 
('Федор','Достоевский'),('Айн','Рэнд'),
('Джек','Лондон'),
('Стивен','Кинг');

insert into books (title, author_id) values 
('Преступление и наказание', 1),
('Источник', 2),
('Атлант расправил плечи', 2),
('Мартин Иден',3),
('Мертвая зона',4),
('11.22.63',4),
('Темная башня',4);

insert into users_books_rates (user_id, book_id, rate) values 
(1,1,10), (1,2,10), (1,3,9), (1,4,10), (1,5,8), (1,6,9), (1,7,7),
(2,1,10), (2,2,9), (2,3,10), (2,4,8), (2,5,8), (2,6,9), (2,7,8), 
(3,1,9), (3,2,10), (3,3,9), (3,4,9), (3,5,9), (3,6,9), (3,7,7);
