drop table apps;
create table apps(
    ID integer primary key autoincrement,
    ApkMd5         nvarchar(256),
	ApkURL         nvarchar(256),
	ApkPublishTime integer,
	AppDownCount   integer,
	AppName        nvarchar(256),
	AuthorName     nvarchar(256),
	AverageRating  float,
	CategoryID     integer,
	CategoryName   nvarchar(256),
	Description    nvarchar(1024),
	EditorIntro    nvarchar(1024),
	FileSize       integer,
	PkgName        nvarchar(256),
	VersionCode    integer,
	VersionName    nvarchar(256),
    UsedSDK        nvarchar(1024),
	Status         nvarchar(1024),
	RankId	       integer,
	UpdateTime     nvarchar(256)
);