### Introduction
- Go also known as Golang is an open source, **compiled** and **statically typed** programming language    developed by Google.
- Go is not an object oriented programming language.

### Advantages of Go
- Easy to write concurrent programs with the help of  **Goroutines** and **channels**.
- Static linking
        The Go compiler supports static linking. The entire Go project can be statically linked into one big fat binary and it can be deployed in cloud servers easily without worrying about dependencies.
- Garbage collection
### Running a go program
**1) go install**
```
go install
```
The above command will compile the program and install(copy) the binary to location ~/go/bin. The name of the binary will be the folder containing main.go. In our case, it will be named learngo.
```
export GOBIN=~/go/bin/
```
The above command specifies that go install should copy the compiled binary to the path `~/go/bin/`

Command to run the compiled binary.
```
~/go/bin/Golang        (only if the path is set to this location)
```
If you want to avoid typing the entire path `~/go/bin/Golang each` time you run the program, you can add `~/go/bin/` to your PATH.
```
export PATH=$PATH:~/go/bin
```
Now you can just type `Golang` in the terminal to run the program.

**2) go build**
The second option to run the program is using `go build`. go build is much similar to `go install` except that it doesn't install(copy) the compiled binary to the path `~/go/bin/`, rather it creates the binary inside the location from which go build was installed.
```
go build
```
The above command will create a binary named `Golang` in the current directory

Type `./Golang` to run the program.

**3) go run**
```
go run <FILE_NAME>
```
-   One subtle difference between the `go run` and `go build/go install` commands is,     go run requires the name of the `.go` file as an argument.
-   `go run` works much similar to `go build`. Instead of compiling and installing the     program to the current directory, it compiles the file to a temporary location and     runs the file from that location.
-   To know the location where `go run` compiles the file to, run `go run` with the       `--work` argument.
```
go run --work main.go
```

### go commands
![go_commands.png](https://github.com/96ankur/Golang/blob/master/img/go-commands.png)