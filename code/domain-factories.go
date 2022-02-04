package code

import (
	"fmt"
	"net/url"
	"strings"
)

type NameFactory struct{}

func (NameFactory) Make(name string) (Name, error) {
	// Factories are responsible for only instantiating valid entities and value objects.

	// They perform validation like the one below
	if len(name) == 0 {
		return "", fmt.Errorf("empty name not allowed")
	}

	if strings.ContainsAny(name, "/%#") {
		return "", fmt.Errorf("IllegalChars")
	}

	return Name(name), nil
}

type PersonalInfoFactory struct{}

func (PersonalInfoFactory) Make(selfieURLs []string) (PersonalInfo, error) {
	// Validates URLs
	for _, selfieURL := range selfieURLs {
		if _, err := url.ParseRequestURI(selfieURL); err != nil {
			return PersonalInfo{}, fmt.Errorf("invalid URL")
		}
	}

	return PersonalInfo{selfieURLs: selfieURLs}, nil
}

type PersonFactory struct{}

func (PersonFactory) Make(name Name, personalInfo PersonalInfo) Person {
	return Person{name: name, personalInfo: personalInfo}
}
