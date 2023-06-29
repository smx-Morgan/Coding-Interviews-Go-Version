package src

func replaceSpace(s string) string {
	tmp := ""
	for _, y := range s {
		if string(y) == " " {
			tmp += "%20"
		} else {
			tmp += string(y)
		}
	}
	return tmp
}
