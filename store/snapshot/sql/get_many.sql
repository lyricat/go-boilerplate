SELECT
  *
FROM snapshots
WHERE
  created_at >= :from
LIMIT
  :limit
;