#coveo-cli

**Simple query**

    $ coveo-cli -n 2 -q layout

    Results: 2, Skipped: 0, Total: 804, Duration: 96ms

    Re: force layout question	|	Gmail - coveodocumentationsamples@gmail.com
    Re: force layout question	|	Gmail - coveodocumentationsamples@gmail.com

**Less simple - simple query**

    $ coveo-cli -n 5 -q "layout @syssource=salesforce"

    # oups no results
     Results: 5, Skipped: 0, Total: 0, Duration: 102ms


**Specifying fields to get**

    coveo-cli -n 2 -q "@sysconcepts" -f systitle,objecttype,sysconcepts

    Results: 2, Skipped: 0, Total: 26396, Duration: 109ms
    Re: Apache OpenOffice.org Calc	|	Message	|	openoffice ; org ; helper column ; search criteria ; confusing think ; bug reporting ; choice of filters ; apache ; Formatting ; programmer ; submission
    Re: [dc.js users] Simple line chart not drawing data points	|	Message	|	list of colors ; emails ; unsubscribe ; unmunged data ; googlegroups ; dc-js-user-group ; empty ; graph ; ordinalColors

**Getting facets**

    coveo-cli -n 5 -q "layout" -g objecttype,sysconcepts
    Results: 5, Skipped: 0, Total: 804, Duration: 101ms
    objecttype:
               Thread : 0
               Message : 0
               File : 0
    sysconcepts:
                d3 : 0
                googlegroups : 0
                unsubscribe : 0
                emails : 0
                stop receiving : 0


**For help**

    coveo-cli --help

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
