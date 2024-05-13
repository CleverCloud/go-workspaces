# Deploy a Go Workspace based app on Clever Cloud
[In its version 1.18](https://go.dev/blog/get-familiar-with-workspaces), Go introduced Workspaces to use multiple modules in the same project easier. This repository contains an application (in the `cmd/` folder) relying on two Go Modules declared in the `go.work` file. To try by yourself, you'll need a [Clever Cloud account](https://console.clever-cloud.com/users/me/information) and our CLI, [Clever Tools](https://github.com/CleverCloud/clever-tools) (you can also use our [Console](https://console.clever-cloud.com)). 

The easier way to install it is through `npm` (but there are lots of [other ways](https://github.com/CleverCloud/clever-tools/blob/master/docs/setup-systems.md)):
```bash
npm i -g clever-tools
clever login
clever profile
```

Then `git clone` this repository and follow our deploy instructions: 
```bash
git clone https://github.com/CleverCloud/go-workspaces.git
cd go-workspaces
```

## How to deploy on Clever Cloud - Makefile
This repository contains a [`Makefile`](Makefile). So you just have to create a Go application on your account, declare where is the built binary through `CC_GO_BINARY` environment variable and `git push`!
```bash
clever create -t go
clever env set CC_GO_BINARY bin/myApp
git add . && git commit -m "First commit"
clever deploy && clever domain
```
After that, your application should be built and deployed. Just type its URL in your browser!

## How to deploy on Clever Cloud - Env vars

If you prefer to rely on environment variables, delete `Makefile`. The create process is this time followed by env vars declaration:
```bash
clever create -t go
clever env set CC_GO_BUILD_TOOL "gomod"
clever env set CC_GO_PKG "cmd/main.go"
git add . && git commit -m "First commit"
clever deploy && clever domain
```
