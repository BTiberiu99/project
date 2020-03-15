package problem

import (
	"bufio"
	"io"
	"strings"
)

type MaC struct {
}

//First should be the intial matrix
//Last should be final matrix
//Between them should be an empty row
func (m *MaC) FromFile(file io.Reader) error {

	scanner := bufio.NewScanner(file)
	matrixs := make([]string, 2)

	//Split matrixes
	str := ""

	for scanner.Scan() {

		row := strings.TrimSpace(scanner.Text())

		if scanner.Err() != nil {
			return scanner.Err()
		}

		if row == "" {
			matrixs[0] = str
			str = ""

		} else {
			str += row + "\n"
		}

	}

	matrixs[1] = str

	//End Split

	return nil
}
