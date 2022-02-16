create table custmer(
    S_NO int auto_increment,
    Name varchar(50) Not Null,
    mobile_number varchar(50) unique,
    email varchar(255) unique,
    address varchar(500) Not Null,
    education varchar(255) Not Null,
    primary key(S_NO)
    );