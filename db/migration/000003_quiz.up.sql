create table "quiz" (
  "id" serial primary key,
  "name" varchar(255) not null,
  "description" text not null,
  "created_by" serial not null,
  "created_at" timestamp not null default now(),
  "updated_at" timestamp not null default now(),
    "deleted_at" timestamp null
);

ALTER TABLE "quiz" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");