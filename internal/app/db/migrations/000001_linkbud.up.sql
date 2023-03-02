CREATE TABLE IF NOT EXISTS users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    short_title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE IF NOT EXISTS users_lists
(
    id         serial                                      not null unique,
    user_id    int references users (id) on delete cascade not null,
    list_id    int references lists (id) on delete cascade not null,
    list_title varchar                                     not null
);

CREATE TABLE IF NOT EXISTS links
(
    id          serial       not null unique,
    title       varchar(255) not null,
    short_title varchar(255) not null,
    url         varchar(255) not null
);

CREATE TABLE IF NOT EXISTS lists_links
(
    id         serial                                      not null unique,
    list_id    int references lists (id) on delete cascade not null,
    link_id    int references links (id) on delete cascade not null,
    link_title varchar                                     not null
);