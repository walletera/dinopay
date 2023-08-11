# DinoPay API

DinoPay is a fictitious payment service provider created as part of the [Walletera](https://fedevmoya.hashnode.dev/designing-a-digital-wallet#heading-a-fictitious-payment-service-provider-dinopay) project. 

What's on this repo
- the Dinopay OpenApi specification
- a generated golang client for the Dinopay API

## How to generate the golang client

To generate the golang client code from the openapi specification you need to install [ogen](https://ogen.dev/) by executing in a terminal
```bash
go install -v github.com/ogen-go/ogen/cmd/ogen@latest
```

Once you have the generator installed open a terminal an execute the following command
```bash
go generate ./...
```