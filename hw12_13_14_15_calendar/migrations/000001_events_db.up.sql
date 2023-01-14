CREATE TABLE IF NOT EXISTS events
(
    id          serial       NOT NULL,
    title       varchar(255) NOT NULL,
    description text NULL,
    date_start  timestamptz   NOT NULL,
    date_end    timestamptz   NOT NULL,
    owner_id    int          NOT NULL,
    notify_in   int          NOT NULL,
    CONSTRAINT events_pk PRIMARY KEY (id)
);