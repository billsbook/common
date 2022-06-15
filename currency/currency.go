package currency

// ISO 4217
type Currency struct {
	Country  string
	Currency string
	Code     string
	Number   uint16
}

var (
	Rial = Currency{"IRAN", "Iranian Rial", "IRR", 364}
)
