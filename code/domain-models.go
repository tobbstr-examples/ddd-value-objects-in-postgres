package code

// Entity
// Should only be instantiated using its factory.
// The state of an entity can only be changed using its methods.
type Person struct {
	name         Name // fields are unexported to protect the entity's state
	personalInfo PersonalInfo
}

func (p *Person) Name() Name {
	return p.name
}

// PersonalInfo returns a copy of its personalInfo to protect it from
// state changes made outside of the entity.
func (p *Person) PersonalInfo() PersonalInfo {
	return p.personalInfo.Copy()
}

// Value object
type Name string

// Value object
type PersonalInfo struct {
	selfieURLs []string
}

// Copy creates a deep copy of a Person's personal info and returns it.
// It exists to protect the entity's state.
func (i PersonalInfo) Copy() PersonalInfo {
	selfieURLs := make([]string, 0, len(i.selfieURLs))
	copy(selfieURLs, i.selfieURLs)

	return PersonalInfo{selfieURLs: selfieURLs}
}
