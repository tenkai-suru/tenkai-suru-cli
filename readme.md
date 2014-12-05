# Tenkai-Suru
### Running the app
After you have the `ts` file, type `./ts --help` to see the all commands you can run.

### To Deploy
`./ts deploy [projectname]`

or 

`./ts d [projectname]`

### Setup for Mac
1. Install Go `brew install go`
2. [Set up your workspace](https://golang.org/doc/code.html)
3. Navigate into your workspace (ex. `cd $GOPATH/src`)
4. Download the project `git clone https://github.com/tenkai-suru/tenkai-suru-cli.git`
5. Navigate into the project `cd tenkai-suru-cli`
6. Get the dependency `go get github.com/codegangsta/cli`
7. Build the app `go build -o ts`
8. Run the app `./tenkai-suru-cli`
