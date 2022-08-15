create table users
(
    id    int auto_increment,
    code  int          not null,
    name  varchar(400) not null,
    email varchar(400) not null,
    constraint user_pk
        primary key (id)
);


create table artists
(
    id   int auto_increment,
    code int          not null,
    name varchar(400) not null,
    constraint artist_pk
        primary key (id)
);

create table songs
(
    id     int auto_increment,
    code   int          not null,
    name   varchar(400) not null,
    artistId int          not null,
    constraint song_pk
        primary key (id),
    constraint song_artist_fk
        foreign key (artistId) references artists (id)
            on update cascade on delete cascade
);
