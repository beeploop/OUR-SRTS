package utils

func StringInSlice(a string) bool {
	multipleEntries := []string{
		"Clearance",
		"AffidavitOfUndertaking",
		"LeaveOfAbsence",
		"AdvancedCreditForm",
		"IncompleteForm",
		"SubjectValidationForm",
		"Substitution",
	}

	for _, b := range multipleEntries {
		if b == a {
			return true
		}
	}

	return false
}
