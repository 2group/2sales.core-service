ALTER TABLE suborders ADD COLUMN order_index INTEGER;

WITH ordered_suborders AS (
    SELECT 
        sub_order_id, 
        ROW_NUMBER() OVER (PARTITION BY to_organization_id ORDER BY status ASC, sub_order_id ASC) AS row_num
    FROM suborders
)
UPDATE suborders s
SET order_index = o.row_num
FROM ordered_suborders o
WHERE s.sub_order_id = o.sub_order_id;
