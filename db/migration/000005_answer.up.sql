create table "answer" (
  "id" serial primary key,
  "question_id" serial not null,
  "answer" text not null,
  "created_by" serial not null,
  "created_at" timestamp not null default now(),
  "updated_at" timestamp not null default now(),
  "deleted_at" timestamp null
);
ALTER TABLE "answer" ADD FOREIGN KEY ("question_id") REFERENCES "question" ("id");