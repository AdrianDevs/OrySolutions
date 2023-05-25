<h1>Ory Self-Service Solution</h1>





<h2>Table of Contents</h2>

[TOC]

# Introduction

A scaffolded project for Web Authentication and Authorisation

# Key Technologies

- **Ory Oathkeeper**
- **Ory Kratos**
- **Ory Keto**
- **Ory Hydra**
- **Go**
- **Chi**
- **SvelteKit**
- **Tailwind**

# Services

- **OathKeeper** for a Proxy service
- **Kratos** for Identity Management
- **Kratos DB** for a PostgreSQL Database for data storage
- **Mailhog** for an Email service
- **Web API Go** for an API service.
- Kratos Self-Service Go for showing UI for user Authentication
- Keto
- Hydra
- Svelte Web App as a Web App
- Kratos Self-Service Svelte for showing UI for user Authentication
- Admin Go
- Admin Svelte
- Traefik
- Prometheus
- Grafana

# Ports

| Service    | External Ports | Internal Ports | Description        |
| ---------- | -------------- | -------------- | ------------------ |
| Oathkeepr  | 8080           | 4455           | Proxy port         |
|            | --             | 4456           | API decisions port |
| Kratos     | --             | 4433           | Public API         |
|            | --             | 4434           | Private Admin API  |
| Kratos DB  | --             | 5432           |                    |
| Mailhog    | 8025           | 8025           | HTTP server        |
|            | 1025           | 1025           | SMTP server        |
| Web API Go | --             | 8090           |                    |
|            |                |                |                    |
|            |                |                |                    |
|            |                |                |                    |

# Running

Clone repo

```sh
$ git clone 
```

Run taskfile

```sh
$ task up
```

