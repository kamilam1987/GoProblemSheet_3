# GoProblemSheet_3
# What is go (programming language)?
Go (often referred to as golang) is a programming language created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing, memory safety features and CSP-style concurrent programming features added. The compiler and other language tools originally developed by Google are all free and open source.
## How to install go
Please visit this link: https://golang.org/doc/install?download=go1.5.windows-amd64.msi2
## How to clone a repository and run programs
Please visit links: <br />
https://help.github.com/articles/cloning-a-repository<br />
https://github.com/processing/processing/wiki/Build-Instructions
### Instruction 
- Step 1: In command line or terminal type: git clone https://github.com/kamilam1987/GoProblemSheet_3.git
- Step 2: Navigate to : cd GoProblemSheet_3
- Step 3: go build .\ProblemGo3.go
- Step 4: .\ProblemGo3.exe or go run .\ProblemGo3.go

https://github.com/processing/processing/wiki/Build-Instructions
## This repository contains problem solving :
1. Random response

Write a function called ElizaResponse in Go that takes a single string as input and returns a single string as output. The function should randomly return one of the following three strings.

“I’m not sure what you’re trying to say. Could you explain it to me?”<br />
“How does that make you feel?”<br />
“Why do you say that?”<br />
Add a main function to your Go code. Test your ElizaResponse function by calling it from main with the following test inputs.<br />

“People say I look like both my mother and father.”<br />
“Father was a teacher.”<br />
“I was my father’s favourite.”<br />
“I’m looking forward to the weekend.”<br />
“My grandfather was French!”<br />
Have main print both the input to and output from the function to the terminal.<br />

2. Recognise “father”<br />

Adapt ElizaResponse to check, using a regular expression, if the input contains the word “father”. If it does, it should return the string “Why don’t you tell me more about your father?” Otherwise, it should return one of the three random responses as before. You should make sure that your regular expression recognises the first three test inputs, but not the last two.

3. Capture “I am “<br />

Adapt the ElizaResponse function to, if the input does not contain the word “father”, check the input begins with “I am “. If it does, use a regular expression to capture the rest of the user input. Return the string “How do you know you are _?”, replacing the underscore with the captured part of the input. Add, along with the five previous test inputs, the following test inputs.

“I am happy.”<br />
“I am not happy with your responses.”<br />
“I am not sure that you understand the effect that your questions are having on me.”<br />
“I am supposed to just take what you’re saying at face value?”<br />
4. “i am “, “I AM “, “I’m “, “Im “, “i’m “<br />

Adapt the function to respond in the same way as with “I am _”, when the input begins with any reasonable variant of “I am “, such as “I’m” or “Im”.<br />

5. Reflect pronouns<br />

Adapt the function to reflect the pronouns in the captured groups, where necessary. For instance, when the following input is given:

“I am not sure that you understand the effect your questions are having on me.”
the function should return:<br />

“How do you know that you are not sure that I understand the effect my questions are having on you.”<br />
6. More input patterns<br />

Add three of your own input patterns to the function, with corresponding outputs.<br />
