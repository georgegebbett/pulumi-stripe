# Stripe Resource Provider

The Stripe Resource Provider lets you manage [Stripe](https://stripe.com) resources.

This is a bridged provider from https://github.com/lukasaron/terraform-provider-stripe

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install pulumi-stripe
```

or `yarn`:

```bash
yarn add pulumi-stripe
```

or indeed `pnpm`:

```bash
pnpm add pulumi-stripe
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-stripe
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/georgegebbett/pulumi-stripe/sdk/go
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Stripe
```

## A note on Go and .NET

I have never done any real development in Go or .NET - indeed this was my first foray into Go at all. I cannot warrant
that I have published the packages correctly, or provided good instructions for installing them.

Please feel free to open a PR with any suggestions!

## Configuration

The following configuration points are available for the `stripe` provider:

- `stripe:apiKey` (environment: `STRIPE_API_KEY`) - the API key for `stripe`
