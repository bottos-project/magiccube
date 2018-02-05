package api

import (
	"fmt"
	proto "../proto"
	"github.com/micro/go-micro/client"
	//"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"github.com/micro/go-web"
	"net/http"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	service.HandleFunc("/register", register)
	service.HandleFunc("/login", login)


	if err := service.Init(); err != nil {
		fmt.Print(err)
	}

	if err := service.Run(); err != nil {
		fmt.Print(err)
	}
}
func login(writer http.ResponseWriter, request *http.Request) {
	// Use the generated client stub
	cl := proto.NewUserClient("go.micro.srv.user", client.DefaultClient)

	// Make request
	rsp, err := cl.Login(context.Background(), &proto.LoginRequest{
		Username: request.FormValue("userName"),
		Signature: request.FormValue("signature"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(writer, rsp.Msg, rsp.Token)
}
func register(writer http.ResponseWriter, request *http.Request) {
	// Use the generated client stub
	cl := proto.NewUserClient("go.micro.srv.user", client.DefaultClient)

	// Make request
	rsp, err := cl.Register(context.Background(), &proto.RegisterRequest{
		Username: request.FormValue("userName"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(writer, rsp.Msg)
}
