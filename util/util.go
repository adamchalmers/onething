package util

func DieIfErr(e error) {
	if e != nil {
		panic(e)
	}
}
