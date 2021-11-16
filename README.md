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