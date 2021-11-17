mongodb web app - web interface for logging events [WIP]

In order to allow system with signature ID ```SYSTEMXYZ``` to log message ```System failure detected```

Configure signature ID in whitelist.json like that:

```json
{
    "systems": [{
            "ID": "SYSTEMXYZ",
            "Description": "Mail server 1"
        },
        {
            "ID": "SYSTEMABC",
            "Description": "File server 1"
        },
        {
            "ID": "SYSTEMWZT",
            "Description": "File server 2"
        }
    ]
}
```


Run:
```go
go run main.go
```

Call /save with parameter ```signature``` "SYSTEMXYZ" and ```message``` "System failure detected":

```bash
http://127.0.0.1:8080/save?signature=SYSTEMXYZ&message=System failure detected
```

Result:
```mongo
{ "_id" : ObjectId("6193e1520dcdee3f35f5faff"), "signature" : "SYSTEMXYZ", "ts" : ISODate("2021-11-16T16:50:26.014Z"), "message" : "system failure detected on XYZ" }

```

If host is not allowed (e.g: ```SYSTEMFOO```, its signature ID is not in whitelist), then result is:

```bash
http://127.0.0.1:8080/save?signature=SYSTEMFOO&message=system%20failure%20detected%20on%20XYZ

```
Result:
```bash
Error: authentication denied
```

In order to retrieve objects for registered systems, you have to call ```/load```  with parameter ```signature```:
```bash
http://127.0.0.1:8080/load?signature=SYSTEMXYZ
```

Result:
```bash
[
    {
        "Signature": "SYSTEMXYZ",
        "TS": "2021-11-16T16:50:26.014Z",
        "Message": "system failure detected on XYZ"
    },
    {
        "Signature": "SYSTEMXYZ",
        "TS": "2021-11-16T17:37:06.46Z",
        "Message": "system failure detected on XYZ"
    },
    {
        "Signature": "SYSTEMXYZ",
        "TS": "2021-11-16T17:59:31.746Z",
        "Message": "system failure detected on XYZ"
    }
]
```

Test:
```golang
(main)$ go test -v -cover ./...
?   	github.com/deeper-x/weblog	[no test files]
=== RUN   TestNewInstance
--- PASS: TestNewInstance (0.00s)
=== RUN   TestCreateCtx
--- PASS: TestCreateCtx (0.00s)
=== RUN   TestCreateClient
--- PASS: TestCreateClient (0.00s)
=== RUN   TestCreateCollection
--- PASS: TestCreateCollection (0.00s)
=== RUN   TestConnect
--- PASS: TestConnect (0.00s)
=== RUN   TestClose
--- PASS: TestClose (0.00s)
=== RUN   TestAddEntry
--- PASS: TestAddEntry (0.01s)
=== RUN   TestSaveEntry
--- PASS: TestSaveEntry (0.00s)
=== RUN   TestGetEntries
--- PASS: TestGetEntries (0.00s)
PASS
coverage: 59.4% of statements
ok  	github.com/deeper-x/weblog/db	0.038s	coverage: 59.4% of statements
?   	github.com/deeper-x/weblog/messages	[no test files]
?   	github.com/deeper-x/weblog/settings	[no test files]
=== RUN   TestReadJSONFile
--- PASS: TestReadJSONFile (0.00s)
PASS
coverage: 46.2% of statements
ok  	github.com/deeper-x/weblog/wauth	0.004s	coverage: 46.2% of statements
=== RUN   TestSave
2021/11/17 12:30:56 Saving entry - Signature: [SYSTEMXYZ]
2021/11/17 12:30:56 Entry saved succesfully ✔
--- PASS: TestSave (0.01s)
=== RUN   TestLoad
--- PASS: TestLoad (0.01s)
PASS
coverage: 40.8% of statements
ok  	github.com/deeper-x/weblog/web	0.029s	coverage: 40.8% of statements
```
