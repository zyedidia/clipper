package clipper

import "bytes"

type Multi struct {
	Clip Clipboard

	data map[string][][]byte
}

func (m *Multi) Init() {
	m.data = map[string][][]byte{
		"clipboard": [][]byte{},
		"primary":   [][]byte{},
	}
}

func (m *Multi) full(reg string) []byte {
	content := m.data[reg]
	buf := &bytes.Buffer{}
	for _, s := range content {
		buf.Write(s)
	}
	return buf.Bytes()
}

func (m *Multi) text(reg string, num int) []byte {
	content := m.data[reg]
	if len(content) <= num {
		return nil
	}
	return content[num]
}

func (m *Multi) ReadCursor(reg string, num, ncursors int) ([]byte, error) {
	clip, err := m.Clip.ReadAll(reg)
	if err != nil {
		return nil, err
	}
	if m.IsValid(reg, clip, ncursors) {
		return m.text(reg, num), nil
	}
	return clip, nil
}

func (m *Multi) WriteCursor(reg string, p []byte, num, ncursors int) error {
	if len(m.data[reg]) != ncursors {
		m.data[reg] = make([][]byte, ncursors)
	}
	if num >= ncursors {
		return nil
	}
	m.data[reg][num] = p
	return m.Clip.WriteAll(reg, m.full(reg))
}

func (m *Multi) IsValid(reg string, system []byte, ncursors int) bool {
	if len(m.data[reg]) != ncursors {
		return false
	}
	return bytes.Equal(system, m.full(reg))
}
