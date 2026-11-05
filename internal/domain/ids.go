package domain

import "fmt"

const (
	skiIDPrefix    = "S-"
	personIDPrefix = "P-"
)

func ValidateSkiID(id string) error {
	return validateID(id, skiIDPrefix)
}

func ValidatePersonID(id string) error {
	return validateID(id, personIDPrefix)
}

func validateID(id, prefix string) error {
	if len(id) != len(prefix)+4 || id[:len(prefix)] != prefix {
		return fmt.Errorf("%w: %q", ErrInvalidID, id)
	}
	for _, character := range id[len(prefix):] {
		if character < '0' || character > '9' {
			return fmt.Errorf("%w: %q", ErrInvalidID, id)
		}
	}
	return nil
}

type IDGenerator struct {
	nextSki    uint16
	nextPerson uint16
}

func (generator *IDGenerator) NewSkiID() (string, error) {
	return generator.nextID(skiIDPrefix, &generator.nextSki)
}

func (generator *IDGenerator) NewPersonID() (string, error) {
	return generator.nextID(personIDPrefix, &generator.nextPerson)
}

func (generator *IDGenerator) nextID(prefix string, counter *uint16) (string, error) {
	if *counter >= 10000 {
		return "", fmt.Errorf("%w: ID range exhausted for %s", ErrInvalidID, prefix)
	}
	(*counter)++
	return fmt.Sprintf("%s%04d", prefix, *counter), nil
}
