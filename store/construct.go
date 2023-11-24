package store

import "log"

func InitializeTables() {
	err := UserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = ProgramTable()
	if err != nil {
		log.Fatal(err)
	}

	err = MajorTable()
	if err != nil {
		log.Fatal(err)
	}

	err = StudentTable()
	if err != nil {
		log.Fatal(err)
	}

	err = RequestTable()
	if err != nil {
		log.Fatal(err)
	}

	err = BirthCertificateTable()
	if err != nil {
		log.Fatal(err)
	}

	err = ClearanceTable()
	if err != nil {
		log.Fatal(err)
	}

	err = TorTable()
	if err != nil {
		log.Fatal(err)
	}

	err = PhotoTable()
	if err != nil {
		log.Fatal(err)
	}

	err = NoticeOfAdmissionTable()
	if err != nil {
		log.Fatal(err)
	}

	err = UsepatTable()
	if err != nil {
		log.Fatal(err)
	}

	err = GoodMoralTable()
	if err != nil {
		log.Fatal(err)
	}

	err = Form138Table()
	if err != nil {
		log.Fatal(err)
	}

	err = PersonalDataSheetTable()
	if err != nil {
		log.Fatal(err)
	}

	err = DataPrivacyProvisionTable()
	if err != nil {
		log.Fatal(err)
	}

	err = HonorableDismissalTable()
	if err != nil {
		log.Fatal(err)
	}

	err = MarriageCertificateTable()
	if err != nil {
		log.Fatal(err)
	}

	err = PromissoryNoteTable()
	if err != nil {
		log.Fatal(err)
	}

	err = HealthStateDeclarationTable()
	if err != nil {
		log.Fatal(err)
	}

	err = MedicalCertificateTable()
	if err != nil {
		log.Fatal(err)
	}

	err = Form137Table()
	if err != nil {
		log.Fatal(err)
	}

	err = ApprovalSheetTable()
	if err != nil {
		log.Fatal(err)
	}

	err = ApplicationForGraduationTable()
	if err != nil {
		log.Fatal(err)
	}

	err = LeaveOfAbsenceTable()
	if err != nil {
		log.Fatal(err)
	}

	err = AffidavitOfUndertakingTable()
	if err != nil {
		log.Fatal(err)
	}

	err = AdvanceCreditFormTable()
	if err != nil {
		log.Fatal(err)
	}

	err = IncompleteFormTable()
	if err != nil {
		log.Fatal(err)
	}

	err = ShiftersFormTable()
	if err != nil {
		log.Fatal(err)
	}

	err = CertificateOfLowIncomeTable()
	if err != nil {
		log.Fatal(err)
	}

	err = SubjectValidationFormTable()
	if err != nil {
		log.Fatal(err)
	}

	err = SubstitutionTable()
	if err != nil {
		log.Fatal(err)
	}

	err = OthersTable()
	if err != nil {
		log.Fatal(err)
	}
}

func UserTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS User (
            id Int NOT NULL AUTO_INCREMENT,
            fullname varchar(255) NOT NULL,
            username varchar(255) NOT NULL, 
            password varchar(255) NOT NULL,
            role enum('admin', 'staff') NOT NULL,
            status enum('active', 'disabled') NOT NULL,
            PRIMARY KEY (id),
            INDEX (username)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func StudentTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Student (
            id Int NOT NULL AUTO_INCREMENT,
            controlNumber varchar(255) NOT NULL UNIQUE,
            lastname varchar(255) NOT NULL,
            firstname varchar(255) NOT NULL,
            middlename varchar(255) NOT NULL,
            type enum('NonTransferee', 'Transferee', 'Graduate') DEFAULT 'NonTransferee' NOT NULL,
            civilStatus enum('single', 'married') DEFAULT 'single' NOT NULL,
            fileLocation varchar(255) NOT NULL,
            programId Int,
            majorId Int,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            INDEX (controlNumber, lastname),
            FOREIGN KEY (programId) REFERENCES Program (id),
            FOREIGN KEY (majorId) REFERENCES Major (id)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func ProgramTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Program (
            id Int NOT NULL AUTO_INCREMENT,
            program varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            PRIMARY KEY (id)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func MajorTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Major (
            id Int NOT NULL AUTO_INCREMENT,
            major varchar(255) NOT NULL UNIQUE,
            programId Int NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (programId) REFERENCES Program (id)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func RequestTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Request (
            id Int NOT NULL AUTO_INCREMENT,
            status enum('active', 'fulfilled', 'rejected') NOT NULL DEFAULT 'active',
            requestorId Int NOT NULL,
            PRIMARY KEY (id),
            FOREIGN KEY (requestorId) REFERENCES User (id)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// SINGLE ENTRY FILES
func PhotoTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Photo (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func BirthCertificateTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS BirthCertificate (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func NoticeOfAdmissionTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS NoticeOfAdmission (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func UsepatTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Usepat (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func GoodMoralTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS GoodMoral (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func Form138Table() error {
	query := `
        CREATE TABLE IF NOT EXISTS Form138 (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func PersonalDataSheetTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS PersonalDataSheet (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func DataPrivacyProvisionTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS DataPrivacyProvision (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func HonorableDismissalTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS HonorableDismissal (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func MarriageCertificateTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS MarriageCertificate (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func PromissoryNoteTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS PromissoryNote (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func HealthStateDeclarationTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS HealthStateDeclaration (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func MedicalCertificateTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS MedicalCertificate (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func Form137Table() error {
	query := `
        CREATE TABLE IF NOT EXISTS Form137 (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func ApprovalSheetTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS ApprovalSheet (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func ApplicationForGraduationTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS ApplicationForGraduation (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func ShiftersFormTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS ShiftersForm (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CertificateOfLowIncomeTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS CertificateOfLowIncome (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func TorTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS TOR (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL UNIQUE,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// TABLE WITH MULTIPLE ENTRIES
func ClearanceTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Clearance (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func LeaveOfAbsenceTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS LeaveOfAbsence (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func AdvanceCreditFormTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS AdvancedCreditForm (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func IncompleteFormTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS IncompleteForm (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func SubjectValidationFormTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS SubjectValidationForm (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func SubstitutionTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Substitution (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func AffidavitOfUndertakingTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS AffidavitOfUndertaking (
            id Int NOT NULL AUTO_INCREMENT,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func OthersTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Other (
            id Int NOT NULL AUTO_INCREMENT,
            filename varchar(255) NOT NULL UNIQUE,
            location varchar(255) NOT NULL,
            owner varchar(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            PRIMARY KEY (id),
            FOREIGN KEY (owner) REFERENCES Student (controlNumber)
        )
    `

	_, err := Db_Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
