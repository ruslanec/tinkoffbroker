package domain

// Данные о стране.
type Country struct {
	AlfaTwo   string `json:"alfa_two,omitempty"`   // Двухбуквенный код страны.
	AlfaThree string `json:"alfa_three,omitempty"` // Трёхбуквенный код страны.
	Name      string `json:"name,omitempty"`       // Наименование страны.
	NameBrief string `json:"name_brief,omitempty"` // Краткое наименование страны.
}
