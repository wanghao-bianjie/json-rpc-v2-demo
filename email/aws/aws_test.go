package aws

import "testing"

func TestSendEmail(t *testing.T) {
	recipient := "wanghao@bianjie.ai"
	subject := "Amazon SES Test (AWS SDK for Go)"
	htmlBody := "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"
	textBody := "This email was sent with Amazon SES using the AWS SDK for Go."
	err := SendEmail(recipient, subject, htmlBody, textBody)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("send email success")
}
