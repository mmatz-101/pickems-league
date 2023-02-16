CREATE TYPE "spread_winner" AS ENUM (
  'HOME',
  'AWAY',
  'PUSH',
  'UNDETERMINED'
);

CREATE TYPE "spread_picks" AS ENUM (
  'HOME',
  'AWAY'
);

CREATE TYPE "spread_type" AS ENUM (
  'FAVORITE',
  'UNDERDOG'
);

CREATE TYPE "season_type" AS ENUM (
  'PRE',
  'REG',
  'POST'
);

CREATE TYPE "league" AS ENUM (
  'NFL',
  'NCAAF'
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hash_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "picks" (
  "username" varchar NOT NULL,
  "game_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "year" int NOT NULL,
  "week" int NOT NULL,
  "league" league NOT NULL,
  "user_pick" spread_picks NOT NULL,
  "user_pick_type" spread_type NOT NULL,
  "game_spread_winner" spread_winner NOT NULL
);

CREATE TABLE "games" (
  "id" varchar PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "hometeam_fullname" varchar NOT NULL,
  "hometeam_shortname" varchar NOT NULL,
  "hometeam_logourl" varchar NOT NULL,
  "awayteam_fullname" varchar NOT NULL,
  "awayteam_shortname" varchar NOT NULL,
  "awayteam_logourl" varchar NOT NULL,
  "channel" varchar NOT NULL,
  "date" timestamp NOT NULL,
  "status" varchar NOT NULL,
  "year" int NOT NULL,
  "week" int NOT NULL,
  "weekname" varchar NOT NULL,
  "homeorunderline" float NOT NULL,
  "homeorunderodd" int NOT NULL,
  "awayoroverline" float NOT NULL,
  "awayoroverodd" int NOT NULL,
  "created_at_vegas" varchar NOT NULL,
  "sportsbookid" int NOT NULL,
  "homescore" int NOT NULL,
  "awayscore" int NOT NULL,
  "game_spread_winner" spread_winner NOT NULL,
  "season_type" season_type NOT NULL,
  "league" league NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "picks" ("username");

CREATE INDEX ON "picks" ("game_id");

-- CREATE UNIQUE INDEX ON "picks" ("username", "game_id");
ALTER TABLE "picks" ADD CONSTRAINT "picks_unique_key" UNIQUE ("username", "game_id");

ALTER TABLE "picks" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "picks" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id");
