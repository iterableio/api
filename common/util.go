package common

func Unzip(m map[string]interface{}) ([]string, []interface{}) {
	keys := make([]string, len(m))
	values := make([]interface{}, len(m))

	i := 0
	for k, v := range m {
		keys[i] = k
		values[i] = v
		i++
	}

	return keys, values
}

func UnzipSlices(m map[string][]interface{}) ([]string, [][]interface{}) {
	keys := make([]string, len(m))
	values := make([][]interface{}, len(m))

	i := 0
	for k, v := range m {
		keys[i] = k
		values[i] = v
		i++
	}

	return keys, values
}
