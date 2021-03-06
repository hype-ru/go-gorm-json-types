package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// ArrayString - массив строк
type ArrayString []string

// Scan - формирование массива
func (s *ArrayString) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

// Value - формирование байтового массива
func (s ArrayString) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
