package main

import (
	"encoding/json"
	"fmt"
)

func jsonCountInt(d interface{}, skipRed bool) float64 {
	switch d := d.(type) {
	case []any:
		var total float64
		for _, i := range d {
			total += jsonCountInt(i, skipRed)
		}
		return total
	case map[string]any:
		var total float64
		for _, v := range d {
			if skipRed && v == "red" {
				return 0
			}
			total += jsonCountInt(v, skipRed)
		}
		return total
	case float64:
		return d
	case string:
		return 0
	default:
		panic(fmt.Sprintf("unexpected type: %T", d))
	}
}

func day12a(input string, skipRed bool) float64 {
	var d any
	if err := json.Unmarshal([]byte(input), &d); err != nil {
		panic(err)
	}
	return jsonCountInt(d, skipRed)
}
