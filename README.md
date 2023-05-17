<h1>Ory Solution</h1>

A full Authentication and Authorisation Solution

<h2>Table of Contents</h2>

[TOC]

# Ory

- **Kratos** - Headless Identity and user management
- **Hydra** - Headless 0Auth 2.0 and OpenID Connect provider
- **Oathkeeper** - Identity and Access Proxy (IAP) - Zero trust networking proxy
- **Keto** - Access, authorisation and permission service - BRAC, ABAC, ACL, based on Google's Zanzibar

# Ory Hosted

## Base Setup

Ory Project Details

- **Ory Project:** Kratos
- **Slug:** keen-keller-k8bgx0ejit

Install the Ory CLI using homebrew

```sh
$ brew install ory/tap/cli
# ory help
```

Ory APIs will be available on `https://{your-slug}.projects.oryapis.com` which would be

 `https://keen-keller-k8bgx0ejit.projects.oryapis.com/ui/registration`

### Local Development

To allow Ory Hosted APIs to be available on your local development domain (`localhost`) and avoid third-party cookie issues use the [Ory CLI Tunnel or Proxy](https://www.ory.sh/docs/guides/cli/proxy-and-tunnel). Using the Proxy is an alternative to [custom domains (CNAME)](https://www.ory.sh/docs/guides/custom-domains), and a useful tool when developing locally. Even though you can use the Proxy in production, we recommend using a CNAME where possible.

```sh
$ export ORY_PROJECT_SLUG=<slug>
$ ory proxy application-url [publish-url] [flags]
```

So for our project

```sh
$ export ORY_PROJECT_SLUG=keen-keller-k8bgx0ejit
$ ory proxy http://localhost:3000 http://localhost:4000 --dev --debug
```

- The URL `http://localhost:3000` is where your application is available
- The `--dev` flag disables a few security checks to make local development easier
- Access the app through the proxy `http://localhost:4000` to interact with the APIs

Try navigate to `http://localhost:4000/ui/registration`

## Go Solution

## Libraries

- **Chi** - router for building Go HTTP service
- **Ory Go Client** - Ory Kratos SDK for Go

## Setup

Install dependencies

```sh
$ go get -u github.com/go-chi/chi/v5
$ go get -u github.com/ory/client-go
```

## Run Project

Run the project

```sh
$ go run .
```

## Self-Hosted

Build docker image from dockerfile and run in it in a container

```sh
$ make up
```

Stop any running containers, remove them and remove the image.

```sh
$ make down
```

