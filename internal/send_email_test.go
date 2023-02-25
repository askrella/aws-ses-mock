package internal

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("AWS SES Tests", func() {
	Context("Sending Emails", func() {
		It("Successful raw email", func() {
			go func() {
				StartServer()
			}()

			client := resty.New()
			response, err := client.R().SetBody(SendEmailRequest{
				Action: "SendEmail",
				Destination: Destination{
					ToAddresses:  []string{"steve+aws-ses-mock@askrella.dev"},
					CcAddresses:  []string{"stanislav+aws-ses-mock@askrella.dev"},
					BccAddresses: []string{"paul+aws-ses-mock@askrella.dev"},
				},
				Message: Message{
					Body: Body{
						Text: Content{
							Data: "This is the message body in plain text format.",
						},
						Html: Content{
							Data: "<html><body><h1>Hello World!</h1><p>This is the message body in HTML format.</p></body></html>",
						},
					},
					Subject: Subject{
						Data: "Test Email",
					},
				},
				Source: "sender@example.com",
				ReplyToAddresses: []string{
					"reply-to@example.com",
				},
			}).Post("http://localhost:8080")
			if err != nil {
				_ = fmt.Errorf("error while sending email: %s", err.Error())
			}

			fmt.Printf("got response %s", string(response.Body()))
		})

	})
})

func TestEmailSending(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "")
}
