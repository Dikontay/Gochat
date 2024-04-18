CREATE TABLE "users" (
                         id bigserial PRIMARY KEY,
                         username varchar(20) NOT NULL,
                         email varchar(20) NOT NULL,
                         password varchar(200) NOT NULL,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         last_login TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "messages"
(
    "id"       bigserial PRIMARY KEY,
    "message" varchar(255) NOT NULL
)