alter table addresses drop column if exists address;

alter table addresses add city text not null default 'Москва';
alter table addresses add street text not null default '2-я Бауманская';
alter table addresses add house text not null default '1';
alter table addresses add flat text not null default '501';
