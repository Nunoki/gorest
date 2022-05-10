DROP TABLE IF EXISTS "json_blob";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "json_blob" (
    "user_id"     UUID UNIQUE PRIMARY KEY,
    "content"     JSON NOT NULL, -- JSON type here serves as a validator
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "modified_at" TIMESTAMPTZ DEFAULT NULL
);


-- NOTE: The reason this ALTER command references the uuid_generate_v4 function explicitly
-- on the `public` schema, is so that the migration will work on the `test_schema` too, which
-- is created and used by integration tests. This is because creating the extension `uuid-ossp`
-- on the `test_schema` fails for a reason unknown to me at the time of writing. A PR with 
-- improvements is welcome.
ALTER TABLE "json_blob" ALTER COLUMN "user_id" SET DEFAULT "public".uuid_generate_v4();
