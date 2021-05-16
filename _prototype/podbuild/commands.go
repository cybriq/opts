package main

var commands = map[string][]string{
	"build": {
		"go generate ./version/.",
		"go build -v ./_prototype/.",
	},
	"install": {
		"go generate ./version/.",
		"go install -v ./_prototype/.",
	},
	"gui": {
		"go generate ./version/.",
		"go run -v ./_prototype/. gui",
	},
	"node": {
		"go generate ./version/.",
		"go run -v ./_prototype/. node",
	},
	"wallet": {
		"go generate ./version/.",
		"go run -v ./_prototype/.",
	},
	"kopach": {
		"go generate ./version/.",
		"go run -v ./_prototype/.",
	},
	"headless": {
		"go generate ./version/.",
		"go install -v -tags headless ./_prototype/.",
	},
	"docker": {
		"go generate ./version/.",
		"go install -v -tags headless ./_prototype/.",
	},
	"appstores": {
		"go generate ./version/.",
		"go install -v -tags nominers ./_prototype/.",
	},
	"tests": {
		"go generate ./version/.",
		"go test ./...",
	},
	"builder": {
		"go generate ./version/.",
		"go install -v ./_prototype/podbuild/.",
	},
	"generate": {
		"go generate ./...",
	},
}
