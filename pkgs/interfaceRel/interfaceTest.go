package interfaceRel

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func InterfaceTest1() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}

type ObPhone struct {
	Price float32
	Type  string
}

func (obPhone ObPhone) call() {
	fmt.Printf("type: %s, price: %f \n", obPhone.Type, obPhone.Price)
}

func InterfaceTest2() {
	var phone Phone

	phone = ObPhone{Price: 1000.0, Type: "Nokia"}
	phone.call()

	phone = ObPhone{Price: 2000.0, Type: "Mi11"}
	phone.call()

}

type MiPhone struct {
	ObPhone

	ProductId string
}

func (miPhone *MiPhone) show() {
	miPhone.call()
	fmt.Printf("ProductId: %s\n", miPhone.ProductId)
}

func InterfaceTest3() {
	var phone Phone

	phone = ObPhone{Price: 2000.0, Type: "Mi11"}
	phone.call()

	miPhone := MiPhone{ObPhone{Price: 2000.0, Type: "Mi11"}, "123"}
	miPhone.show()

	miPhone2:=new(MiPhone)
	miPhone2.Price=1000.0
	miPhone2.Type="RedMi"
	miPhone2.ProductId="678"
	miPhone2.show()
}