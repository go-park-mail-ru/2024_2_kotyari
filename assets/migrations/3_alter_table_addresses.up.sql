alter table addresses drop column if exists city;
alter table addresses drop column if exists street;
alter table addresses drop column if exists house;
alter table addresses drop column if exists flat;

alter table addresses add address text;