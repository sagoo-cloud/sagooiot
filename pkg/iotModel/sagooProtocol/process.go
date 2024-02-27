package sagooProtocol

type (
	ProProgress struct {
		Id     string            `json:"id"`
		Params ProProgressParams `json:"params"`
	}

	ProProgressParams struct {
		Step   string `json:"step"`
		Desc   string `json:"desc"`
		Module string `json:"module"`
	}
)
