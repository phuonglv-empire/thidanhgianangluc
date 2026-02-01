-- Initialize database for online exam system
-- This script is run automatically by Docker on first startup

-- Create extensions if needed
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Ensure database is UTF8
ALTER DATABASE examdb SET timezone TO 'UTC';

-- Create initial schema (tables will be created by GORM migrations)
