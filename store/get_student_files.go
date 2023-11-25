package store

import (
	"database/sql"

	"github.com/registrar/models"
)

func GetStudentFiles(controlNumber string) (*models.StudentFiles, error) {
	query := `
        SELECT
            stdnt.lastname as lastname,
            stdnt.firstname as firstname,
            stdnt.middlename as middlename,
            stdnt.controlNumber as controlNumber,
            prog.program as program,
            mjr.major as major,
            stdnt.fileLocation as fileLocation,
            stdnt.type as type,
            stdnt.civilStatus as civilStatus,
            bc.location as birthCertificate,
            noap.location as noap,
            usepat.location as usepat,
            tor.location as tor,
            gm.location as goodMoral,
            form138.location as form138,
            pds.location as pds,
            dpp.location as dpp,
            hd.location as hd,
            mc.location as mc,
            pn.location as pn,
            hsd.location as hsd,
            med.location as med,
            form137.location as form137,
            aprvl.location as approval,
            afg.location as afg,
            sf.location as sf,
            col.location as col
        FROM
            Student stdnt
            LEFT JOIN Program prog on stdnt.programId = prog.id
            LEFT JOIN Major mjr on stdnt.programId = mjr.id
            LEFT JOIN Photo pic on stdnt.controlNumber = pic.owner
            LEFT JOIN BirthCertificate bc on stdnt.controlNumber = bc.owner
            LEFT JOIN NoticeOfAdmission noap on stdnt.controlNumber = noap.owner
            LEFT JOIN Usepat usepat on stdnt.controlNumber = usepat.owner
            LEFT JOIN TOR tor on stdnt.controlNumber = tor.owner
            LEFT JOIN GoodMoral gm on stdnt.controlNumber = gm.owner
            LEFT JOIN Form138 form138 on stdnt.controlNumber = form138.owner
            LEFT JOIN PersonalDataSheet pds on stdnt.controlNumber = pds.owner
            LEFT JOIN DataPrivacyProvision dpp on stdnt.controlNumber = dpp.owner
            LEFT JOIN HonorableDismissal hd on stdnt.controlNumber = hd.owner
            LEFT JOIN MarriageCertificate mc on stdnt.controlNumber = mc.owner
            LEFT JOIN PromissoryNote pn on stdnt.controlNumber = pn.owner
            LEFT JOIN HealthStateDeclaration hsd on stdnt.controlNumber = hsd.owner
            LEFT JOIN MedicalCertificate med on stdnt.controlNumber = med.owner
            LEFT JOIN Form137 form137 on stdnt.controlNumber = form137.owner
            LEFT JOIN ApprovalSheet aprvl on stdnt.controlNumber = aprvl.owner
            LEFT JOIN ApplicationForGraduation afg on stdnt.controlNumber = afg.owner
            LEFT JOIN ShiftersForm sf on stdnt.controlNumber = sf.owner
            LEFT JOIN CertificateOfLowIncome col on stdnt.controlNumber = col.owner
        WHERE stdnt.controlNumber = ?
    `

	var files models.StudentFiles
	err := Db_Conn.Get(&files, query, controlNumber)
	if err != nil {
		return nil, err
	}

	err = getClearance(&files, controlNumber)
	if err != nil {
		return nil, err
	}

	err = getAffidavitOfUndertaking(&files, controlNumber)
	if err != nil {
		return nil, err
	}

	err = getAdvancedCredit(&files, controlNumber)
	if err != nil {
		return nil, err
	}

	err = getSubjectValidation(&files, controlNumber)
	if err != nil {
		return nil, err
	}

	err = getLOA(&files, controlNumber)
	if err != nil {
		return nil, err
	}

	err = getINC(&files, controlNumber)
	if err != nil {
		return nil, err
	}

	err = getSubstitution(&files, controlNumber)
	if err != nil {
		return nil, err
	}

	err = getOthers(&files, controlNumber)
	if err != nil {
		return nil, err
	}

	return &files, nil
}

func getClearance(files *models.StudentFiles, controlNumber string) error {
	query := `
        SELECT location FROM Clearance WHERE owner = ?
    `

	var clearances []sql.NullString
	err := Db_Conn.Select(&clearances, query, controlNumber)
	if err != nil {
		return err
	}

	files.Clearance = clearances

	return nil
}

func getAffidavitOfUndertaking(files *models.StudentFiles, controlNumber string) error {
	query := `
        SELECT location FROM AffidavitOfUndertaking WHERE owner = ?
    `

	var results []sql.NullString
	err := Db_Conn.Select(&results, query, controlNumber)
	if err != nil {
		return err
	}

	files.AffidavitOfUndertaking = results

	return nil
}

func getLOA(files *models.StudentFiles, controlNumber string) error {
	query := `
        SELECT location FROM LeaveOfAbsence WHERE owner = ?
    `

	var results []sql.NullString
	err := Db_Conn.Select(&results, query, controlNumber)
	if err != nil {
		return err
	}

	files.LeaveOfAbsence = results

	return nil
}

func getAdvancedCredit(files *models.StudentFiles, controlNumber string) error {
	query := `
        SELECT location FROM AdvancedCreditForm WHERE owner = ?
    `

	var results []sql.NullString
	err := Db_Conn.Select(&results, query, controlNumber)
	if err != nil {
		return err
	}

	files.AdvancedCreditForm = results

	return nil
}

func getINC(files *models.StudentFiles, controlNumber string) error {
	query := `
        SELECT location FROM IncompleteForm WHERE owner = ?
    `

	var results []sql.NullString
	err := Db_Conn.Select(&results, query, controlNumber)
	if err != nil {
		return err
	}

	files.IncompleteForm = results

	return nil
}

func getSubjectValidation(files *models.StudentFiles, controlNumber string) error {
	query := `
        SELECT location FROM SubjectValidationForm WHERE owner = ?
    `

	var results []sql.NullString
	err := Db_Conn.Select(&results, query, controlNumber)
	if err != nil {
		return err
	}

	files.SubjectValidationForm = results

	return nil
}

func getSubstitution(files *models.StudentFiles, controlNumber string) error {
	query := `
        SELECT location FROM Substitution WHERE owner = ?
    `

	var results []sql.NullString
	err := Db_Conn.Select(&results, query, controlNumber)
	if err != nil {
		return err
	}

	files.Substitution = results

	return nil
}

func getOthers(files *models.StudentFiles, controlNumber string) error {
	query := `
        SELECT filename, location FROM Other WHERE owner = ?
    `

	var results []models.Other
	err := Db_Conn.Select(&results, query, controlNumber)
	if err != nil {
		return err
	}

	files.Other = results

	return nil
}
