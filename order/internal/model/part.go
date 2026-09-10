package model

type Part struct {
	UUID  string
	Price float64
}

type PartsFilter struct {
	UUIDs []string
}
