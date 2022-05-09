create table dchat.main.messages
(
    "Id"       serial
        constraint messages_pk
            primary key,
    "From"     text not null
        constraint messages_users_username_fk
            references dchat.main.users ("Username")
            on update cascade on delete cascade,
    "To"       text not null
        constraint messages_users_username_fk_
            references dchat.main.users ("Username")
            on update cascade on delete cascade,
    "Messages" text not null,
    "Date"     int  not null
);

create unique index messages_id_uindex
    on dchat.main.messages ("Id");

