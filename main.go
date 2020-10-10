package main

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

// run command line --profile=prod  --port=8094
func main() {
	// config object to read config.yml
	configuration := Abstractions.NewConfigurationBuilder().AddYamlFile("config").Build()
	// init web host
	YoyoGo.NewWebHostBuilder().
		UseConfiguration(configuration).
		Configure(func(app *YoyoGo.WebApplicationBuilder) {
			app.UseEndpoints(func(router Router.IRouterBuilder) {
				router.POST("/alert", PostAlert)
			})
		}).Build().Run()
}
