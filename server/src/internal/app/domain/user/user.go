package user

type Model struct {
	Firstname string
	Lastname  string
}

func (m *Model) GetFirstname() string {
	return m.Firstname
}
