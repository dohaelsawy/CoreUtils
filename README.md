# CoreUtils
## Description
Goal is to reimplement several core Unix utilities. giving exposure to Go basics, including file I/O, flag parsing, string manipulation, and more.
## 1. head
* Print the first 10 lines of input by default.
* Add a ``-n`` flag to specify the number of lines to print.
``` 
go build head.go
```
``` 
./head -n 4 \input.txt
```
## 2. tail
* Print the last 10 lines of input by default.
* Add a ``-n`` flag to specify the number of lines to print.
``` 
go build tail.go
```
``` 
./tail -n 4 \input.txt
```
## 3. wc (word count)
* Count lines, words, and characters in the input.
* Add ``-l``, ``-w``, and ``-c`` flags to display only lines, words, or characters respectively.
``` 
go build wc.go
```
``` 
./wc -l -w -c \input.txt
```
## 5. cat
* Concatenate and print files.
* Add a ``-n`` flag to number output lines.
``` 
go build echo.go
```
``` 
./echo "hello world"
```
## 5. echo
* Print arguments to standard output.
* Add a -n flag to omit the trailing newline.
``` 
go build echo.go
```
``` 
./echo -n "hello world"
```
## 6. yes
*  It outputs a string repeatedly until terminated.
``` 
go build yes.go
```
``` 
./yes hello world
```
## 7. true
* return always successful responce
``` 
go build true.go
```
``` 
./true 
```
## 8. false
* return always failed responce
``` 
go build false.go
```
``` 
./false 
```
## 9. env
* prints out all the environment variables of the current process. 
``` 
go build env.go
```
``` 
./env 
```
## 10. tree
* recursive directory listing program that produces a depth-indented listing of files and directories.
* Add a ``-L`` flag to specify the number of depth level of Dictionary 
``` 
go build tree.go
```
``` 
./tree -L 4 \Dir
```

 
