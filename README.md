#  AWS SES Mock


[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Docker](https://github.com/askrella/whatsapp-chatgpt/actions/workflows/docker.yml/badge.svg)
![Build](https://img.shields.io/github/actions/workflow/status/askrella/aws-ses-mock/docker.yml?branch=master)


![Askrella](https://avatars.githubusercontent.com/u/77694724?s=100)

We created this project as a new version of aws-ses-local, which doesn't seem to be maintained for a few years.
Our goal is to provide more features, small containers and be more accurate than the alternatives.

# :gear: Getting Started

## Running the Docker Container

```bash
docker run -p 8080:8080 ghcr.io/askrella/aws-ses-mock:1.0.1
```

## Usage with NodeJS

Using the AWS SDK you can set the endpoint for SES manually by specifying the endpoint in your configuration:

```javascript
import AWS from 'aws-sdk'
const ses = new AWS.SES({ region: 'us-east-1', endpoint: 'http://localhost:8080' })
```

## Usage with Golang

Using the AWS SDK you can set the endpoint for SES manually by overriding the endpoint resolver:

```golang
customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
    if overrideEndpoint, exists := os.LookupEnv("OVERRIDE_SES_ENDPOINT"); exists {
        return aws.Endpoint{
            PartitionID:   "aws",
            URL:           overrideEndpoint,
            SigningRegion: "eu-central-1",
        }, nil
    }

    return aws.Endpoint{}, &aws.EndpointNotFoundError{}
})

cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"), config.WithEndpointResolverWithOptions(customResolver))
```

## Manual testing

The `POST` endpoint is available under `http://localhost:8080/` and should contain the raw JSON body used for SES messages:
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

## :test_tube: Running Tests

To run tests, run the following command

```bash
  go test ./internal/*
```


# :wave: Contributors

<a href="https://github.com/askrella/aws-ses-mock/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=askrella/aws-ses-mock" />
</a>

* [Askrella Software Agency](askrella.de)
  * [Steve](https://github.com/steve-hb) (Maintainer)
  * [Navo](https://github.com/navopw) (Maintainer)

Feel free to open a new pull request with changes or create an issue here on GitHub! :)

# :warning: License
Distributed under the MIT License. See LICENSE.txt for more information.

# :handshake: Contact Us

In case you need professional support, feel free to <a href="mailto:contact@askrella.de">contact us</a>
