create table if not exists product (
    id bigint unsigned primary key auto_increment,
    name varchar(10) not null,
    price float not null default '0.0',
    created_at datetime not null default current_timestamp,
    updated_at datetime null on update current_timestamp
); 