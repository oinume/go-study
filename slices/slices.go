package main
import "fmt"

type QueryParams struct {
	Fields []string
}

func main() {
	fmt.Println("--- Simple range ---")
	var accounts []string = []string{ "oinume", "oranie" }
	for _, value := range accounts {
		fmt.Println(value)
	}

	fmt.Println("--- Params ---")

	var params *Params = &Params{Key:"hoge", Fields: []string{ "oinume", "oranie" }}
	for _, value := range params.Fields {
		fmt.Println(value)
	}
}
