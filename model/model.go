package model

type Model struct {
	id         string
	attribute1 string
}

func (m Model) Id() string {
	return m.id
}

func (m *Model) SetId(id string) {
	m.id = id
}

func (m Model) Attribute1() string {
	return m.attribute1
}

func (m *Model) SetAttribute1(attribute1 string) {
	m.attribute1 = attribute1
}
