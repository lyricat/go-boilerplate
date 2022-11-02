UPDATE
  "properties"
SET
  "value"=:value,
  "updated_at"=NOW()
WHERE
  "key"=:key
;