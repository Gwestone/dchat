--1_create_dchat
CREATE DATABASE dchat;
--2_create_schema_main
create schema main;
--3_create_messages_table
create table main.users
(
    "Id"       serial
        constraint users_pk
            primary key,
    "Username" text not null,
    "Password" text not null,
    "UserId"   text not null
);

create unique index users_id_uindex
    on main.users ("Id");

create unique index users_password_uindex
    on main.users ("Password");

create unique index users_userid_uindex
    on main.users ("UserId");

create unique index users_username_uindex
    on main.users ("Username");

--4_create_users_table
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

