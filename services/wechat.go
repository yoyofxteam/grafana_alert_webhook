package services

import (
	"encoding/json"
	"fmt"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
)

// wechat send markdown message
func WechatSendMarkdownMessage(request GrafanaAlertRequest, config Abstractions.IConfiguration) string {
	tag := request.GetTag()
	logger := XLog.GetXLogger("wechat")
	js, _ := json.Marshal(request)
	logger.Info("Request json: %s", string(js))
	if tag == "" {
		logger.Info("no send")
		return ""
	}
	sendUrl := config.Get(fmt.Sprintf("alert.%s.webhook_url", tag)).(string)
	linkUrl := config.Get(fmt.Sprintf("alert.%s.link_url", tag)).(string)
	logger.Info("request tag:%s", tag)
	logger.Info(sendUrl)
	logger.Info(linkUrl)

	var message *MarkdownMessage
	if request.State == "alerting" && len(request.EvalMatches) > 0 {
		message = &MarkdownMessage{
			Markdown: struct {
				Content string `json:"content" gorm:"column:content"`
			}{
				Content: "## " + request.RuleName + ",请相关同事注意。\n" +
					" > [报警信息] : " + request.Message + "\n" +
					" > [报警次数] : <font color=\"warning\">" + request.GetMetricValue() + "次</font>" + "\n" +
					" > [报警明细] : (" + linkUrl + ")\n",
			},
			Msgtype: "markdown",
		}
	}
	msg, _ := json.Marshal(message)
	msgStr := string(msg)
	logger.Info("send message:%s", msgStr)

	return HttpPost(sendUrl, msgStr)
}
