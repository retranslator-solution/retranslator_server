package models

type StringResource struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type ArrayValue struct {
	Value string `json:"value" binding:"required"`
}

type ArrayResource struct {
	Name   string       `json:"name" binding:"required"`
	Values []ArrayValue `json:"values" binding:"required"`
}

type PluralValue struct {
	Value    string `json:"value" binding:"required"`
	Quantity string `json:"quantity" binding:"required"`
}

type PluralResource struct {
	Name   string        `json:"name" binding:"required"`
	Values []PluralValue `json:"values" binding:"required"`
}

type Resource struct {
	Name            string           `json:"name" binding:"required"` // ru_RU, en_EN etc
	ArrayResources  []ArrayResource  `json:"arrayResources" binding:"required"`
	StringResources []StringResource `json:"stringResources" binding:"required"`
	PluralResources []PluralResource `json:"pluralResources" binding:"required"`
}
