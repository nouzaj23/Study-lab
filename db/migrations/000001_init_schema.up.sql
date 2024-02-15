CREATE TABLE "quizzes" (
                           "id" bigserial PRIMARY KEY,
                           "name" varchar NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "questions" (
                             "id" bigserial PRIMARY KEY,
                             "title" varchar NOT NULL,
                             "quiz_id" bigint NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "answers" (
                           "id" bigserial PRIMARY KEY,
                           "text" varchar NOT NULL,
                           "question_id" bigint NOT NULL,
                           "is_correct" bool NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tags" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar UNIQUE NOT NULL
);

CREATE INDEX ON "quizzes" ("name");

CREATE INDEX ON "questions" ("quiz_id");

CREATE INDEX ON "answers" ("question_id");

CREATE INDEX ON "tags" ("name");

ALTER TABLE "questions" ADD FOREIGN KEY ("quiz_id") REFERENCES "quizzes" ("id") ON DELETE CASCADE;

ALTER TABLE "answers" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id") ON DELETE CASCADE;

CREATE TABLE "tags_quizzes" (
                                "tag_id" bigserial,
                                "quiz_id" bigserial,
                                PRIMARY KEY ("tag_id", "quiz_id")
);

ALTER TABLE "tags_quizzes" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE;

ALTER TABLE "tags_quizzes" ADD FOREIGN KEY ("quiz_id") REFERENCES "quizzes" ("id") ON DELETE CASCADE;

