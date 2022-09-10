package main

type MessageFilter func(string) string

func getWelcomeWording(realMessage string, filter MessageFilter) string {
	return filter(realMessage)
}

func main() {

	welcomeWording1 := getWelcomeWording("John", func(name string) string {
		if name == "John" {
			return "Welcome John"
		} else if name == "Doe" {
			return "Welcome Doe"
		} else {
			return "Who are you?"
		}
	})
	println(welcomeWording1)

	welcomeWording2 := func(name string) string {
		if name == "John" {
			return "Welcome John"
		} else if name == "Doe" {
			return "Welcome Doe"
		} else {
			return "Who are you?"
		}
	}

	println(getWelcomeWording("Doe", welcomeWording2))
}
