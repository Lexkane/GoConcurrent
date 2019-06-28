package main

import "fmt"

func main() {

	button := MakeButton()
	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	button.AddEventListiner("click", handlerOne)
	button.AddEventListiner("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler One" + msg)
		}
	}()
	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler Two", msg)
		}
	}()

	button.TriggerEvent("click", "Button clicked")
	button.RemoveEventListiner("click", handlerTwo)
	button.TriggerEvent("click", "Button clicked again")

	fmt.Scanln()

}

type Button struct {
	eventListiners map[string][]chan string
}

func MakeButton() *Button {
	result := new(Button)
	result.eventListiners = make(map[string][]chan string)
	return result
}

func (this *Button) AddEventListiner(event string, responseChannel chan string) {
	if _, present := this.eventListiners[event]; present {
		this.eventListiners[event] = append(this.eventListiners[event], responseChannel)
	} else {
		this.eventListiners[event] = []chan string{responseChannel}
	}
}

func (this *Button) RemoveEventListiner(event string, listinerChannel chan string) {
	if _, present := this.eventListiners[event]; present {
		for idx, _ := range this.eventListiners[event] {
			if this.eventListiners[event][idx] == listinerChannel {
				this.eventListiners[event] = append(this.eventListiners[event][:idx],
					this.eventListiners[event][idx+1:]...)
				break
			}
		}
	}
}

func (this *Button) TriggerEvent(event string, response string) {
	if _, present := this.eventListiners[event]; present {
		for _, handler := range this.eventListiners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
