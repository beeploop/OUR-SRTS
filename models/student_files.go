package models

import "database/sql"

type StudentFiles struct {
	Lastname                 string           `db:"lastname"`
	Firstname                string           `db:"firstname"`
	Middlename               string           `db:"middlename"`
	ControlNumber            string           `db:"controlNumber"`
	Program                  sql.NullString   `db:"program"`
	Major                    sql.NullString   `db:"major"`
	File                     string           `db:"fileLocation"`
	Type                     string           `db:"type"`
	CivilStatus              string           `db:"civilStatus"`
	Photo                    sql.NullString   `db:"pic"`
	BirthCertificate         sql.NullString   `db:"birthCertificate"`
	TOR                      sql.NullString   `db:"tor"`
	NoticeOfAdmission        sql.NullString   `db:"noap"`
	Usepat                   sql.NullString   `db:"usepat"`
	GoodMoral                sql.NullString   `db:"goodMoral"`
	Form138                  sql.NullString   `db:"form138"`
	PersonalDataSheet        sql.NullString   `db:"pds"`
	DataPrivacyProvision     sql.NullString   `db:"dpp"`
	HonorableDismissal       sql.NullString   `db:"hd"`
	MarriageCertificate      sql.NullString   `db:"mc"`
	PromissoryNote           sql.NullString   `db:"pn"`
	HealthStateDeclaration   sql.NullString   `db:"hsd"`
	MedicalCertificate       sql.NullString   `db:"med"`
	Form137                  sql.NullString   `db:"form137"`
	ApprovalSheet            sql.NullString   `db:"approval"`
	ApplicationForGraduation sql.NullString   `db:"afg"`
	ShiftersForm             sql.NullString   `db:"sf"`
	CertificateOfLowIncome   sql.NullString   `db:"col"`
	NMAT                     sql.NullString   `db:"nmat"`
	Indigency                sql.NullString   `db:"indigency"`
	Clearance                []sql.NullString `db:"clearance"`
	AffidavitOfUndertaking   []sql.NullString `db:"aou"`
	LeaveOfAbsence           []sql.NullString `db:"loa"`
	AdvancedCreditForm       []sql.NullString `db:"acf"`
	IncompleteForm           []sql.NullString `db:"inc"`
	SubjectValidationForm    []sql.NullString `db:"svf"`
	Substitution             []sql.NullString `db:"sub"`
	Other                    []Other          `db:"other"`
}

type Other struct {
	Filename string `db:"filename"`
	Location string `db:"location"`
}
