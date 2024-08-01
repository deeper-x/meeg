


### Migrations:
```sql

CREATE TABLE databases (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    ts_created timestamp default now()
);

CREATE TABLE migrations (
    id SERIAL PRIMARY KEY,
    query text NOT NULL,
    fk_database int REFERENCES databases(id) NOT NULL,
    ts_created timestamp DEFAULT NOW()
);

CREATE TABLE hosts (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    ts_created timestamp DEFAULT NOW()
);

CREATE TABLE registrations (
    id SERIAL PRIMARY KEY,
    fk_migration int REFERENCES migrations(id) NOT NULL,
    fk_host int REFERENCES hosts(id) NOT NULL,
    ts_created timestamp DEFAULT NOW()
);




```