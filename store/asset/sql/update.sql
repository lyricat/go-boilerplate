UPDATE assets
SET
  "name"=:name,
  "symbol"=:symbol,
  "icon_url"=:icon_url,
  "chain_id"=:chain_id,
  "price_usd"=:price_usd,
  "updated_at"=NOW()
WHERE
  "asset_id"=:asset_id
;

