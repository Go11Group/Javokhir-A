
create table book
(
    id          serial primary key,
    name        varchar not null,
    page        int,
    author_id   int references author(id)
);

insert into author (name) values('J.K Rowling')
insert into author (name) values ('John Doe');

insert into book(name, page, author_id) values('Harry Potter', 1234, 1)
insert into book(name, page, author_id) values('Grokking algo', 327, 2)

book_store=# insert into book(name, page, author_id) values('Harry Potter', 327, 3);
ERROR:  insert or update on table "book" violates foreign key constraint "book_author_id_fkey"
DETAIL:  Key (author_id)=(3) is not present in table "author".
book_store=# 

book_store=# delete from author where id = 1;
ERROR:  update or delete on table "author" violates foreign key constraint "book_author_id_fkey" on table "book"
DETAIL:  Key (id)=(1) is still referenced from table "book".


INSERT INTO author (name) VALUES ('Shekspeere');
INSERT INTO author (name) VALUES ('qwerty');
INSERT INTO author (name) VALUES ('qaz');
INSERT INTO author (name) VALUES ('wsdcvbg');

insert into book(name, page, author_id) values('vfd', 327, 2)
insert into book(name, page, author_id) values('ebg', 327, 2)
insert into book(name, page, author_id) values('sddfs', 327, 1)
insert into book(name, page, author_id) values('dftghy', 327, 4)
insert into book(name, page, author_id) values('dfvfdb', 327, 3)
insert into book(name, page, author_id) values('fe', 327, 5)
insert into book(name, page, author_id) values('fewf', 327, 5)
insert into book(name, page, author_id) values('fewv', 327, 3)
insert into book(name, page, author_id) values('hgasfd', 327, 1)
insert into book(name, page, author_id) values('vds', 327, 6)


