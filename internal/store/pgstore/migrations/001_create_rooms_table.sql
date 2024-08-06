
create table if not exists rooms (
  "id" uuid                 primary key not null default gen_random_uui() --id da sala
  "theme" varchar(255)                  not null --nome da sala
);

---- create above / drop below ----

drop table if exists rooms;

