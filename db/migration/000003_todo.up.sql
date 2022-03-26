create table "todo" (
  "id" serial primary key,
  "title" varchar(500) not null,
  "description" varchar(1000),
  "is_completed" boolean not null default false,
  "user_id" integer not null,
  "created_at" timestamp not null default now(),
  "updated_at" timestamp not null default now(),
  "deleted_at" timestamp null
);

ALTER TABLE "todo" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
