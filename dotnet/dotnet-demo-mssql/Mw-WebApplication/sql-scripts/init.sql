IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'localdb')
BEGIN
    -- If the database doesn't exist, create it
    CREATE DATABASE localdb; -- Add semicolon here
END
GO

USE localdb;
GO

IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'Person')
BEGIN
    -- If the Person table doesn't exist, create it
    CREATE TABLE Person (
        PersonID INT IDENTITY(1,1) PRIMARY KEY,
        FirstName VARCHAR(50) NOT NULL,
        LastName VARCHAR(50) NOT NULL,
        Age INT
    );

-- Insert sample data
    INSERT INTO Person (FirstName, LastName, Age)
    VALUES
        ('John', 'Doe', 30),
        ('Jane', 'Smith', 25),
        ('Michael', 'Johnson', NULL),
        ('Emily', 'Brown', 40);

END
GO

IF NOT EXISTS (SELECT * FROM sys.procedures WHERE name = 'GetAllPersons')
BEGIN
    EXEC('
            CREATE PROCEDURE GetAllPersons
            AS
            BEGIN
                SELECT * FROM Person;
            END
        ');
END
