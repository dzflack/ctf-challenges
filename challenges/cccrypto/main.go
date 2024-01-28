package main

import (
	"syscall/js"
)

func submit(this js.Value, inputs []js.Value) interface{} {
	println("In golang submit function")
	event := inputs[0]
	callback := inputs[1]

	username := event.Get("target").Call("querySelector", "[name='username']").Get("value").String()
	creditCardNumber := event.Get("target").Call("querySelector", "[name='creditcardnumber']").Get("value").String()

	if username == "796952" && (creditCardNumber == "4971 9660 3200 7972" || creditCardNumber == "4971966032007972") {
		println("returning true")

		callback.Invoke("Flag: CTF{15f5d85dd1fe2c86a6d04f96328b16b3}")
	} else {
		println("returning false")

		callback.Invoke("Nope")
	}

	return nil
}

func registerCallbacks() {
	js.Global().Set("submit", js.FuncOf(submit))
}

func main() {
	c := make(chan struct{})
	println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
