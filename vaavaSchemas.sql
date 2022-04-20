drop table if exists user_;
create table user_
(
    id       serial8 primary key,
    username text not null default '',
    password text not null default ''
);
select *
from user_;
insert into user_(id)
values (default),
       (default),
       (default);

drop table if exists unit;
create table unit
(
    id         serial8,
    userID     int8 references user_ (id) not null,
    levelID    int8 references level (id) default null,
    classID    int8 references class (id) default null,
    locationID int8 references town  (id) default null,
    primary key (id)
);

select *
from unit;

insert into unit(userID)
values (1),
       (1),
       (2);
create table town
(
    id          serial8 primary key,
    name        text                                            default 'New city',
    barrier     int8 check ( town.barrier <= town.barrier_max ) default 0,
    barrier_max int8                                            default 0
);
select *
from user_;
select id, (select count(*) from unit where userid = user_.id) as units
from user_;

-- essential
-- create table user_();
-- create table unit();
-- create table town();
drop table if exists inventory;
create table inventory
(
    id     serial8 primary key,
    userID int8 not null references user_ (id),
    itemID int8 not null references item (id)
);
create table item
(
    id   serial8 primary key,
    name text unique not null
);
create table level
(
    id    serial8 primary key,
    level int8 unique not null
);
create table mine
(
);
create table dungeon
(
);
create table monster
(
    id     serial8 primary key,
    typeID int references monsterTypes (id)
);
create table monsterTypes
(
    id   serial primary key,
    name text unique not null
);

-- secondary
drop table if exists class;
create table class
(
    id   serial8 primary key,
    name text unique not null
);
create table rarity
(
);
create table magic
(
);
create table building
(
);
create table history
(
);
