create table "question" (
  "id" serial primary key,
  "quiz_id" serial not null,
  "question" text not null,
  "created_by" serial not null,
  "created_at" timestamp not null default now(),
  "updated_at" timestamp not null default now(),
  "deleted_at" timestamp null
);
ALTER TABLE "question" ADD FOREIGN KEY ("quiz_id") REFERENCES "quiz" ("id");
ALTER TABLE "question" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");