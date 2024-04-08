package test

import (
	"encoding/json"
	"strconv"
)

type Visitor struct {
	ID      int      `json:"id"`
	Name    string   `json:"name,omitempty"`
	Phones  []string `json:"phones,omitempty"`
	Company string   `json:"company,omitempty"`
}

// MarshalJSON реализует интерфейс json.Marshaler.
func (v Visitor) MarshalJSON() ([]byte, error) {
	// чтобы избежать рекурсии при json.Marshal, объявляем новый тип
	type VisitorAlias Visitor

	aliasValue := struct {
		VisitorAlias
		// переопределяем поле внутри анонимной структуры
		ID string `json:"id"`
	}{
		// встраиваем значение всех полей изначального объекта (embedding)
		VisitorAlias: VisitorAlias(v),
		// задаём значение для переопределённого поля
		ID: strconv.Itoa(v.ID),
	}

	return json.Marshal(aliasValue) // вызываем стандартный Marshal
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler.
func (v *Visitor) UnmarshalJSON(data []byte) (err error) {
	// чтобы избежать рекурсии при json.Unmarshal, объявляем новый тип
	type VisitorAlias Visitor

	aliasValue := &struct {
		*VisitorAlias
		// переопределяем поле внутри анонимной структуры
		ID string `json:"id"`
	}{
		VisitorAlias: (*VisitorAlias)(v),
	}
	// вызываем стандартный Unmarshal
	if err = json.Unmarshal(data, aliasValue); err != nil {
		return
	}
	v.ID, err = strconv.Atoi(aliasValue.ID)
	return
}
