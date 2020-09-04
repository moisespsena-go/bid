# pgbid

BID extension for PostgreSQL

## Installation

    sudo make install

## Functions

- pgbid_get_utc_time(BYTEA):TIMESTAMP

    returns UTC TIMESTAMP from BID BYTEA.

    Example:

    ```sql
    SELECT pgbid_get_utc_time('\x5e1e56000000000000000000'::BYTEA);
    ````

    result:

         pgbid_get_utc_time

        ---------------------
         2020-01-15 00:00:00

- pgbid_get_utc_date(BYTEA):TIMESTAMP

    returns UTC DATE from BID BYTEA.

    Example:

    ```sql
    SELECT pgbid_get_utc_date('\x5e1e56000000000000000000'::BYTEA);
    ````

    result:

         pgbid_get_utc_date

        ---------------------
         2020-01-15

- pgbid_get_time(BYTEA):TIMESTAMP

    returns TIMESTAMP WITH TIME ZONE from BID BYTEA.

    Example:

    ```sql
    SELECT pgbid_get_time('\x5e1e56000000000000000000'::BYTEA);
    ````

    result:

         pgbid_get_time

        ------------------------
         2020-01-14 21:00:00-03

- pgbid_get_date(BYTEA):TIMESTAMP

    returns TIMESTAMP WITH TIME ZONE DATE from BID BYTEA. It's alias of `pgbid_get_time(bid::BYTEA)::DATE`.

    Example:

    ```sql
    SELECT pgbid_get_date('\x5e1e56000000000000000000'::BYTEA),

      pgbid_get_time('\x5e1e56000000000000000000'::BYTEA)::DATE;
    ````

    result:

          pgbid_get_date | pgbid_get_time

         ----------------+----------------
          2020-01-14     | 2020-01-14

- pgbid_to_text(BYTEA):TEXT

    returns encoded BYTEA to BASE64 URI TEXT.

    Example:

    ```sql
    SELECT pgbid_to_text('\x5e1e56000000000000000000'::BYTEA);
    ````

    result:

         pgbid_to_text

        -------------------
         Xh5WAAAAAAAAAAAA

- pgbid_to_bytea(TEXT):BYTEA

    returns BYTEA decoded from BASE64 URI TEXT.

    Example:

    ```sql
    SELECT pgbid_to_bytea('Xh5WAAAAAAAAAAAA');
    ````

    result:

         pgbid_to_bytea

        ----------------------------
         \x5e1e56000000000000000000

## Usage in Table Partition

Partition by UTC TIMESTAMP range:

    ```sql
    CREATE TABLE my_log (
        id BYTEA NOT NULL,
        message TEXT
    ) PARTITION BY RANGE (pgbid_get_utc_time(id));
    ```

Partition by UTC DATE range:

    ```sql
    CREATE TABLE my_log (
        id BYTEA NOT NULL,
        message TEXT NOT NULL
    ) PARTITION BY RANGE (pgbid_get_utc_date(id));
    ```