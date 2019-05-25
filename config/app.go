package config

func NewAppConfig() *AppConfig {
	return &AppConfig{
		/*
			|--------------------------------------------------------------------------
			| Environment Name
			|--------------------------------------------------------------------------
			|
			| With the help of the environment name we can implement different behaviour for
			| our environments. For example for local development we like to have another
			| logger encoding with colors, but in production we have the JSON encoding.
			|
		*/

		Env: Env("APP_ENV", "production"),

		/*
			|--------------------------------------------------------------------------
			| Application Name
			|--------------------------------------------------------------------------
			|
			| This value is the name of your application. This value is used when the
			| framework needs to place the application's name in a notification or
			| any other location as required by the application or its packages.
			|
		*/

		Name: Env("APP_NAME", "go-trophy-server"),

		/*
			|--------------------------------------------------------------------------
			| Application Port
			|--------------------------------------------------------------------------
			|
			| This value define on witch port the application is available. Default is
			| the standard port 8080
			|
		*/

		Port: Env("PORT", "8080"),

		/*
			|--------------------------------------------------------------------------
			| Application Name
			|--------------------------------------------------------------------------
			|
			| This value is the name of your application. This value is used when the
			| framework needs to place the application's name in a notification or
			| any other location as required by the application or its packages.
			|
		*/

		Prefix: Env("APP_PREFIX", "/api"),

		/*
			|--------------------------------------------------------------------------
			| Log Level
			|--------------------------------------------------------------------------
			|
			| TODO
			|
		*/

		LogLevel: Env("APP_LOG_LEVEL", "error"),

		/*
			|--------------------------------------------------------------------------
			| Simultaneous Connections
			|--------------------------------------------------------------------------
			|
			| By default, http.ListenAndServe (which gin.Run wraps) will serve an unbounded
			| number of requests. Limiting the number of simultaneous connections can
			| sometimes greatly speed things up under load.
			|
		*/

		Connection: EnvInt("APP_CONNECTIONS", 20),

		/*
			|--------------------------------------------------------------------------
			| Application Banner
			|--------------------------------------------------------------------------
			|
			| If you wish to see a nice log statement when the application has been
			| bootstrap then activate this config.
			|
		*/

		ShowBanner: EnvBool("APP_SHOW_BANNER", true),
	}
}

type AppConfig struct {
	Env        string
	Name       string
	Port       string
	Prefix     string
	LogLevel   string
	ShowBanner bool
	Connection int
}
