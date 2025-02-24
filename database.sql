-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

-- This is test table. Remove this table and replace with your own tables. 
CREATE TABLE estates (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	width int4 NOT NULL,
	length int4 NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT estates_pkey PRIMARY KEY (id),
	CONSTRAINT estates_width_check CHECK (width BETWEEN 1 AND 50000),
    CONSTRAINT estates_length_check CHECK (length BETWEEN 1 AND 50000)
);

CREATE TABLE trees (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	estate_id uuid NOT NULL,
	x int4 NOT NULL,
	y int4 NOT NULL,
	height int4 NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT trees_pkey PRIMARY KEY (id),
	CONSTRAINT trees_estate_id_x_y_key UNIQUE (estate_id, x, y),
	CONSTRAINT trees_height_check CHECK (height BETWEEN 1 AND 30)
);

CREATE INDEX idx_estate_id ON trees USING btree (estate_id);
CREATE INDEX idx_tree_coords ON trees USING btree (estate_id, x, y);

ALTER TABLE trees 
	ADD CONSTRAINT trees_estate_id_fkey FOREIGN KEY (estate_id) 
	REFERENCES estates(id) ON DELETE CASCADE;