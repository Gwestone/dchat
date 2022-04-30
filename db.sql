BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users"
(
    Id       integer
        constraint users_pk
            primary key autoincrement,
    UserId   text not null,
    Username text not null,
    Password text not null
);
/*INSERT INTO "users" VALUES(1,'4a0785f5-9720-41ca-8618-5cf7ffb69610','user1','secret1');
INSERT INTO "users" VALUES(2,'87f26fbf-5d5b-4cb3-90cb-30597fb2b860','uniquie1','secret1');*/
CREATE TABLE IF NOT EXISTS "messages"
(
    Id      integer not null
        constraint messages_pk
            primary key autoincrement,
    "from"  text    not null,
    "to"    text    not null,
    message text    not null,
    date    integer not null,
    constraint messages_users_Username_Username_fk
        foreign key ("from", "to") references users (Username, Username)
            on update cascade on delete cascade
);
/*INSERT INTO "messages" VALUES(1,'1','2','Hello world',1648133278);
INSERT INTO "messages" VALUES(2,'1','2','hello userUniquie1',1648136683);*/
CREATE TABLE secrets
(
    id     integer not null
        constraint secrets_pk
            primary key autoincrement,
    "from" integer not null,
    "to"   integer not null,
    secret text    not null,
    constraint secrets_users_Id_Id_fk
        foreign key ("from", "to") references users (Id, Id)
);
CREATE UNIQUE INDEX users_Id_uindex
    on users (Id);
CREATE UNIQUE INDEX users_UserId_uindex
    on users (UserId);
CREATE UNIQUE INDEX users_Username_uindex
    on users (Username);
CREATE UNIQUE INDEX messages_Id_uindex
    on messages (Id);
CREATE UNIQUE INDEX secrets_id_uindex
    on secrets (id);
COMMIT;
PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
COMMIT;