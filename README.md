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

The README only includes the basic user guide for the project. For more information, please read the details of our full documentation (bit.ly/APIGatewayDocumentation, "here").


# 1. Basic structure

The system comprises three primary components: the API Gateway, the backend RPC servers, and a registry centre for service discovery. The API Gateway serves as an intermediary, accepting incoming requests from clients, translating and forwarding them to appropriate backend servers. The backend servers, implemented using Kitex, process these requests and return responses back through the Gateway to the clients. The registry centre Etcd is a framework extension from Kitex.


# 2. README Highlight

	**Hertz**: hertz server

	**Idl**:This folder contains the Interface Definition Language (IDL) files used for defining services, methods, and data types in the Thrift protocol. The IDL files play a crucial role in the communication between different microservices.

	**Kitex**: kitex server

	**Tutorial code**: code used previously for the project


# 3. Tech Stack:

* Programming Language: Go (Golang)
* RPC Framework: Kitex
* Service Registry: Etcd (with kitex-contrib/registry-etcd)
* Data Formats: JSON and Thrift



# 4 Usage Instructions:

This guide provides step-by-step instructions on installing and setting up the API gateway for development or use. Users can follow the outlined prerequisites, installation instructions, and initial configuration if necessary.

Ensure that the following is properly installed.


* Go
* Kitex
* Hertz
* etcd


## 4.1: Installing of API gateway




1. Clone this repository:


```
   git clone https://github.com/Chelseacax/orbital-X-Bytedance

```



2. Navigate to the project directory:


```
   cd orbital-X-Bytedance
```



## 4.2: Configuring Go Environment:



* Update your `~/.zshrc` file with the following configuration:
    * `export GOROOT=/usr/local/src/go`
    * `export GOPATH=$HOME/go`
    * `export PATH=$PATH:$GOROOT/bin`
    * `export PATH=$GOPATH/bin:$PATH`
* Run `source ~/.zshrc` to apply the changes to your terminal.


## 4.3: Running etcd service registry

Start service with` etcd `. 


* etcd must be run first before the other servers
* The default IP address of etcd is 127.0.0.2379


## 4.4: Starting the Hertz & Kitex servers

- Start the server using `go run . `under /hertz directory.

- Start the server using `go run . `under /kitex directory.


## 4.5: Testing Cases for.

Open another terminal under, use the curl commands below:



1. For HelloService:
* `curl 'http://127.0.0.1:42000/hello?name=Chelsea'`
2. For CalculatorService:
* Add Method:` curl -X POST -H "Content-Type: application/json" -d '{ "FirstInt": 567, "SecondInt": 123 }' http://127.0.0.1:42000/add`
* Subtract Method:` curl -X POST -H "Content-Type: application/json" -d '{ "FirstInt": 567, "SecondInt": 123 }' http://127.0.0.1:42000/subtract`





