package valueObjects

import "encoding/json"

type StringVO struct {
	Value string
}

func (stringValueObject StringVO) String() string {
	return stringValueObject.Value
}

func (stringValueObject StringVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(stringValueObject.Value)
}

func (stringValueObject *StringVO) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &stringValueObject.Value)
}
