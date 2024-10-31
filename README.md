# 🚀 GOTH Stack Complete Setup

[![Go Version](https://img.shields.io/badge/Go-1.23.2-00ADD8?style=flat-square&logo=go)](https://golang.org/doc/go1.23)
[![HTMX](https://img.shields.io/badge/HTMX-1.9.11-purple?style=flat-square)](https://htmx.org)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-3.4.1-38B2AC?style=flat-square&logo=tailwind-css)](https://tailwindcss.com)
[![Alpine.js](https://img.shields.io/badge/Alpine.js-3.13.7-2D3441?style=flat-square)](https://alpinejs.dev)
[![Templ](https://img.shields.io/badge/Templ-0.2.778-red?style=flat-square)](https://templ.guide)

A modern, fully configured starter template for building fast, type-safe web applications using Go, HTMX, Tailwind CSS, Alpine.js, and Templ. This stack provides an incredible developer experience with type safety from your database all the way to your HTML.

## ✨ Features

- **📜 Type-Safe Templates**: Using [Templ](https://templ.guide) for compile-time checked templates
- **🔥 Hot Reload**: Using [Air](https://github.com/cosmtrek/air) for instant feedback during development
- **🎨 HTMX + Tailwind**: [Modern, interactive UIs](https://htmx.org) without complex client-side JavaScript
- **🛠️ Modern JS Utilities**: [Alpine.js](https://alpinejs.dev) for lightweight interactivity
- **📱 Responsive Design**: [Mobile-first approach](https://tailwindcss.com) with Tailwind CSS

## 🚀 Quick Start

### Prerequisites

- Go v1.23.2 or higher
- npm v10.2.3
- Air v1.61.1
- Templ CLI v0.2.778

### Installation

Install dependencies:
```bash
npm install
```

### Running Locally

Start the development server:
```bash
make run
```

This will:
- Start the Templ proxy server on the value of TEMPL_PROXY_PORT defined in the .env file.
- Start the Go server on the value of APP_PORT, also defined in the .env file.
- Enable hot reloading for all file changes.

Access the application at: 
```bash
http://localhost:<TEMPL_PROXY_PORT>
```

### Running as Container

1. Build the container:
```bash
docker build -t <image-tag> .
```

2. Run the container:
```bash
docker run --rm \
  --env-file .env \
  -p <local-port>:<app-port> \
  -t <image-tag>
```

Access the containerized application at:
```bash
http://localhost:<local-port>
```

Note: The `<app-port>` should match the `APP_PORT` in your `.env` file.

## In detail explanation

There is an in detail explanation of how the hot reloading works in this article.
