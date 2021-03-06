# Go Learning Series

Go lang tutorials and some examples

- All the practices are seperated into different directories
- At the end of learning series, a mini-api created using `json-server` and some function coded into **project** directory to send request and get response from the rest service

# Usage

Only execute `go run main.go` command after the installation of go language runtime, dependencies and executable binaries \
Add go module reference with `go mod init golesson` \
To list all the added modules in the project root directory use `go list -m all`

# Dependencies

`json-server` is used to create rest-api and its endpoints quickly for testing purposes \
To install `json-server` use `npm i -g json-server` \
First of all create a fake database into `[filename].json`. (please click [here](https://github.com/typicode/json-server) to learn more about the json-server) \
To start `json-server` use `json-server --watch [filename].json`
