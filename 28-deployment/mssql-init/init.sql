-- init.sql

-- Retry loop to wait for the database engine to be ready
:SETVAR RETRIES 10
DECLARE @retry INT = 0;
WHILE NOT EXISTS (SELECT name FROM sys.databases WHERE name = 'lesson28_mssql')
BEGIN
    WAITFOR DELAY '00:00:05';  -- Wait 5 seconds before next check
    SET @retry += 1;
    IF @retry >= $(RETRIES) BREAK; -- Exit after RETRIES attempts
END

-- Create the lesson28_mssql database if it does not exist
IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = 'lesson28_mssql')
BEGIN
    CREATE DATABASE lesson28_mssql;
    WAITFOR DELAY '00:00:05'; -- Pause to ensure DB is ready
END
GO

-- Switch to the target database
USE lesson28_mssql;
GO

-- Create SQL login for external connections if it doesn't exist
IF NOT EXISTS (SELECT * FROM sys.sql_logins WHERE name = 'lesson28_user')
BEGIN
    CREATE LOGIN lesson28_user WITH PASSWORD = 'lesson28_pass';
END
GO

-- Create corresponding database user and grant db_owner role
IF NOT EXISTS (SELECT * FROM sys.database_principals WHERE name = 'lesson28_user')
BEGIN
    CREATE USER lesson28_user FOR LOGIN lesson28_user;
    EXEC sp_addrolemember 'db_owner', 'lesson28_user';
END
GO

-- Create the 'users' table if it doesn't already exist
IF NOT EXISTS (
    SELECT * FROM INFORMATION_SCHEMA.TABLES 
    WHERE TABLE_NAME = 'users' AND TABLE_SCHEMA = 'dbo'
)
BEGIN
    CREATE TABLE users (
        id INT PRIMARY KEY IDENTITY(1,1),        -- Auto-incrementing ID
        name NVARCHAR(100) NOT NULL,             -- User's name
        email NVARCHAR(100) NOT NULL UNIQUE      -- User's email (must be unique)
    );
END
GO

-- Seed with an initial admin user if the table is empty
IF NOT EXISTS (SELECT * FROM users)
BEGIN
    INSERT INTO users (name, email)
    VALUES ('Admin', 'admin@example.com');
END
GO

-- Blurb: Purpose of `init.sql`
-- 
-- This script bootstraps your SQL Server instance during container setup:
--
-- - Waits for the SQL Server engine to be ready (retry loop).
-- - Creates the target database only if it doesn't exist.
-- - Creates a login and maps it to a database user with full permissions.
-- - Creates the `users` table with `id`, `name`, and `email` fields.
-- - Seeds the table with a default admin user if the table is empty.
--
-- This script is executed automatically by the `mssql-init` service in your
-- Docker Compose setup to ensure your database is ready for immediate use.
