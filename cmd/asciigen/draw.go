package main

const (
	scaleX = 8
	scaleY = 3
)

func Draw(
	layouts map[string]map[string][]QMKKeyDictionary,
	keymap QMKKeymap,
) string {
	result := ""

	for layout := range layouts {
		if keymap.Layout != layout {
			continue
		}

		keys, ok := layouts[layout]["layout"]
		if !ok {
			continue
		}

		// Preprocess keys
		// Y aka row -> X aka col
		var maxX4, maxY4 int
		for i := range keys {
			if int(keys[i].W) == 0 {
				keys[i].W = 1
			}

			if int(keys[i].H) == 0 {
				keys[i].H = 1
			}

			// TODO: Better way to handle this
			// Because 0.25
			keys[i].NewX = int(keys[i].X * scaleX)
			keys[i].NewY = int(keys[i].Y * scaleY)
			keys[i].NewW = int(keys[i].W * scaleX)
			keys[i].NewH = int(keys[i].H * scaleY)

			if keys[i].NewX+keys[i].NewW > maxX4 {
				maxX4 = keys[i].NewX + keys[i].NewW
			}

			if keys[i].NewY+keys[i].NewH > maxY4 {
				maxY4 = keys[i].NewY + keys[i].NewH
			}
		}

		for _, layer := range keymap.Layers {
			// Preprocess table
			table := make([][]string, maxY4)
			for i := 0; i < maxY4; i++ {
				table[i] = make([]string, maxX4)
			}

			// Fill layout
			count := 0
			for _, key := range keys {
				keyStr := layer[count]
				keyWidthLimit := key.NewW - 2

				// TODO: Better way to handle this
				if len(keyStr) > keyWidthLimit {
					keyStr = keyStr[:keyWidthLimit]
				} else {
					keyWidthLimit = len(keyStr)
				}

				for i := key.NewY; i < key.NewY+key.NewH; i++ {
					for j := key.NewX; j < key.NewX+key.NewW; j++ {
						if i == key.NewY {
							if j == key.NewX {
								table[i][j] = "+"
							} else {
								table[i][j] = "-"
							}
						} else if i == key.NewY+key.NewH/2 {
							// Write key in the middle
							if j == key.NewX {
								table[i][j] = "|"
							} else if j < key.NewX+keyWidthLimit+1 {
								table[i][j] = string(keyStr[j-key.NewX-1])
							} else {
								table[i][j] = " "
							}
						} else {
							table[i][j] = " "
						}
					}
				}

				count++
			}

			// Postprecess table
			newTable := make([][]string, 0, maxY4+1)

			// Remove empty row
			for i := 0; i < len(table); i++ {
				isEmptyRow := true
				for j := 0; j < len(table[i]); j++ {
					if table[i][j] != " " {
						isEmptyRow = false
						break
					}
				}

				if isEmptyRow {
					continue
				}

				// Padding
				paddingRight := "|"
				if table[i][len(table[i])-1] == "-" {
					paddingRight = "+"
				}

				newTable = append(newTable, append(table[i], paddingRight))
			}

			// Padding
			paddingRow := make([]string, 0, maxX4)
			for j := 0; j < maxX4; j++ {
				paddingBottom := "-"
				if newTable[len(newTable)-1][j] == "|" {
					paddingBottom = "+"
				}
				paddingRow = append(paddingRow, paddingBottom)
			}
			newTable = append(newTable, append(paddingRow, "+"))

			str := ""
			for i := range newTable {
				for j := range newTable[i] {
					str += newTable[i][j]
				}
				str += "\n"
			}

			result += str + "\n"
		}

	}

	return result
}
