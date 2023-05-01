package valueObjects

import (
	"encoding/json"
)

type StringVO struct {
	value string
}

func CreateStringVO(value string) StringVO {
	return StringVO{value: value}
}

func (stringValueObject StringVO) String() string {
	return stringValueObject.value
}

func (stringValueObject StringVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(stringValueObject.value)
}

func (stringValueObject *StringVO) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &stringValueObject.value)
}
