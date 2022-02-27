drop table if exists unit;
create table unit
(
    id     serial primary key,
    name   text default 'Kai',
    grade  text default 'Common',
    status text default 'Idle',
    level  int  default 0,
    xp     int  default 0
);
select *
from unit order by id;

insert into unit (name)
values ('Miner'), ('Guard'), ('Healer'),
       ('Adventure'), ('Blacksmith'), ('Magician');


drop table if exists unit_info;
create table unit_info
(
    unit_id      int references unit(id),
    attack       int                      default 1,
    defense      int                      default 0,
    mana         int                      default 0,
    fraction     text                     default 'main',
    location     text                     default null,
    targeting    int references unit (id) default null,
    container_id text                     default null
);
select *
from unit_info;



drop table if exists items;
create table items
(
    id           serial primary key,
    name         text,
    amount       int default 0,
    amount_limit int default 1000000
);
select id, name, amount
from items;

insert into items (name)
values ('Rock'),
       ('Mana_crystal'),
       ('Iron_ore');

update items
set amount = amount +1
where id = 1;


drop table if exists levels;
create table levels (
    level int default null,
    xp int default null,
    total int default null
);
select * from levels;

insert into levels(level, xp, total)
values (1, 0, 0);