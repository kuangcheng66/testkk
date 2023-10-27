package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"io"
	"os"
)

func main() {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// docker images
	/*images, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", images)*/

	// docker pull image
	/*reader, err := cli.ImagePull(ctx, "ghcr.io/kosmos-io/clustertree-knode-manager:v0.1.7", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)*/

	// docker save -o *.tar.gz images
	/*path := "/root/testkk/kosmos2.tar.gz"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	saveResponse, err := cli.ImageSave(ctx, []string{"14ff280f8e75"})
	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(file, saveResponse); err != nil {
		panic(err)
	}*/

	// docker load -i *.tar.gz
	tar, err := os.Open("/root/clusterlink-operator-v0.1.5.tar")
	imageLoadResponse, err := cli.ImageLoad(ctx, tar, true)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(imageLoadResponse.Body)
	if err != nil {
		panic(err)
	}
	str := string(body)
	fmt.Printf("%s", string(str))

	// docker push image
	/*	authConfig := types.AuthConfig{
			Username: "PaaS_cnpt_test_suwanliang2022", //harbor用户名 （这一块我是读的配置文件，需要的可以进行修改）
			Password: "Minami@toI1",                   //harbor 密码
		}
		encodedJSON, err := json.Marshal(authConfig)
		if err != nil {
			panic(err)
		}
		authStr := base64.URLEncoding.EncodeToString(encodedJSON)

		image := "clusterlink-eb203842.ecis-suzhou-1.cmecloud.cn/kosmos-io/clusterlink/clusterlink-network-manager:0.2.0"
		out, err := cli.ImagePush(ctx, image, types.ImagePushOptions{RegistryAuth: authStr})
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(out)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", string(body))*/

	// docker tag image1 image2
	/*err = cli.ImageTag(ctx, "ghcr.io/kosmos-io/clustertree-knode-manager:v0.1.7", "ghcr.io/kosmos-io/clustertree-knode-manager:vkk")
	if err != nil {
		panic(err)
	}*/
}
