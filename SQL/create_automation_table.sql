CREATE TABLE automation(
id VARCHAR(255) NOT NULL PRIMARY KEY,
list_id VARCHAR(255) REFERENCES list(id),
create_time VARCHAR(255),
segment_id VARCHAR(255),
start_time VARCHAR(255),
status VARCHAR(255),
title VARCHAR(255),
trigger_settings JSONB )