create table users
(
    Id       integer not null /*autoincrement needs PK*/,
    UserId   text    not null,
    Username text    not null,
    password text    not null
);

create unique index users_Id_uindex
    on users (Id);

create unique index users_UserId_uindex
    on users (UserId);

create unique index users_Username_uindex
    on users (Username);
