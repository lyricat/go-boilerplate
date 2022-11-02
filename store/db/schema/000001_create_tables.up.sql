BEGIN;

CREATE TABLE "assets" (
  "asset_id" varchar(255) PRIMARY KEY,
  "name" varchar(255),
  "symbol" varchar(255) UNIQUE,
  "icon_url" varchar(255),
  "chain_id" varchar(255),
  "price_usd" numeric(64,8),
  "created_at" timestamptz,
  "updated_at" timestamptz,
  "deleted_at" timestamptz,
  "order" bigint DEFAULT 0
);

CREATE INDEX idx_assets_order
ON "assets" (
  "order" ASC
);

CREATE INDEX idx_assets_symbol
ON "assets" (
  "symbol" ASC
);

CREATE TABLE "properties" (
  "key" varchar(255) PRIMARY KEY,
  "value" varchar(255),
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "snapshots" (
  "snapshot_id" varchar(36) PRIMARY KEY,
  "trace_id" varchar(36),
  "source" varchar(32),
  "transaction_hash" varchar(64),
  "receiver" varchar(256),
  "sender" varchar(256),
  "type" varchar(32),
  "user_id" varchar(36),
  "opponent_id" varchar(36),
  "asset_id" varchar(36),
  "memo" varchar(256),
  "amount" numeric(64,8),
  "created_at" timestamptz
);

COMMIT;