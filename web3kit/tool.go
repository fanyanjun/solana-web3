package web3kit

func Recover(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = e
		} else {
			panic(r)
		}
	}
}

func Map[T any, O any, TS ~[]T](input TS, mapper func(int, T) O) []O {
	var output = make([]O, len(input))
	for i, data := range input {
		output[i] = mapper(i, data)
	}
	return output
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
func Must1[T any](arg T, err error) T {
	if err != nil {
		panic(err)
	}
	return arg
}
func Must2[T any, T2 any](arg T, arg2 T2, err error) (T, T2) {
	if err != nil {
		panic(err)
	}
	return arg, arg2
}
func Must3[T any, T2 any, T3 any](arg T, arg2 T2, arg3 T3, err error) (T, T2, T3) {
	if err != nil {
		panic(err)
	}
	return arg, arg2, arg3
}