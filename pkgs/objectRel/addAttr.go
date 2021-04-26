package objectRel

import "fmt"

type Item struct {
	Id       string
	Score    float64
	IsLang   bool
	IsNew    bool
	ExpLayer int
}

func AddAttrTest1() {
	item := Item{Id: "a1", Score: 0.5, IsLang: false, IsNew: true, ExpLayer: 3}
	fmt.Printf("item: %+v\n", item)

	item2 := Item{Id: "a2", Score: 0.5, IsLang: false}
	fmt.Printf("item: %+v\n", item2)
}
