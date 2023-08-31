CREATE TABLE segments
{
    id  SERIAL PRIMARY KEY,
    name varchar(255) not null
};


CREATE TABLE userSegmentPairs
{
    id  SERIAL PRIMARY KEY,
    userId int not null,
    segmentName varchar not null
};

