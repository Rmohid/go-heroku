# go-template
###Template web api for go cli applications

This is a command-line application template that contains a simple key value store for internal use with a web api (on port 7100 by default) that supports the following:

1. Dump key values in plaintext
  * curl localhost:7100
2. Individually add a key/value
  * curl localhost:7100?newkey=value
3. Dump kv store to JSON
  * curl localhost:7100/json > kvdata.json
4. Update kv store from JSON
  * curl -H "Content-Type: application/json" --data @kvdata.json http://localhost:7100
5. Delete individual key values
  * curl -X DELETE localhost:7100?someKey
6. Dump an individual key
  * curl localhost:7100/key/someKey

This functionality can be added to a existing go application by including `github.com/rmohid/go-template/config` as a standalone package. Main and webExternal are not needed.

It also includes a debug logger `dbg.Log()` that supports the following:

1. Runtime changes to log verbosity, higher values for more detailed output
  * `dbg.Log(2,"Will only show for verbosity 2 or above")`
  * curl localhost:7100?dbg.verbosity=2
2. Log output to sdout/sderr or /dev/null
  * curl localhost:7100?dbg.debugWriter=sdtout
  * curl localhost:7100?dbg.debugWriter=sdterr
  * curl localhost:7100?dbg.debugWriter=devnull
3. Log output as form data to a web server
  * curl localhost:7100?dbg.httpUrl=localhost:7000
  * curl localhost:7100?dbg.debugWriter=http
4. Log output to a file
  * curl localhost:7100?dbg.logfile=./somefile.log
  * curl localhost:7100?dbg.debugWriter=file

This functionality can be added by including `github.com/rmohid/go-template/dbg`
