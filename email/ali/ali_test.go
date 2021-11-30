package ali

import (
	"testing"

	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	"github.com/alibabacloud-go/tea/tea"
)

func TestCreateClient(t *testing.T) {
	accessKeyId := "***"
	accessKeySecret := "***"
	accountName := "***"
	toAddress := "wanghao@bianjie.ai"
	subject := "AliCloud Email Test"
	htmlBody := "<h1>AliCloud Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	client, err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		t.Fatal(err)
	}
	singleSendMailRequest := &dm20151123.SingleSendMailRequest{
		AccountName:    tea.String(accountName),
		AddressType:    tea.Int32(1),
		ReplyToAddress: tea.Bool(true),
		ToAddress:      tea.String(toAddress),
		Subject:        tea.String(subject),
		HtmlBody:       tea.String(htmlBody),
	}
	_res, err := client.SingleSendMail(singleSendMailRequest)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(_res.String())
}
