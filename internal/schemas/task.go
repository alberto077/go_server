package schemas

type TaskIn struct {
	Date            string `json:"date"`
	Name            string `json:"name"`
	ItemDescription string `json:"item"`
	Priority        int32  `json:"priority"`
}

// Have a validator for the time
type TaskOut struct {
}
