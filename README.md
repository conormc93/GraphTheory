### Developed By:
*Conor McGrath*

### Lecturer:
*Dr. Ian McLoughlin*

## This is a go application that can:
1. Build a postfix expression from an infix expression. 
2. Build an NFA from a regular postfix expression.
3. Check to see if a string matches a regular expression.

## Running the program
This program is made using the Go programming language.

Click on the following link to download and install Go. [INSTALL GO](https://golang.org/dl/)

To clone the repository to your local machine, in command prompt enter 
```
git clone https://github.com/conormc93/GraphTheory
```

## How to run: 

1. **Build and Execute** 
Navigate to the folder where you cloned the repository to and enter the following to compile the code 
```
go build g00291461.go
```

Then type the name of that file without any file extension.
```
g00291461
```


2. **Build and Run** 
Navigate to the folder where you cloned the repository to and enter the following to compile the code 
```
go build g00291461.go
```

This will create a .exe file in your current directory also.
You can then type the following to run the program
```
go run g00291461.go
```

#### Description
The aim of this project is to write a program in the Go programming language that can build a non-deterministic finite automaton (NFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text. This program does not use the regexp package from the Go standard library nor any other external library. A regular expression is a string containing a series of characters, some of which may have a special meaning. For example, this program can determine between characters ".", "|", "*", "+" and "?" which have the special meanings "concatenate", "or", "Kleene star", "match-one-or-more quantifier" and "match-zero-or-one quantifier" respectively.

#### Research
As with any project, my first task was to break it down into manageable sections or pieces. I went online and found a few different resources on Regex, and Go. My lecturer had online tutorials on how to go about tackling the project also. Once I had finished my information gathering I looked at various examples online about how others had gone about coding something similar, maybe perhaps in a different language.

#### Design
For the design process I basically looked at resources the lecturer had put up. The methods used and the implementation of the code isn't as unique as I would like but to thoroughly understand the concepts and task of the project I was relying a lot on my lecturers help through his online videos. felt I completed it as best I could at this moment in time with all the knowledge of Go and Regex I have. 

#### Resources
- Dr Ian McLoughlin - Galway - Mayo Institute of Technology
+ (https://stackoverflow.com/)
+ (https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d)
+ (https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b)
+ (https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e)
