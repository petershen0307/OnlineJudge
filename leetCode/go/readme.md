# How to build Go?

## Example path:
* c:\leetCode\
    * go\
        * src\
            * 104max_depth\
                * main.go
        * bin\
            * 104max_depth.exe
```
set GOPATH=c:\leetCode\go
go install 104max_depth
// enter folder name to build go project, not .go file name or package name in .go file
```

## Keyword:
* package
    * set to **main** will be the entry point(build to an execution file)
    * if package name is not **main**, it will build to a library file(xxx.a)

## Environment setting:
* IDE: VS code
    * install extension **Go** which is Microsoft/vscode-go on github
    * user preference setting:
    ```
    {
        "go.buildOnSave": true,
        "go.lintOnSave": true,
        "go.vetOnSave": true,
        "go.buildTags": "",
        "go.buildFlags": [],
        "go.lintTool": "golint",
        "go.lintFlags": [],
        "go.vetFlags": [],
        "go.coverOnSave": false,
        "go.useCodeSnippetsOnFunctionSuggest": false,
        "go.formatOnSave": true,
        "go.formatTool": "goreturns",
        "go.formatFlags": [],
        "go.goroot": "C:\\programming\\Go", // this is where go installed path
        "go.gopath": "C:\\programming\\gotool", // this is default go path and it can be the place where go tools for vscode download
        "go.gocodeAutoBuild": false
    }
    ```
    * customize build in vscode:
        * **ctrl+shift+p or f1** to open **Tasks: Configure Task Runner**
        * it will create a **tasks.json** by vscode project
        * the settings
        ```
        {
            "version": "0.1.0",
            "command": "go",
            "isShellCommand": true,
            "showOutput": "silent",
            "options": {
                "env": {
                    "GOPATH": "C:\\programming\\Go"  //"YOUR_GO_PATH"
                }
            },
            "tasks": [
                {
                    "taskName": "install",
                    "args": [ "-v", "./..."],
                    "isBuildCommand": true
                },
                {
                    "taskName": "test",
                    "args": [ "-v", "./..."],
                    "isTestCommand": true
                }
            ]
        }
        ```
        * **ctrl+shift+b** to build the project
    * [ref1](golang-zhtw.netdpi.net)
    * [ref2](http://www.evanlin.com/dive-with-vscode-golang/)
    * [ref3](https://github.com/Microsoft/vscode-go)
