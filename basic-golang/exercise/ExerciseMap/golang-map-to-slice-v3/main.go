package main

func MapToSlice(mapData map[string]string) [][]string {
	var res [][]string
	for key, val := range mapData {
		var tmp []string
		tmp = append(tmp, key, val)
		res = append(res, tmp)
	}
	return res // TODO: replace this
}
