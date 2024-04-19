CREATE TABLE IF NOT EXISTS users (
                                     id bigserial PRIMARY KEY,
                                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(), name text NOT NULL,
                                     email varchar(255) UNIQUE NOT NULL,
                                     password_hash bytea NOT NULL,
                                     activated bool NOT NULL,
                                     version integer NOT NULL DEFAULT 1
);

CREATE TABLE messages
(
    "id"       bigserial PRIMARY KEY,
    "message" varchar(255) NOT NULL
)