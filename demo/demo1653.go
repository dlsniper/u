/**
 * @author Florin Patan <florinpatan@gmail.com>
 */

// Package 1653
package main

type (
	funcTiOn func() ([]string, Cmd)

	Cmd interface {
		Ini(keys map[string]string)
	}
)

var (
	funcs = []funcTiOn{}
)

func main() {
	for idx := range funcs {
		var handler funcTiOn
		_, handler, _ = funcs[idx]()
		handler.Ini(map[string]string) // Unresolved reference 'Ini'
	}

	for _, val := range funcs {
		_, handler := val()
		handler.Ini(map[string]string) // Unresolved reference 'Ini'
	}
}
