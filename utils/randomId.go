package utils

/*
A UUID is designed as a number that is globally unique in space and time.
Two calls to UUID() are expected to generate two different values,
even if these calls are performed on two separate computers that are not connected to each other.
*/

/*
Custom type for UUID for Gorm. Adapted from:
https://gist.github.com/dipesh429/46781db11a11df8dd5a1d1bd062a17d2#file-customized-data-type-go
*/

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

//RandomId -> new datatype
type RandomId uuid.UUID

// StringToRandomId -> parse string to RandomId
func StringToRandomId(s string) (RandomId, error) {
	id, err := uuid.Parse(s)
	return RandomId(id), err
}

//String -> String Representation of Binary16
func (my RandomId) String() string {
	return uuid.UUID(my).String()
}

//GormDataType -> sets type to binary(16)
func (my RandomId) GormDataType() string {
	return "binary(16)"
}

func (my RandomId) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(my)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (my *RandomId) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*my = RandomId(s)
	return err
}

// Scan --> tells GORM how to receive from the database
func (my *RandomId) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = RandomId(parseByte)
	return err
}

// Value -> tells GORM how to save into the database
func (my RandomId) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
