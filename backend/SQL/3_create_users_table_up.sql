create table dchat.main.users
(
    "Id"       serial
        constraint users_pk
            primary key,
    "Username" text not null,
    "Password" text not null,
    "UserId"   text not null
);

create unique index users_id_uindex
    on dchat.main.users ("Id");

create unique index users_password_uindex
    on dchat.main.users ("Password");

create unique index users_userid_uindex
    on dchat.main.users ("UserId");

create unique index users_username_uindex
    on dchat.main.users ("Username");
