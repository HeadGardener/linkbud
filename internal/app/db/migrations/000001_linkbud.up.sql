CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    short_title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id         serial                                      not null unique,
    user_id    int references users (id) on delete cascade not null,
    list_title varchar                                     not null
);

CREATE TABLE links
(
    id    serial       not null unique,
    title varchar(255) not null,
    url   varchar(255) not null
);

CREATE TABLE lists_links
(
    id         serial  not null unique,
    list_title varchar not null,
    lnk_title  varchar not null
);