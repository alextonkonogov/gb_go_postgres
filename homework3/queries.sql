
-- поиск пользователя-читателя по фамилии (по вхождению)
select surname, name from users where surname like 'Tonk%';

-- поиск автора по фамилии (по вхождению)
select surname, name from authors where surname like 'Досто%';

-- вывод списка книг написанных конкретным автором
select 
	books.id, 
	books.title 
from books
left join authors on books.author_id = authors.id
where authors.surname like 'Кин%';


--вывод названий книг и имен авторов
select 
	books.id, books.title,
	concat(authors.name, ' ',authors.surname) as author 
from books 
left join authors on authors.id = books.author_id;


-- вывод списка читателей и оценок, которые они поставили книгам
select 
	concat(users.name, ' ',users.surname) as user,
	concat( books.title, ' ', '(', authors.name, ' ',authors.surname, ')') as book,
	rate
from users_books_rates as rates
left join users on users.id = rates.user_id
left join books on books.id = rates.book_id
left join authors on authors.id = books.author_id;
