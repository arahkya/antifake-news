create table if not EXISTS Contents(
	Id integer primary key autoincrement,
    TitleEn nvarchar(150) not null,
    TitleTh nvarchar(150) not null,
    Detail nvarchar(1000) not null,
    Author nvarchar(100) not null,
    Organize nvarchar(50) not null
)