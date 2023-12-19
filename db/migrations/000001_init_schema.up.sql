-- Create User table
CREATE TABLE IF NOT EXISTS User (
    id Int NOT NULL AUTO_INCREMENT,
    fullname varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE, 
    password varchar(255) NOT NULL,
    role enum('admin', 'staff') DEFAULT 'staff' NOT NULL,
    status enum('enabled', 'disabled') DEFAULT 'enabled' NOT NULL,
    PRIMARY KEY (id),
    INDEX (username)
);

-- Create Request table 
CREATE TABLE IF NOT EXISTS Request (
    id varchar(255) NOT NULL,
    status enum('active', 'fulfilled', 'rejected') NOT NULL DEFAULT 'active',
    requestorId Int,
    PRIMARY KEY (id),
    FOREIGN KEY (requestorId) REFERENCES User (id) ON DELETE SET NULL
);

-- Create Progam table
CREATE TABLE IF NOT EXISTS Program (
    id Int NOT NULL AUTO_INCREMENT,
    program varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (program, id)
);

-- Create Major table
CREATE TABLE IF NOT EXISTS Major (
    id Int NOT NULL AUTO_INCREMENT,
    major varchar(255) NOT NULL,
    programId Int NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY (major, programId),
    INDEX (major, id),
    FOREIGN KEY (programId) REFERENCES Program (id) ON DELETE CASCADE
);

-- Create Student table
CREATE TABLE IF NOT EXISTS Student (
    id Int NOT NULL AUTO_INCREMENT,
    controlNumber varchar(255) NOT NULL UNIQUE,
    lastname varchar(255) NOT NULL,
    firstname varchar(255),
    middlename varchar(255),
    type enum('NonTransferee', 'Transferee', 'Graduate') DEFAULT 'NonTransferee' NOT NULL,
    civilStatus enum('single', 'married') DEFAULT 'single' NOT NULL,
    fileLocation varchar(255) NOT NULL,
    programId Int,
    majorId Int,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (controlNumber, lastname),
    FOREIGN KEY (programId) REFERENCES Program (id) ON DELETE SET NULL,
    FOREIGN KEY (majorId) REFERENCES Major (id) ON DELETE SET NULL
);

-- Create Photo table
CREATE TABLE IF NOT EXISTS Photo (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create BirthCertificate table
CREATE TABLE IF NOT EXISTS BirthCertificate (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create Admission table
CREATE TABLE IF NOT EXISTS NoticeOfAdmission (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create Usepat table
CREATE TABLE IF NOT EXISTS Usepat (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create GoodMoral table
CREATE TABLE IF NOT EXISTS GoodMoral (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create Form138 table
CREATE TABLE IF NOT EXISTS Form138 (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create PersonalDataSheet table
CREATE TABLE IF NOT EXISTS PersonalDataSheet (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create DataPrivacyProvision table
CREATE TABLE IF NOT EXISTS DataPrivacyProvision (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create HonorableDismissal table
CREATE TABLE IF NOT EXISTS HonorableDismissal (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create MarriageCertificate table
CREATE TABLE IF NOT EXISTS MarriageCertificate (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create PromissoryNote table
CREATE TABLE IF NOT EXISTS PromissoryNote (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create HealthStateDeclaration table
CREATE TABLE IF NOT EXISTS HealthStateDeclaration (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create MedicalCertificate table
CREATE TABLE IF NOT EXISTS MedicalCertificate (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create Form137 table
CREATE TABLE IF NOT EXISTS Form137 (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create ApprovalSheet table
CREATE TABLE IF NOT EXISTS ApprovalSheet (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create ApplicationForGraduation table
CREATE TABLE IF NOT EXISTS ApplicationForGraduation (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create ShiftersForm table
CREATE TABLE IF NOT EXISTS ShiftersForm (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create CertificateOfLowIncome table
CREATE TABLE IF NOT EXISTS CertificateOfLowIncome (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create TOR table
CREATE TABLE IF NOT EXISTS TOR (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create NMAT table
CREATE TABLE IF NOT EXISTS NMAT (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create Indigency table
CREATE TABLE IF NOT EXISTS Indigency (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create := table
CREATE TABLE IF NOT EXISTS Clearance (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create LeaveOfAbsence table
CREATE TABLE IF NOT EXISTS LeaveOfAbsence (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create AdvancedCreditForm table
CREATE TABLE IF NOT EXISTS AdvancedCreditForm (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create IncompleteForm table
CREATE TABLE IF NOT EXISTS IncompleteForm (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create SubjectValidationForm table
CREATE TABLE IF NOT EXISTS SubjectValidationForm (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create Substitution table
CREATE TABLE IF NOT EXISTS Substitution (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create AffidavitOfUndertaking table
CREATE TABLE IF NOT EXISTS AffidavitOfUndertaking (
    id Int NOT NULL AUTO_INCREMENT,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber)
);

-- Create Other table
CREATE TABLE IF NOT EXISTS Other (
    id Int NOT NULL AUTO_INCREMENT,
    filename varchar(255) NOT NULL,
    location varchar(255) NOT NULL,
    owner varchar(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX (owner),
    FOREIGN KEY (owner) REFERENCES Student (controlNumber),
    UNIQUE KEY (filename, owner)
);

