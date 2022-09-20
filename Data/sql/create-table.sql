create database AntifakeNews;

use AntifakeNews;

create table Contents(
	Id int not null auto_increment,
    TitleEn varchar(150) not null,
    TitleTh varchar(150) not null,
    Detail varchar(1000) not null,
    Author varchar(100) not null,
    Organize varchar(50) not null,
    PRIMARY KEY(Id)
)