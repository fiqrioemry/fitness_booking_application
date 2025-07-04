package config

func InitConfiguration() {
	LoadEnv()
	InitRedis()
	InitMailer()
	InitDatabase()
	InitCloudinary()
	InitGoogleOAuthConfig()
	GetTrustedProxies()
	InitStripe()
}
