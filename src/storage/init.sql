CREATE SCHEMA data;

CREATE TABLE data.candles (
  id SERIAL PRIMARY KEY,
  ticker VARCHAR(10) NOT NULL,
  open NUMERIC(10, 2) NOT NULL,
  close NUMERIC(10, 2) NOT NULL,
  min NUMERIC(10, 2) NOT NULL,
  max NUMERIC(10, 2) NOT NULL,
  volume BIGINT NOT NULL
);

GRANT USAGE ON SCHEMA data TO invest_user;
GRANT ALL PRIVILEGES ON TABLE data.candles TO invest_user;
GRANT ALL PRIVILEGES ON SEQUENCE data.candles_id_seq TO invest_user;