# GO Server

## Overview

This is a simple Go server project. It serves as a backend application and can be expanded for various use cases.

## Features

- Lightweight and efficient Go server

- REST API support

- Easy to configure and extend

## Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (latest stable version recommended)

- [Git](https://git-scm.com/) (optional, for version control)

## Installation

1. Clone the repository:

```
git clone ParthKansara98/GO-SERVER
```

2. Navigate to the project directory:

```
cd GO-SERVER
```

## Running the Server

To start the server, run:

```
go run main.go
```

## Building the Project

To build the project into a binary:

```
go build -o server
```

Then, run the binary:

```
./server
```

## API Endpoints

1. ```/``` Is end point to ```index.html``` file.
2. ```/form.html``` Is end point for ```form.html``` file.
3. ```/form``` Is end point for display result of form when we click on ```Print Values```.

## File Structure

```
GO SERVER
├── static/
│   ├── form.html
│   ├── index.html
└── main.go
```
