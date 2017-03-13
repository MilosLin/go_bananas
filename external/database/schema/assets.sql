CREATE TABLE IF NOT EXISTS `test_db`.`assets` (
    item_name varchar(32) primary key, -- A common attribute for all items
    dynamic_cols  blob  -- Dynamic columns will be stored here
);
