create table c_course(
                         id int(32) not null auto_increment,
                         name varchar(255) not null unique ,
                         avatar varchar(255),
                         hours int(32) not null,
                         sid int(32) not null,
                         primary key(id)
)
    engine=innodb default charset=utf8mb4;
alter table c_course add index(name);
