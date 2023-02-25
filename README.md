# AWS SES Mock

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Docker](https://github.com/askrella/whatsapp-chatgpt/actions/workflows/docker.yml/badge.svg)
![Build](https://img.shields.io/github/actions/workflow/status/askrella/aws-ses-mock/docker.yml?branch=master)

Example Json for SendMail:

```json
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
```