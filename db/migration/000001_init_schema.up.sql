create table "user" (
  "id" serial primary key,
  "first_name" varchar(255) not null,
  "last_name" varchar(255) not null,
  "email" varchar(255) not null unique,
  "password" varchar(255) not null,
  "profile_image" varchar(255) default null,
  "created_at" timestamp not null default now(),
  "updated_at" timestamp not null default now(),
  "deleted_at" timestamp null
);
