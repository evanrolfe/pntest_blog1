package main

// reference: https://github.com/onaci/docker-plugin-seaweedfs/blob/golang/Makefile
//------------------------------------------------------------------------------
// build:
//------------------------------------------------------------------------------
// docker build -t my-plugin-image .
// docker run my-plugin-image
//------------------------------------------------------------------------------
// rootfs:
//------------------------------------------------------------------------------
// docker build -q -t my-plugin-image:rootfs .
// mkdir rootfs
// docker create --name tmp my-plugin-image:rootfs
// docker export tmp | tar -x -C ./rootfs
// docker rm -vf tmp
//------------------------------------------------------------------------------
// create plugin:
//------------------------------------------------------------------------------
// docker plugin rm -f my-plugin-image:latest || true
// docker plugin create my-plugin-image:latest .
// docker plugin enable my-plugin-image:latest
/*
type MyPlugin struct{}

func (d *MyPlugin) CreateEndpoint(request *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	return nil, nil
}

func (d *MyPlugin) CreateNetwork(r *network.CreateNetworkRequest) error {
	fmt.Printf("Hello!\n")
	return nil
}

func (d *MyPlugin) DeleteNetwork(r *network.DeleteNetworkRequest) error {
	fmt.Printf("Goodbye!\n")
	return nil
}

func (d *MyPlugin) DeleteEndpoint(request *network.DeleteEndpointRequest) error {
	return nil
}

func (d *MyPlugin) Join(r *network.JoinRequest) (*network.JoinResponse, error) {
	fmt.Printf("Container %s joined network\n", r.NetworkID)
	return &network.JoinResponse{}, nil
}

func (d *MyPlugin) Leave(r *network.LeaveRequest) error {
	fmt.Printf("Container %s left network\n", r.NetworkID)
	return nil
}

func (d *MyPlugin) DiscoverNew(r *network.DiscoveryNotification) error {
	fmt.Println("Discovery notification:", r.DiscoveryType)
	return nil
}

func (d *MyPlugin) DiscoverDelete(r *network.DiscoveryNotification) error {
	fmt.Println("Discovery notification:", r.DiscoveryType)
	return nil
}

func (d *MyPlugin) ProgramExternalConnectivity(r *network.ProgramExternalConnectivityRequest) error {
	fmt.Println("ProgramExternalConnectivity:", r.EndpointID)
	return nil
}

func (d *MyPlugin) RevokeExternalConnectivity(r *network.RevokeExternalConnectivityRequest) error {
	fmt.Println("RevokeExternalConnectivity:", r.EndpointID)
	return nil
}

func (d *MyPlugin) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {

	return nil, nil
}

func (d *MyPlugin) EndpointInfo(request *network.InfoRequest) (*network.InfoResponse, error) {
	return nil, nil
}

func (d *MyPlugin) FreeNetwork(*network.FreeNetworkRequest) error {
	return nil
}

func (d *MyPlugin) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return &network.CapabilitiesResponse{
		Scope: network.LocalScope,
	}, nil
}

func main() {
	fmt.Println("my-plugin starting...")
	driver := &MyPlugin{}
	handler := network.NewHandler(driver)
	err := handler.ServeUnix("test", 0)
	if err != nil {
		panic(err)
	}
}
*/
