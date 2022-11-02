SELECT
  *
FROM snapshots
WHERE
  snapshot_id = :id
LIMIT 1
;