package settings

var Settings = &yamlSettings{
	Log: log{
		Level: "debug",
		File:  "",
	},
	Bot: bot{
		UseProxy: true,
	},
}
