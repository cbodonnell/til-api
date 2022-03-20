-- Initialization Script

DO $$
BEGIN
	-- Create tables

	CREATE TABLE IF NOT EXISTS public.tils (
		"id" serial NOT NULL,
		"text" text NOT NULL,
		"user_uuid" text NOT NULL,
		CONSTRAINT tils_pk PRIMARY KEY ("id")
	);

END
$$

