package cmd

import (
	"os"
	"fmt"
	"flag"
	"bytes"
	"testing"
	"google.golang.org/grpc"
	mocks "leveler/mocks"
	resources "leveler/resources"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

var testConfig = `{
	"Name": "action",
    "Usage": "action",
    "ShortDescription": "Perform an operation on an Action resource",
    "LongDescription": "Defined as a command and shell which can be later applied to a destination resource",
    "SupportedOperations": ["create", "get", "list", "update", "patch", "delete", "apply"],
    "Spec": {
       "StringOptions": [
          {
             "Name": "name",
             "Required": true,
             "Description": "A descriptive name for the action",
             "Default": "", 
             "IsSecondaryKey": true
          },
          {
             "Name": "description",
             "Required": false,
             "Description": "A concise description for the goal of the action",
             "Default": ""
          },
          {
             "Name": "command",
             "Required": true,
             "Description": "A shell command that achieves the action",
             "Default": ""
          },
          {
             "Name": "shell",
             "Required": true,
             "Default": "/bin/bash",
             "Description": "A concise description for the goal of the action"
          }
        ]
    }
}
`
func generateCmdConfig() *resources.CmdConfig {
	// read config from resources.json

	var m = &resources.CmdConfig{}
	var jsonUnmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: false,
	}

	reader := bytes.NewReader([]byte(testConfig))
	err := jsonUnmarshaler.Unmarshal(reader, m)
	if err != nil {
		fmt.Printf("Error processing resource configuration: %v", err)
		os.Exit(1)
	}

	return m
}

func getResourceClient() *ResourceClient {
	var mock = flag.Bool("mock", false, "Specify whether or not to use mock objects (for unit testing) or real objects (for integration testing)")
	var host = flag.String("host", "127.0.0.1", "The gRPC server host to connect to for integration tests -- 127.0.0.1 is the default")
	var port = flag.Int("port", -1, "The gRPC server port to connect to for integration tests -- no default")
	flag.Parse()

	var r *ResourceClient 

	if *mock {
		r = &ResourceClient{
			Client: &mocks.MockResourceEndpointClient{},
			CmdConfig: *generateCmdConfig(),
		}
	} else {
		if *port == -1 {
			fmt.Println("gRPC server port not specified -- please rerun with the '-port' option!")
			os.Exit(1)
		} else {
			opts = append(opts, grpc.WithInsecure())  // TODO: if further config is necessary add options to command line
			clientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), opts...)

			if err != nil {
				fmt.Printf("Couldn't connect to server: %s", err)
				os.Exit(1)
			}

			r = &ResourceClient{
				Client: resources.NewResourceEndpointClient(clientConn), 
				CmdConfig: *generateCmdConfig(),
			}
		}
	}

	return r
}

var resourceClient = getResourceClient()

func TestResourceClient_Usage(t *testing.T) {
	if resourceClient.Usage() != "action" {
		t.Errorf("Usage() returned the wrong value!")
	}
}

func TestResourceClient_ShortDescription(t *testing.T)  {
	if resourceClient.ShortDescription() != "Perform an operation on an Action resource" {
		t.Errorf("ShortDescription() returned the wrong value!")
	}
}

func TestResourceClient_LongDescription(t *testing.T) {
	if resourceClient.LongDescription() != "Defined as a command and shell which can be later applied to a destination resource" {
		t.Errorf("LongDescription() returned the wrong value!")
	}
}

func TestResourceClient_AddOptions(t *testing.T) {
	
}

func TestResourceClient_getId(t *testing.T) {
	
}

func TestResourceClient_processFlags(t *testing.T) {
	
}

func TestResourceClient_doGet(t *testing.T) {
	
}

func TestResourceClient_CreateRequest(t *testing.T) {
	
}

func TestResourceClient_GetRequest(t *testing.T) {
	
}

func TestResourceClient_ListRequest(t *testing.T) {
	
}

func TestResourceClient_UpdateRequest(t *testing.T) {
	
}

func TestResourceClient_DeleteRequest(t *testing.T) {
	
}