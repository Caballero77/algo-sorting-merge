package main

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/merge", func(ctx iris.Context) {
		ctx.Write(parseAndSort([]byte("[" + ctx.URLParam("array") + "]")))
	})
	app.Post("/merge", func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		ctx.Write(parseAndSort(body))
	})
	app.Run(iris.Addr(":80"))
}

func parseAndSort(bytes []byte) []byte {
	var array []int
	json.Unmarshal(bytes, &array)

	b, _ := json.Marshal(map[string][]int{"result": sort(array)})

	return b
}

func innerSort(left []int, right []int) []int {
	i := 0
	j := 0
	res := make([]int, len(left)+len(right))
	for i < len(left) || j < len(right) {
		if len(right) == j || (i != len(left) && left[i] < right[j]) {
			res[i+j] = left[i]
			i++
		} else {
			res[i+j] = right[j]
			j++
		}
	}
	return res
}

func sort(list []int) []int {
	length := len(list)
	if length <= 1 {
		return list
	}
	return innerSort(sort(list[0:length/2]), sort(list[length/2:length]))
}
