package urls

type Base string

const (
	pagcloud Base = ""
	prod          = "prod"
	qa            = "qa"
	dev           = "dev"
)

func (base *Base) String() string {

	switch *base {
	case "service":
		return ""
	case "prod":
		return ""
	case "qa":
		return ""
	case "dev":
		return ""
	}

	return ""
}
