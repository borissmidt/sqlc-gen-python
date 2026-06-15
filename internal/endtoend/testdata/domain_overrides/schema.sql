CREATE DOMAIN job_status AS text
CHECK (
    VALUE IN (
                 'QUEUED',
                 'PENDING',
                 'RUNNING',
                 'COMPLETED',
                 'FAILED'
        )) NOT NULL;

CREATE DOMAIN positive_int AS integer
CHECK (VALUE > 0);


CREATE TABLE jobs (
          id       BIGSERIAL PRIMARY KEY,
          status   job_status NOT NULL,
          priority positive_int
);
