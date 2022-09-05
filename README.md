# Using ossf scorecard API

- https://api.securityscorecards.dev/
- https://github.com/ossf/scorecard#scorecards-rest-api

## Using the API
The API is available at https://api.securityscorecards.dev/. This API doesnt require any authentication. You can use the API to get the scorecard for a repository. The API is a REST API and it returns JSON.

## Example
```
curl -X GET "https://api.securityscorecards.dev/projects/github.com/ossf/scorecard" -H "accept: application/json" | jq
```

## Demo code
The demo code uses the API to get the scorecard for all the dependencies of a repository.
This code uses parse the `go.mod` file to get the dependencies and then uses the API to get the scorecard for each dependency.

## Running the demo code
```
go run main.go PATH_TO_GO_MOD_FILE_DIR
```

The demo code shows all the dependencies that have been fuzzed.



