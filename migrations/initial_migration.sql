-- Initial Migration Script

-- Migrations table
CREATE TABLE "Migrations" (
    "ID" SERIAL PRIMARY KEY,
    "migrationID" VARCHAR(255) NOT NULL UNIQUE,
    "appliedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Animals table
CREATE TABLE "Animals" (
    "ID" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "species" VARCHAR(255) NOT NULL,
    "breed" VARCHAR(255),
    "age" INT,
    "sex" VARCHAR(50),
    "description" TEXT,
    "arrivalDate" DATE NOT NULL,
    "healthStatus" VARCHAR(255),
    "sterilisationStatus" BOOLEAN,
    "chipNumber" VARCHAR(255),
    "internalNotes" TEXT,
    "reasonOnboarded" VARCHAR(255),
    "latestVaccinationDate" DATE,
    "currentLocation" VARCHAR(255),
    "status" VARCHAR(255) NOT NULL
);

-- Venues table
CREATE TABLE "Venues" (
    "ID" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "contactInfo" VARCHAR(255),
    "googleMapsLink" VARCHAR(255),
    "operatingHours" VARCHAR(255)
);

-- Adopters table
CREATE TABLE "Adopters" (
    "ID" SERIAL PRIMARY KEY,
    "cognitoID" VARCHAR(255) NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255),
    "contactInfo" VARCHAR(255),
    "registrationDate" DATE NOT NULL
);

-- Volunteers table
CREATE TABLE "Volunteers" (
    "ID" SERIAL PRIMARY KEY,
    "cognitoID" VARCHAR(255) NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255),
    "contactInfo" VARCHAR(255),
    "registrationDate" DATE NOT NULL
);

-- Donors table
CREATE TABLE "Donors" (
    "ID" SERIAL PRIMARY KEY,
    "cognitoID" VARCHAR(255) NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255),
    "contactInfo" VARCHAR(255),
    "registrationDate" DATE NOT NULL,
    "totalDonations" DECIMAL(10, 2) DEFAULT 0
);

-- Staff Members table
CREATE TABLE "StaffMembers" (
    "ID" SERIAL PRIMARY KEY,
    "cognitoID" VARCHAR(255) NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255),
    "contactInfo" VARCHAR(255),
    "registrationDate" DATE NOT NULL
);

-- Foster Parents table
CREATE TABLE "FosterParents" (
    "ID" SERIAL PRIMARY KEY,
    "cognitoID" VARCHAR(255) NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255),
    "contactInfo" VARCHAR(255),
    "registrationDate" DATE NOT NULL
);

-- Adoptions table
CREATE TABLE "Adoptions" (
    "ID" SERIAL PRIMARY KEY,
    "animalID" INT NOT NULL,
    "adopterID" INT NOT NULL,
    "adoptionDate" DATE NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "outcome" VARCHAR(255),
    FOREIGN KEY ("animalID") REFERENCES "Animals"("ID"),
    FOREIGN KEY ("adopterID") REFERENCES "Adopters"("ID")
);

-- Pop Up Days table
CREATE TABLE "PopUpDays" (
    "ID" SERIAL PRIMARY KEY,
    "venueID" INT NOT NULL,
    "date" DATE NOT NULL,
    "volunteers" JSON NOT NULL,
    "dogs" JSON NOT NULL,
    FOREIGN KEY ("venueID") REFERENCES "Venues"("ID")
);

-- Insert initial migration record
INSERT INTO "Migrations" ("migrationID") VALUES ('initial_migration');
