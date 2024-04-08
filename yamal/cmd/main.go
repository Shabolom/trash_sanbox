package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: Gopher
values:
- 11
- 22
- 33
`

func main() {
	// памятка о преобразовании типов в байты
	var v Data
	gog := 123459485
	gPo := "213"
	fgf := make([]int, 1)
	fgf = append(fgf, 1)
	byteGpo := []byte(gPo)
	byteFgf := make([]byte, len(fgf))
	byteGog := []byte{byte(gog)}

	fmt.Println(byteGpo)
	fmt.Println(byteFgf)
	fmt.Println(byteGog)

	err := yaml.Unmarshal([]byte(yamlData), &v)
	if err != nil {
		panic(err)
	}
	out, err := toml.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	fmt.Println(v)

	// вставьте недостающий код
	// 1) десериализуйте yamlData в переменную типа Data
	// 2) преобразуйте полученную переменную в TOML
	// 3) выведите в консоль результат

}
