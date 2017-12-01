//Author: Kmila Michel g00340498
//Problem sheet 3 : Problem set: Regular expressions
package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func ElizaResponse(input string) string {
	//\b word boundry
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father"
	}
	if matched, _ := regexp.MatchString(`(.*) ?remember ?(.*)`, input); matched {
		return "Can you tell me more ?"
	}
	if matched, _ := regexp.MatchString(`(.*) ?Hello?(.*)`, input); matched {
		return "Hello it's nice to meet you"
	}
	if matched, _ := regexp.MatchString(`(.*) ?sorry?(.*)`, input); matched {
		return "Why should I be?"
	}

	re := regexp.MustCompile(`(?i)I am ([^.?!].*)[.?!]?`)
	//re := regexp.MustCompile(`(?i).*\bI'?\s*a?m \b([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {

		return re.ReplaceAllString(input, "How do you know you are $1?")
	}

	//Array of random answers
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	//Return random string from answers array
	return answers[rand.Intn(len(answers))]
}

func main() {

	//Prints random input and output from eliza that matched string
	rand.Seed(time.Now().Unix())
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()
	fmt.Println("Hello.")
	fmt.Println(ElizaResponse("Hello."))
	fmt.Println()
	fmt.Println("Are you a computer?")
	fmt.Println(ElizaResponse("Are you a computer?"))
	fmt.Println()
	fmt.Println("Are you sorry?")
	fmt.Println(ElizaResponse("Are you sorry?"))
	fmt.Println()
	fmt.Println("I remember when I was young.")
	fmt.Println(ElizaResponse("I remember when I was young."))
	fmt.Println()
	fmt.Println("Father was a teacher")
	fmt.Println(ElizaResponse("Father was a teacher"))
	fmt.Println()
	fmt.Println("I was my father’s favourite")
	fmt.Println(ElizaResponse("I was my father’s favourite!"))
	fmt.Println()
	fmt.Println("I like my friend")
	fmt.Println(ElizaResponse("I like my friend!"))
	fmt.Println()
	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()
	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()
	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println()
	fmt.Println("I am not happy with your responses")
	fmt.Println(ElizaResponse("I am not happy with your responses"))
	fmt.Println()
	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println()
}
