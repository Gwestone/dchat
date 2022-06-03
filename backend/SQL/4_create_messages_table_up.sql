create table main.messages
(
    "Id"       serial
        constraint messages_pk
            primary key,
    "From"     text not null
        constraint messages_users_username_fk
            references main.users ("Username")
            on update cascade on delete cascade,
    "To"       text not null
        constraint messages_users_username_fk_
            references main.users ("Username")
            on update cascade on delete cascade,
    "MessageText" text not null,
    "Date"     int  not null
);

create unique index messages_id_uindex
    on main.messages ("Id");

