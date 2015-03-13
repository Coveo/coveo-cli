![Logo](doc/coveo-cli.png)

#coveo-cli


## Get

https://github.com/Coveo/coveo-cli/releases

**Windows**

https://github.com/Coveo/coveo-cli/releases/download/v0.2/coveo-cli-windows-x64.zip

**Linux**
https://github.com/Coveo/coveo-cli/releases/download/v0.2/coveo-cli-linux-amd64.zip

**OSX**
https://github.com/Coveo/coveo-cli/releases/download/v0.2/coveo-cli-darwin-amd64.zip

## Developers

**Respect some standards**
```bash
go fmt .
go fix .
go vet .
```

**Build**
```bash
go generate
GOOS=darwin  go build -o dist/darwin-amd64/coveo-cli .
GOOS=windows go build -o dist/windows-amd64/coveo-cli.exe .
GOOS=linux   go build -o dist/linux-amd64/coveo-cli .
```

**Tests (if there were any)**
```bash
go test .
go test -cover .
```

## Usage
**Simple query**

```bash

    $ coveo-cli -n 2 -q layout

    Result: [0-1]/804, Duration: 103ms


    Re: force layout question
    imap://imap.gmail.com:993/account:coveodocumentationsamples@gmail.com/mailbox:[Gmail]/mailbox:All Mail/mail:19386
    	systitle: Re: force layout question
    	syssource: Gmail - coveodocumentationsamples@gmail.com
    Re: force layout question
    imap://imap.gmail.com:993/account:coveodocumentationsamples@gmail.com/mailbox:[Gmail]/mailbox:All Mail/mail:19397
    	systitle: Re: force layout question
    	syssource: Gmail - coveodocumentationsamples@gmail.com
```

**Less simple - simple query**

```bash
$ coveo-cli -n 5 -q "layout @syssource=salesforce"

# oups no results
Result: [0-4]/0, Duration: 84ms
```

**Specifying fields to get**

```bash
$ coveo-cli -n 2 -q "@sysconcepts" -f objecttype,sysconcepts

Result: [0-1]/26413, Duration: 99ms


use of -blibpath not consistent (on AIX, maybe for all platforms)
imap://imap.gmail.com:993/account:coveodocumentationsamples@gmail.com/mailbox:[Gmail]/mailbox:All Mail/mail:19426
	objecttype: Message
	sysconcepts: lib ; usr ; multiplicity ; side-by-side ; packaging ; libressl ; archive member ; NIX whith ; libcrypto
Re: spell check
imap://imap.gmail.com:993/account:coveodocumentationsamples@gmail.com/mailbox:[Gmail]/mailbox:All Mail/mail:19423
	objecttype: Message
	sysconcepts: spell check ; User Profile ; suspect
```

**Getting facets**

```bash
$ coveo-cli -n 5 -q "layout" -g objecttype,sysconcepts
Result: [0-4]/804, Duration: 108ms

objecttype:
           Thread : 1
           Message : 778
           File : 24
sysconcepts:
            d3 : 539
            googlegroups : 539
            unsubscribe : 509
            emails : 377
            stop receiving : 225
```

**For help**

```bash
$ coveo-cli --help

coveo-cli: usage
  -e="https://cloudplatform.coveo.com/rest/search/": access endpoint
  -f="systitle,syssource": fields to show
  -g="": Facets to query, if you query facets you cant query normal results
  -h=false: show query count & duration
  -help=false: show query count & duration
  -json=false: print original json format
  -n=10: numbers of results to return
  -p="": Password
  -q="": Query "q" term
  -s=true: show query count & duration
  -skip=0: number of results to skip
  -t="52d806a2-0f64-4390-a3f2-e0f41a4a73ec": access token
  -u="": Username
```
