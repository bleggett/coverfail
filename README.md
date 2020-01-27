> NOTE: This is deprecated in favor of https://github.com/klmitch/overcover

# Coverfail

`go test -cover ./...` will report coverage for each package it finds in the repo separately,
but cannot give you an overall coverage percentage for the whole module, which this does.

Additionally, this lets you return a nonzero exit code to the parent process if the module's overall coverage falls
below a fixed amount provided via an argument. This makes it easy to integrate into most CI systems.

> NOTE: This was created as a quick hack to work around the (to me) puzzling inability of `go test -cover` to generate overall coverage for a Go module composed of multiple packages, and should be abandoned as soon as upstream adds a more robust implementation of this relatively basic capability.

## Dependencies
This runs `go test -cover` and parses the output, so if you can build this
you can run it.

## Example invocation
1. `cd <yourgoreporoot>`
2. `coverfail -threshold 55.6`

#### Example output if your code coverage > threshold:
```
Threshold is:  51
Overall coverage: 51.2% of statements    

(returncode == 0)
```                                                                                                                                                                       

#### Example output if your code coverage < threshold:
```
Threshold is:  53
Overall coverage: 51.2% of statements  

Overall coverage is lower than provided threshold number, bailing with nonzero exit code...

(returncode == 1)
```


#### Example output if `go test -cover -coverpkgs=./... ./...` fails (just dumps that output):
```
?       gitlab.com/mytool        [no test files]
?       gitlab.com/mytool/apply  [no test files]
--- FAIL: TestGetAccessLevel (0.00s)
    client_test.go:27:
                Error Trace:    client_test.go:27
                Error:          Should not be: "Developer"
                Test:           TestGetAccessLevel
                Messages:       Expected 30 to map to 'Developer' but got Developer instead
2019/04/16 15:11:12 JSON group create body: {"name":"test1","path":"test1"}
2019/04/16 15:11:12 Couldn't find group test2 in subgroups of group id: 1
2019/04/16 15:11:12 JSON group create body: {"name":"test2","path":"test2","parent_id":1}
2019/04/16 15:11:12 Couldn't find group test2 in subgroups of group id: 1
2019/04/16 15:11:12 JSON group create body: {"name":"test2","path":"test2","parent_id":1}
FAIL
coverage: 49.8% of statements in ./...
FAIL    gitlab.com/mytool/client      0.031s
?       gitlab.com/mytool/cmd    [no test files]
ok      gitlab.com/mytool/parse  0.028s  coverage: 4.7% of statements in ./...
'go test' exited with an error, no coverage results available

(returncode == 1)
```
