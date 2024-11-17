package domain

// CurrencyRate stores information about exchange rates
type CurrencyRate struct {
	Base   string
	Target string
	Rate   float64
}
