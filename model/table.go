package model

type Table struct {
	rows int
	columns int
	elements []string
}

func (t Table) GetElement(row int, column int) string {
	index := t.calculateIndex(row, column)

	return t.elements[index]
}

func (t Table) SetValue(row int, column int, value string) {
	index := t.calculateIndex(row, column)

	t.elements[index] = value
}

func (t Table) calculateIndex(row int, column int) int {
	if (row <= 0 || row > t.rows) {
		// erro, fora do índice
	}
	if (column <= 0 || column > t.columns) {
		// erro, fora do índice
	}

	return (column + t.columns * (row - 1)) - 1
}

func CreateTable(rows int, columns int) Table {
	if (rows <= 0) {
		// erro, índice inválido
	}
	if (columns <= 0) {
		// erro, índice inválido
	}

	return Table{
		rows: rows,
		columns: columns,
		elements: make([]string, (columns*rows), (columns*rows)),
	}
}
