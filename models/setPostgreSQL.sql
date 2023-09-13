-- postgres
drop table cc_tag;
drop table cc_hot_list;

create table cc_tag (
  id         serial primary key,
  name       varchar(50) not null unique,
  sort       smallint(3) default null,
  source_key   varchar(50) not null unique,
  icon_color   varchar(50) default null,
  created_at timestamp not null   
);

create table cc_hot_list (
  id         serial primary key,
  tag_id     integer references cc_tag(id),
  title      varchar(200) not null,
  link      varchar(300) not null,
  extra   varchar(50) default null,
  created_at timestamp not null   
);