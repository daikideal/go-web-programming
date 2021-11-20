drop table posts cascade if exixts;
drop table comments if exixts;

create table posts (
  id  serial primary key,
  content text,
  author  varchar(255)
);
create table comments (
  id  serial primary key,
  content text,
  author  varchar(255),
  post_id integer reference posts(id)
);
