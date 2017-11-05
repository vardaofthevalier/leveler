## Installing

- First, follow the instructions in the official documentation to bootstrap your Golang development environment
- Install the `glide` package manager for Golang
- Clone this repository to $GOHOME/src/leveler
- Run `glide init` in your Golang workspace to generate a `glide.yml` file
- Run `glide install` in your Golang workspace to install dependencies

## Building

To build the client:

```
cd $GOHOME/src/leveler/client
go build
```

To build the server:

```
cd $GOHOME/src/leveler/server
go build
```

## Running

First, start the server:

```
cd $GOHOME/src/leveler/server
./server &
```

Test that the client can communicate with the server by running a simple command:

```
cd $GOHOME/src/leveler/client
./client create action --name "foo" --description "a foo" --command "echo foo"
```

You should see the action resource object echoed in the server component's logs.

## Project Structure

This project has been designed with the intention of easily supporting modifications or extensions to the underlying gRPC service implementation.

// TODO: write a more thorough explanation

## Future Work