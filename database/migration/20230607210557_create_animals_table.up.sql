CREATE TABLE IF NOT EXISTS animals (
    id bigserial primary key,
    animal varchar(255) not null,
    breed varchar(255) null,
    gender varchar(255) null,
    age integer not null,
    size varchar(255) null,
    city varchar(255) null,
    state varchar(255) null,
    dewormed varchar(255) null,
    castrated varchar(255) null,
    vaccinated varchar(255) null,
    special_care varchar(255) null,
    picture varchar(255) null,
    created_at timestamp not null,
    udpated_at timestamp not null,
    deleted_at timestamp null
);