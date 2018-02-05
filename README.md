I've tried **three** different ways of setting up dependencies.

However, regardless of the approach I've used, no approach seems to work for all three of these build / hosting scenarios:

| hosting / build scenario            | classic `src` dependecies | `dep` vendoring | `dep` + separate `gaedef` |
| :---------------------------------- | :-----------------------: | :-------------: | :-----------------------: |
| Local `go run main.go`              | SUCCESS                   | SUCCESS         | SUCCESS                   |
| Local `dev_appserver.py app.yaml`   | ERROR                     | ERROR           | ERROR                     |
| Remote `gcloud app deploy app.yaml` | SUCCESS                   | ERROR           | ERROR                     |

## Local `go run main.go`

All the ways of setting up my dependencies work as expected when treating the app as just a simple Go app, ala:

```
> go run main.go
2018/02/04 16:00:18 Initializing helloservice
2018/02/04 16:00:18 Starting app engine
2018/02/04 16:01:42 `Hello` sent to client
```

## Local `dev_appserver.py app.yaml`

However, I run into problems using _any_ approach when running the local dev server, ala:

```
> dev_appserver.py app.yaml
INFO     2018-02-04 16:08:12,414 devappserver2.py:105] Skipping SDK update check.
INFO     2018-02-04 16:08:13,013 api_server.py:308] Starting API server at: http://localhost:61261
INFO     2018-02-04 16:08:13,115 dispatcher.py:255] Starting module "default" running at: http://localhost:8080
INFO     2018-02-04 16:08:13,122 admin_server.py:146] Starting admin server at: http://localhost:8000
ERROR    2018-02-04 16:08:14,420 module.py:1602]
INFO     2018-02-04 16:08:15,125 shutdown.py:45] Shutting down.
INFO     2018-02-04 16:08:15,127 api_server.py:971] Applying all pending transactions and saving the datastore
INFO     2018-02-04 16:08:15,128 api_server.py:974] Saving search indexes
```

## Remote `gcloud app deploy`

Deploying to `gcloud` works just fine and dandy when using the `src` method of handling dependencies.  However, it totally blows up when I try to deploy using the `dep` methodology:

```
gcloud app deploy app.yaml
Services to deploy:

descriptor:      [C:\Projects\junk-dev-app\src\helloworld\app.yaml]
source:          [C:\Projects\junk-dev-app\src\helloworld]
target project:  [junk-dev-app]
target service:  [default]
target version:  [20180204t161006]
target url:      [https://junk-dev-app.appspot.com]


Do you want to continue (Y/n)?  y

Beginning deployment of service [default]...
Some files were skipped. Pass `--verbosity=info` to see which ones.
You may also view the gcloud log file, found at
[C:\Users\Jonathan\AppData\Roaming\gcloud\logs\2018.02.04\16.09.45.084000.log].
╔════════════════════════════════════════════════════════════╗
╠═ Uploading 0 files to Google Cloud Storage                ═╣
╚════════════════════════════════════════════════════════════╝
File upload done.
Updating service [default]...failed.
ERROR: (gcloud.app.deploy) Error Response: [9] Deployment contains files that cannot be compiled: Compile failed:
2018/02/04 14:11:07 go-app-builder: Failed parsing input: parser: bad import "syscall" in vendor/golang.org/x/net/icmp/endpoint.go
```

So, regardless of the approach I've used, no approach seems to work for all three scenarios.


# Tree 1 - Dependencies via `src`

Just to be clear, as I'm a total noob, this is what I mean when I say I am using the `src` dependency method.

```
GOPATH (c:\projects\junk-test-app\)
└───src
    ├───github.com
    │   └───golang
    ├───golang.org
    │   └───x
    ├───google.golang.org
    │   └───appengine
    ├───helloservice
    └───helloworld
        ├───main.go
        └───app.yaml
```

## Tree 1 - Errors

No errors occur when running the program locally or deploying to GAE via `gcloud app deploy`.

However, locally hosting the app via `dev_appserver.py app.yaml` terminates with the following error:
```
ERROR    2018-02-04 16:08:14,420 module.py:1602]
```

# Tree 2 -- Vendoring via `dep`

And, this is what I mean when I say I am using the `dep` dependency method.

```
GOPATH (c:\projects\junk-test-app\)
└───src
    ├───helloservice
    └───helloworld
        │   ├───main.go
        │   └───app.yaml
        └───vendor
            ├───github.com
            │   └───golang
            ├───golang.org
            │   └───x
            └───google.golang.org
                └───appengine
```

