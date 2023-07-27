<!-----

Yay, no errors, warnings, or alerts!

Conversion time: 0.5 seconds.


Using this Markdown file:

1. Paste this output into your source file.
2. See the notes and action items below regarding this conversion run.
3. Check the rendered output (headings, lists, code blocks, tables) for proper
   formatting and use a linkchecker before you publish this page.

Conversion notes:

* Docs to Markdown version 1.0β34
* Sun Jul 23 2023 19:32:22 GMT-0700 (PDT)
* Source doc: readme
----->

## check out the details of our full documentation here:https://docs.google.com/document/d/1_plI0Uo0BMp_dmGk1guOZPVVn63LVT6k3Y7NAu9aBss/edit#heading=h.s0ri7fwxlsfu

## 1. Basic structure

The system comprises three primary components: the API Gateway, the backend RPC servers, and a registry centre for service discovery. The API Gateway serves as an intermediary, accepting incoming requests from clients, translating and forwarding them to appropriate backend servers. The backend servers, implemented using Kitex, process these requests and return responses back through the Gateway to the clients. The registry centre Etcd is a framework extension from Kitex.


## 2. README Highlight

	**Hertz**: hertz server

	**Idl**:This folder contains the Interface Definition Language (IDL) files used for defining services, methods, and data types in the Thrift protocol. The IDL files play a crucial role in the communication between different microservices.

	**Kitex**: kitex server

	**Tutorial** **code**: code used previously for the project


## 3. Tech Stack:



* Programming Language: Go (Golang)
* RPC Framework: Kitex
* Service Registry: Etcd (with kitex-contrib/registry-etcd)
* Data Formats: JSON and Thrift


## 4.Usage Examples

This guide provides step-by-step instructions on installing and setting up the API gateway for development or use. Users can follow the outlined prerequisites, installation instructions, and initial configuration if necessary.


### 4.1 Installation

1. Clone this repository:


```
   git clone https://github.com/Chelseacax/orbital-X-Bytedance
```


2. Navigate to the project directory:


```
   cd orbital-X-Bytedance
```


3. Run the application:


```
   go run .
```



### 4.2 Go environment configuration:

   - Update your `~/.zshrc` file with the following configuration:


```
     export GOROOT=/usr/local/src/go
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOROOT/bin
     export PATH=$GOPATH/bin:$PATH
```


   - Run `source ~/.zshrc` to apply the changes to your terminal.

   - Use `go env` to verify the settings and ensure proper configuration.


### 4.3 Commands according to endpoints:

To handle different operations in your API gateway, you can use specific endpoints for each operation. Currently we have implemented two services: addService and helloService and you can use them as the end points.


### 4.4 Make rpc call

To set up the Hz server:

- Generate Hz code scaffolding using the `hz new` command with the appropriate IDL file.

- Generate Kitex client code scaffolding using the `kitex` command and the relevant IDL file.

- Start the server using ``go run .`.`

To use Kitex server:

- Generate Kitex server code scaffolding with the ``kitex`` command and the corresponding IDL file.

- Complete the business logic in the `handler.go` file.

- Start the server using ``go run .`.`


