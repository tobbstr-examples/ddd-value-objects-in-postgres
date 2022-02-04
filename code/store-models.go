package code

type PersonRecord struct {
	ID           string
	Name         string
	PersonalInfo PersonalInfoAttr
}

type PersonalInfoAttr struct {
	SelfieURLs []string
}
