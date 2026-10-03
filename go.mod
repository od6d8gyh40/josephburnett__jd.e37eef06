module github.com/josephburnett/jd

go 1.26.0

require (
	github.com/go-openapi/jsonpointer v1.0.2
	github.com/josephburnett/jd/v2 v2.0.0-20240818191833-6125a15c637a
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/kr/text v0.2.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/josephburnett/jd/v2 => ./v2
