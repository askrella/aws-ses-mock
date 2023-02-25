package internal

import (
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("AWS SES Tests", func() {
	Context("Sending Emails", func() {
		client := resty.New()

		go func() {
			StartServer()
		}()

		It("Successful raw email", func() {
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

			// Error should be not be nil
			Ω(err).Should(BeNil())

			// Error should be not be nil
			Ω(string(response.Body())).Should(ContainSubstring("SendEmailResult"))
		})

		It("Should fail if ToAddresses is not provided", func() {
			client := resty.New()

			response, err := client.R().SetBody(SendEmailRequest{
				Action: "SendEmail",
				Destination: Destination{
					ToAddresses:  []string{},
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

			// Error should be not be nil
			Ω(err).Should(BeNil())

			// Response should be: one or more required fields was not sent
			Ω(string(response.Body())).Should(ContainSubstring("one or more required fields was not sent"))
		})
	})
})

func TestEmailSending(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "")
}
