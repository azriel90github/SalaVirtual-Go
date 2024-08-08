create table if not exists messages (
  "id"               uuid       primary key not null    default gen_random_uui(), --id da sala
  "room_id"          uuid                   not null,   --id paa chave estrangeira
  "message"          varchar(255)           not null,   --campo para as perguntas
  "reaction_count"   bigint                 not null    default 0, --campo para as reaçoes, por default 0
  "ansewered"        boolean                not null    default false,

  foreign key (room_id) references rooms(id) --referencias para chave estrangeira
);


---- create above / drop below ----

drop table if exists messages;