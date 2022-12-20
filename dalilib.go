package lua

var daliFuncs = map[string]LGFunction{
	"on":   daliOn,
}

func daliOn(L *LState) *lFile {
	addr := L.CheckInt(1)
	println("using DALI addr", addr)
	return nil
}