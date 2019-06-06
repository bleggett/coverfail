# Coverfail

Dumb little tool to get a consolidated code
coverage number for all packages in a Golang repo, and return a nonzero exit
code to the parent process if that number is less than a provided threshold.

`go test -cover ./...` will report coverage for each module it finds in the repo separately,
but cannot give you an overall coverage percentage for all `.go` files found, which this does.

## Dependencies
This runs `go test -cover` and parses the output, so if you can build this
you can run it.

## Example invocation
1. `cd <yourgoreporoot>`
2. `coverfail -threshold 55.6`

#### Example output if your code coverage > threshold:
```
Overall code coverage percentage is:  57.5
Threshold is:  55.6

(returncode == 0)
```

#### Example output if your code coverage < threshold:
```
Overall code coverage percentage is:  54.5
Threshold is:  55.6

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
