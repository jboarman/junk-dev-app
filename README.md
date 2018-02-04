I've tried two different ways of setting up dependencies.

However, regardless of the approach I've used, no approach seems to work for all three of these scenarios:

1. Local `go run main.go`
2. Local `dev_appserver.py app.yaml`
3. Remote `gcloud app deploy app.yaml`

Local `go run main.go`
======================

Both ways of setting up my dependencies work as expected when treating the app as just a simple Go app, ala:

```
> go run main.go
2018/02/04 16:00:18 Initializing helloservice
2018/02/04 16:00:18 Starting app engine
2018/02/04 16:01:42 `Hello` sent to client
```

Local `dev_appserver.py app.yaml`
=================================

However, I run into problems using _either_ approach when running the local dev server, ala:

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

Remote `gcloud app deploy`
==========================

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


Dependencies via `src`
=============================

Just to be clear, as I'm a total noob, this is what I mean when I say I am using the `src` dependency method.

```
GOPATH (c:\projects\junk-test-app\)
├───src
    ├───github.com
    │   └───golang
    ├───golang.org
    │   └───x
    ├───google.golang.org
    │   └───appengine
    ├───helloservice
    └───helloworld
```

Vendoring via `dep`
======================

And, this us what I mean when I say I am using the `dep` dependency method.

```
GOPATH (c:\projects\junk-test-app\)
├───src
    ├───helloservice
    └───helloworld
        └───vendor
            ├───github.com
            │   └───golang
            ├───golang.org
            │   └───x
            └───google.golang.org
                └───appengine
```