package code

// PersonRecord is a data structure that represents a record in the Person table
type PersonRecord struct {
	ID           string
	Name         string
	PersonalInfo PersonalInfoAttr
}

type PersonalInfoAttr struct {
	SelfieURLs []string
}
