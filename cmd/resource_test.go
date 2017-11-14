package cmd

import (
	"testing"
	mocks "leveler/mocks"
	resources "leveler/resources"
)

func generateCmdConfig() *CmdConfig {
	// read config from resources.json
}

func getResourceClient() *ResourceClient {
	var mock = flag.Bool("mock", false, "Specify whether or not to use mock objects (for unit testing) or real objects (for integration testing)")
	var host = flag.String("host", "127.0.0.1", "The gRPC server host to connect to for integration tests -- 127.0.0.1 is the default")
	var port = flag.Int("port", -1, "The gRPC server port to connect to for integration tests -- no default")
	flag.Parse()

	var r *ResourceClient 

	if mock {
		r = &ResourceClient{
			Client: &mocks.MockResourceEndpointClient{},
			CmdConfig: generateCmdConfig(),
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
				CmdConfig: generateCmdConfig(),
			}
		}
	}

	return r
}

var resourceClient = getResourceClient()

func (t *testing.T) TestResourceClient_Usage() {

}

func (t *testing.T) TestResourceClient_ShortDescription() {
	
}

func (t *testing.T) TestResourceClient_LongDescription() {
	
}

func (t *testing.T) TestResourceClient_AddOptions() {
	
}

func (t *testing.T) TestResourceClient_getId() {
	
}

func (t *testing.T) TestResourceClient_processFlags() {
	
}

func (t *testing.T) TestResourceClient_doGet() {
	
}

func (t *testing.T) TestResourceClient_CreateRequest() {
	
}

func (t *testing.T) TestResourceClient_GetRequest() {
	
}

func (t *testing.T) TestResourceClient_ListRequest() {
	
}

func (t *testing.T) TestResourceClient_UpdateRequest() {
	
}

func (t *testing.T) TestResourceClient_DeleteRequest() {
	
}