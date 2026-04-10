CREATE TABLE learngo.td_user (
user_id uuid NOT NULL DEFAULT learngo.uuid_generate_v7(),
username varchar(50) NOT NULL,
"password" varchar(50) NOT NULL,
email varchar(100) NULL,
created_at timestamptz NULL DEFAULT now(),
CONSTRAINT td_user_pkey PRIMARY KEY (user_id)
);