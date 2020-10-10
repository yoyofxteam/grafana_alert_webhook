package main

import (
	"GrafanaAlertWebHook/services"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

/*
PostAlert : Receive grafana alert post json message
 - Endpoint : /alert
*/
func PostAlert(ctx *Context.HttpContext) {
	var request services.GrafanaAlertRequest
	_ = ctx.Bind(&request) // json binding to object
	var config Abstractions.IConfiguration
	_ = ctx.RequiredServices.GetService(&config) // IOC get config object

	ctx.JSON(200, Context.H{
		"Message": services.WechatSendMarkdownMessage(request, config),
	})
}
