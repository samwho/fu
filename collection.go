package fu

type Collection struct {
	is []interface{}
}

func Ints(in []int) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}

func Int32s(in []int32) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}

func Int64s(in []int64) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}

func Uints(in []uint) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}

func Uint32s(in []uint32) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}

func Uint64s(in []uint64) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}

func Float32s(in []float32) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}

func Float64s(in []float64) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}

func Strings(in []string) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{is}
}
