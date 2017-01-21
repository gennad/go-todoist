package todoist

import (
	"github.com/satori/go.uuid"
	"strconv"
	"fmt"
)

type ID string

func NewID(id string) (ID, error) {
	if _, err := strconv.Atoi(id); err == nil {
		return ID(id), nil
	}
	if IsTempID(ID(id)) {
		return ID(id), nil
	}
	return "", fmt.Errorf("Invalid ID: %s", id)
}

func (i ID) MarshalJSON() ([]byte, error) {
	s := string(i)
	if IsTempID(i) {
		s = `"` + s + `"`
	}
	return []byte(s), nil
}

func (i *ID) UnmarshalJSON(b []byte) (err error) {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)  // integer id
	}
	id, err := NewID(s)
	if err != nil {
		return err
	}
	*i = id
	return nil
}

func GenerateTempID() ID {
	return ID(uuid.NewV4().String())
}

func IsTempID(id ID) bool {
	if _, err := uuid.FromString(string(id)); err == nil {
		return true
	}
	return false
}

type UUID string

func GenerateUUID() UUID {
	return UUID(uuid.NewV4().String())
}

func (i UUID) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(string(i))), nil
}

func (i *UUID) UnmarshalJSON(b []byte) (err error) {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	id, err := uuid.FromString(s)
	if err != nil {
		return err
	}
	*i = UUID(id.String())
	return nil
}