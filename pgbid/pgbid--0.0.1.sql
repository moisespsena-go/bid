CREATE OR REPLACE FUNCTION pgbid_get_utc_time(bid BYTEA)
RETURNS TIMESTAMP AS $$
SELECT to_timestamp(
        (get_byte(bid,0)<<24) +
        (get_byte(bid,1)<<16) +
        (get_byte(bid,2)<<8) +
        (get_byte(bid,3))
    )::TIMESTAMPTZ AT TIME ZONE 'UTC'
$$ LANGUAGE SQL IMMUTABLE STRICT;

CREATE OR REPLACE FUNCTION pgbid_get_utc_date(bid BYTEA) RETURNS DATE AS $$
SELECT pgbid_get_utc_time(bid)::DATE
$$ LANGUAGE SQL IMMUTABLE STRICT;

CREATE OR REPLACE FUNCTION pgbid_get_time(bid BYTEA)
  RETURNS TIMESTAMPTZ
AS $$
SELECT (pgbid_get_utc_time(bid) AT TIME ZONE 'UTC')::TIMESTAMPTZ
$$ LANGUAGE SQL IMMUTABLE STRICT;

CREATE OR REPLACE FUNCTION pgbid_get_date(bid BYTEA) RETURNS DATE AS $$
SELECT pgbid_get_time(bid)::DATE
$$ LANGUAGE SQL IMMUTABLE STRICT;

CREATE OR REPLACE FUNCTION pgbid_to_text(b BYTEA) RETURNS TEXT AS
$$
    SELECT rtrim(replace(replace(encode(b, 'base64'), '/', '_'), '+', '-'), '=');
$$ LANGUAGE SQL IMMUTABLE STRICT;

CREATE OR REPLACE FUNCTION pgbid_to_bytea(t TEXT) RETURNS BYTEA AS
$$
	-- Base64 encode length: (BYTEA_LENGTH + 2) / 3 * 4 -> (12 + 2) / 3 * 4 = 16
    SELECT decode(rpad(replace(replace(t, '_', '/'), '-', '+'), 16, '='), 'base64');
$$ LANGUAGE SQL IMMUTABLE STRICT;