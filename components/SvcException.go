package components

var resCode = map[int]string{
	101: "LOGIN FAILED",
	102: "INVALID PARAM",
}

func Eksepsi(c int) string {
	return resCode[c]
}
