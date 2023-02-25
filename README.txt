# AWS SES Mock

Example Json for SendMail:

`
{
    "Action": "SendEmail",
    "Destination": {
        "ToAddresses": [
            "recipient@example.com"
        ],
        "CcAddresses": [
            "cc@example.com"
        ],
        "BccAddresses": [
            "bcc@example.com"
        ]
    },
    "Message": {
        "Body": {
            "Text": {
                "Data": "This is the message body in plain text format."
            },
            "Html": {
                "Data": "<html><body><h1>Hello World!</h1><p>This is the message body in HTML format.</p></body></html>"
            }
        },
        "Subject": {
            "Data": "Test email"
        }
    },
    "Source": "sender@example.com",
    "ReplyToAddresses": [
        "reply-to@example.com"
    ]
}
`