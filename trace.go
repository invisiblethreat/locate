package locate

import "runtime"

// WhereAmI returns the current package and function as a string
func WhereAmI() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "lost"
	}

	fn := runtime.FuncForPC(pc)
	return fn.Name()
}