## Tree 2 - Errors

No errors occur when running the program locally.

However, locally hosting the app via `dev_appserver.py app.yaml` or deploying to GAE via `gcloud app deploy` terminates with the errors below.

`gcloud app deploy` error:
```
ERROR: (gcloud.app.deploy) Error Response: [9] Deployment contains files that cannot be compiled: Compile failed:
2018/02/04 14:11:07 go-app-builder: Failed parsing input: parser: bad import "syscall" in vendor/golang.org/x/net/icmp/endpoint.go
```

`dev_appserver.py app.yaml` error:
```
ERROR    2018-02-04 16:08:14,420 module.py:1602]
```

Tree 3 - Vendoring via `dep` (alternate placement)
==============================================================

This tree suggested by [@derekperkins](https://github.com/derekperkins) places the GAE definition into a folder that has no subfolders based on an assumption that `gcloud app deploy` might not like to see dependencies contained in subfolders of the `main.go` app bootstrap definition.

```
GOPATH (c:\projects\junk-test-app\)
└───src
    ├───helloservice
    └───helloworld
        ├───gaedef (Derek calls this folder "service")
        │   ├───main.go
        │   └───app.yaml
        └───vendor
            ├───github.com
            │   └───golang
            ├───golang.org
            │   └───x
            └───google.golang.org
                └───appengine
```

## Tree 3 - Errors

In this scenario, no errors occur when running the program locally.

However, locally hosting the app via `dev_appserver.py app.yaml` or deploying to GAE via `gcloud app deploy` terminates with the errors below, *which are different* than than the errors seen in *Tree 2*. 

`gcloud app deploy` error:
```
ERROR: (gcloud.app.deploy) Error Response: [9] Deployment contains files that cannot be compiled: Compile failed:
2018/02/04 17:09:39 go-app-builder: Failed parsing input: package "helloworld/vendor/google.golang.org/appengine" cannot import internal package "google.golang.org/appengine/internal/modules"
```

`dev_appserver.py app.yaml` error:
```
ERROR    2018-02-04 19:08:43,239 instance_factory.py:196] Failed to build Go application: (Executed command: C:\Users\Jonathan\AppData\Local\Google\Cloud SDK\google-cloud-sdk\platform\google_appengine\goroot-1.8\bin\go-app-builder.exe -app_base C:\Projects\junk-dev-app\src\helloworld\gaedef -api_version go1.8 -arch 6 -dynamic -goroot C:\Users\Jonathan\AppData\Local\Google\Cloud SDK\google-cloud-sdk\platform\google_appengine\goroot-1.8 -gopath C:/Projects/junk-dev-app -nobuild_files ^^$ -incremental_rebuild -unsafe -binary_name _go_app -extra_imports appengine_internal/init -work_dir c:\users\jonathan\appdata\local\temp\tmpoyfixwappengine-go-bin -gcflags -I,C:\\Users\\Jonathan\\AppData\\Local\\Google\\Cloud SDK\\google-cloud-sdk\\platform\\google_appengine\\goroot-1.8\\pkg\\linux_amd64_appengine -ldflags -L,C:\\Users\\Jonathan\\AppData\\Local\\Google\\Cloud SDK\\google-cloud-sdk\\platform\\google_appengine\\goroot-1.8\\pkg\\linux_amd64_appengine main.go)
C:\Projects\junk-dev-app\src\helloworld\vendor\golang.org\x\net\context\go17.go:10: import C:\Users\Jonathan\AppData\Local\Google\Cloud SDK\google-cloud-sdk\platform\google_appengine\goroot-1.8\pkg\linux_amd64_appengine/context.a: object is [linux amd64 1.8.5 (appengine-1.9.65) X:framepointer] expected [windows amd64 1.8.5 (appengine-1.9.65) X:framepointer]
C:\Projects\junk-dev-app\src\helloservice\helloservice.go:4: import C:\Users\Jonathan\AppData\Local\Google\Cloud SDK\google-cloud-sdk\platform\google_appengine\goroot-1.8\pkg\linux_amd64_appengine/fmt.a: object is [linux amd64 1.8.5 (appengine-1.9.65) X:framepointer] expected [windows amd64 1.8.5 (appengine-1.9.65) X:framepointer]
```

