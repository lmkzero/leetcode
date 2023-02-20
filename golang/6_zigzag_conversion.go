package main

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([][]byte, numRows)
	var (
		curRow    int
		goingDown bool
	)
	for _, c := range []byte(s) {
		row := rows[curRow]
		row = append(row, c)
		rows[curRow] = row
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}
	var res string
	for _, row := range rows {
		res += string(row)
	}
	return res
}
