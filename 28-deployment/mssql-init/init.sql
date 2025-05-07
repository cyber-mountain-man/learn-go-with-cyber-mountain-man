-- init.sql
-- Wait until the lesson28_mssql database exists (retry loop)
:SETVAR RETRIES 10
DECLARE @retry INT = 0;
WHILE NOT EXISTS (SELECT name FROM sys.databases WHERE name = 'lesson28_mssql')
BEGIN
    WAITFOR DELAY '00:00:05';
    SET @retry += 1;
    IF @retry >= $(RETRIES) BREAK;
END

-- Create the database if it still doesn't exist
IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = 'lesson28_mssql')
BEGIN
    CREATE DATABASE lesson28_mssql;
    WAITFOR DELAY '00:00:05'; -- Give the new DB a moment to initialize
END
GO

USE lesson28_mssql;
GO

-- Create login if it doesn't exist
IF NOT EXISTS (SELECT * FROM sys.sql_logins WHERE name = 'lesson28_user')
BEGIN
    CREATE LOGIN lesson28_user WITH PASSWORD = 'lesson28_pass';
END
GO

-- Create database user and assign role if not already created
IF NOT EXISTS (SELECT * FROM sys.database_principals WHERE name = 'lesson28_user')
BEGIN
    CREATE USER lesson28_user FOR LOGIN lesson28_user;
    EXEC sp_addrolemember 'db_owner', 'lesson28_user';
END
GO

-- Create 'users' table if it does not exist
IF NOT EXISTS (
    SELECT * FROM INFORMATION_SCHEMA.TABLES 
    WHERE TABLE_NAME = 'users' AND TABLE_SCHEMA = 'dbo'
)
BEGIN
    CREATE TABLE users (
        id INT PRIMARY KEY IDENTITY(1,1),
        name NVARCHAR(100) NOT NULL,
        email NVARCHAR(100) NOT NULL UNIQUE
    );
END
GO

-- Optional: Insert a default user if the table is empty
IF NOT EXISTS (SELECT * FROM users)
BEGIN
    INSERT INTO users (name, email)
    VALUES ('Admin', 'admin@example.com');
END
GO
