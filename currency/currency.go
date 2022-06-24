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

// return the currency for the given code
func GetCurrency(code string) Currency {
	switch code {
	case "IRR":
		return Rial
	default:
		return Rial
	}
}
