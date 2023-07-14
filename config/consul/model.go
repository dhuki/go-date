package consul

var (
	searchPath = []string{
		"$HOME/.go-date",
		"$GOPATH/src/github/go-date",
		".",
	}

	configName = map[string]string{
		"LOCAL": "config.local.yaml",
		"DEV":   "config.dev.yaml",
		"UAT":   "config.uat.yaml",
		"PROD":  "config.prod.yaml",
	}
)
