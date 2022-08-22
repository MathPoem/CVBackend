CREATE TABLE "university"
(
    "id"      bigserial not null PRIMARY KEY,
    "country" varchar   not null,
    "city"    varchar   not null,
    "name"    varchar   not null unique,
    "url"     varchar
);

CREATE TABLE "school"
(
    "id"            bigserial not null PRIMARY KEY,
    "name"          varchar   not null,
    "university_id" int       not null,
    "url"           varchar
);

CREATE TABLE "department"
(
    "id"        bigserial not null PRIMARY KEY,
    "name"      varchar   not null,
    "school_id" int       not null,
    "url"       varchar
);

CREATE TABLE "person"
(
    "id"            bigserial not null PRIMARY KEY,
    "department_id" int       not null,
    "first_name"    varchar   not null,
    "middle_name"   varchar   not null,
    "second_name"   varchar   not null,
    "age"           int,
    "url"           varchar,
    "first_degree"  varchar,
    "second_degree" varchar,
    "third_degree"  varchar,
    "useradd_id"    int
);

CREATE TABLE "program"
(
    "id"         bigserial not null primary key,
    "school_id"  int       not null,
    "name"       varchar   not null,
    "year_start" int       not null,
    "semester"   int       not null,
    "url"        varchar
);

CREATE TABLE "course"
(
    "id"                    bigserial not null primary key,
    "name"                  varchar   not null,
    "program_id"            int       not null,
    "credits"               int       not null,
    "hours_lecture"         int       not null,
    "hours_seminar"         int       not null,
    "estimation_in_diploma" boolean   not null,
    "exam"                  boolean   not null,
    "description"           text      not null,
    "url"                   varchar   not null
);

CREATE TABLE "lecture"
(
    "id"        bigserial not null primary key,
    "year"      int       not null,
    "person_id" int       not null,
    "course_id" int       not null,
    "url"       varchar
);

CREATE TABLE "seminar"
(
    "id"        bigserial not null primary key,
    "year"      int       not null,
    "person_id" int       not null,
    "course_id" int       not null,
    "url"       varchar
);

CREATE TABLE users
(
    "id"           bigserial not null primary key,
    "username"     varchar   not null unique,
    "password"     varchar   not null,
    "email_1"      varchar   not null,
    "univ_1_id"    int       not null,
    "program_1_id" int       not null,
    "email_2"      varchar,
    "univ_2_id"    int,
    "program_2_id" int,
    "email_3"      varchar,
    "univ_3_id"    int,
    "program_3_id" int
);

CREATE TABLE estimations
(
    "id"          bigserial not null primary key,
    "date_create" date      not null,
    "user_id"     int       not null,
    "lecture_id"  int       not null,
    "seminar_id"  int       not null,
    "question_1"  int       not null,
    "question_2"  int       not null,
    "question_3"  int       not null,
    "question_4"  int       not null,
    "question_5"  int       not null,
    "question_6"  int       not null,
    "title"       varchar,
    "description" varchar
);

CREATE TABLE estimate_value
(
    "id" bigserial not null primary key
);

ALTER TABLE "users"
    ADD FOREIGN KEY ("univ_1_id") REFERENCES "university" ("id");

ALTER TABLE "users"
    ADD FOREIGN KEY ("univ_2_id") REFERENCES "university" ("id");

ALTER TABLE "users"
    ADD FOREIGN KEY ("univ_3_id") REFERENCES "university" ("id");

ALTER TABLE "users"
    ADD FOREIGN KEY ("program_1_id") REFERENCES "program" ("id");

ALTER TABLE "users"
    ADD FOREIGN KEY ("program_2_id") REFERENCES "program" ("id");

ALTER TABLE "users"
    ADD FOREIGN KEY ("program_3_id") REFERENCES "program" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("question_1") REFERENCES "estimate_value" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("question_2") REFERENCES "estimate_value" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("question_3") REFERENCES "estimate_value" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("question_4") REFERENCES "estimate_value" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("question_5") REFERENCES "estimate_value" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("question_6") REFERENCES "estimate_value" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("lecture_id") REFERENCES "lecture" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("seminar_id") REFERENCES "seminar" ("id");

ALTER TABLE "estimations"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "lecture"
    ADD FOREIGN KEY (person_id) REFERENCES person ("id");

ALTER TABLE "lecture"
    ADD FOREIGN KEY ("course_id") REFERENCES "course" ("id");

ALTER TABLE "seminar"
    ADD FOREIGN KEY (person_id) REFERENCES person ("id");

ALTER TABLE "seminar"
    ADD FOREIGN KEY ("course_id") REFERENCES "course" ("id");

ALTER TABLE "course"
    ADD FOREIGN KEY ("program_id") REFERENCES "program" ("id");

ALTER TABLE "program"
    ADD FOREIGN KEY ("school_id") REFERENCES "school" ("id");

ALTER TABLE person
    ADD FOREIGN KEY ("department_id") REFERENCES "department" ("id");

ALTER TABLE "school"
    ADD FOREIGN KEY ("university_id") REFERENCES "university" ("id");

ALTER TABLE "department"
    ADD FOREIGN KEY ("school_id") REFERENCES "school" ("id");