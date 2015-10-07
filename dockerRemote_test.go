package collector

import (
	"fmt"
	"testing"
)

func TestListDanglingImages(t *testing.T) {
	fmt.Println("ListDanglingImages")
	var err error
	DockerTransport, err = NewDockerTransport(DOCKERPROTO, DOCKERADDR)
	if err != nil {
		t.Fatal(err)
	}
	imageList, err := ListDanglingImages()
	if err != nil {
		t.Fatal(err)
	}
	for _, image := range imageList {
		fmt.Printf("image ID %s\n", string(image))
	}
}

func TestRemoveImageByID(t *testing.T) {
	var err error
	DockerTransport, err = NewDockerTransport(DOCKERPROTO, DOCKERADDR)
	if err != nil {
		t.Fatal(err)
	}
	RegistrySpec = "index.docker.io"
	RegistryAPIURL, HubAPI, BasicAuth, XRegistryAuth = GetRegistryURL()
	metadata := ImageMetadataInfo{
		Repo: "banyanops/nginx",
		Tag:  "1.7",
	}
	fmt.Println("TestPullImage %v", metadata)
	PullImage(metadata)

	id := "02a791aafe15"
	resp, err := RemoveImageByID(ImageIDType(id))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(resp))
}
