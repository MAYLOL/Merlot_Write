package rpc
import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "merlot_project/merlot_write/service"
	"merlot_project/merlot_write/service"
)
//Article
var ClientArticle  service.ArticleServiceClient
var ClientArticleCtx  context.Context
var ClientError error


//topic
var ClientTopic service.TopicServiceClient
var ClientTopicCtx context.Context


//diary
var ClientDiary service.DiaryServiceClient
var ClientDiaryCtx context.Context
const (
	Address2     = "localhost:50051"
)

func init(){
//1.创建文章RPC的连接
StartArticleRPC()
//2.创建日记RPC的连接
StartDiaryRPC()
//3.创建主题RPC的连接
StartTopicRPC()
}
//FUNC:创建文章RPC服务连接
func StartArticleRPC(){
	conn, err := grpc.Dial(Address2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		print("did not connect")
	}
	ClientArticle = pb.NewArticleServiceClient(conn)
	ClientArticleCtx = context.Background()
}
//FUNC:创建主题RPC服务连接
func StartTopicRPC(){
	// Set up a connection to the server.
	conn, err := grpc.Dial(Address2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ClientTopic = pb.NewTopicServiceClient(conn)
	ClientTopicCtx = context.Background()
}
//FUNC：创建日记RPC的服务链接
func StartDiaryRPC(){
	// Set up a connection to the server.
	conn, err := grpc.Dial(Address2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ClientDiary = pb.NewDiaryServiceClient(conn)
	ClientDiaryCtx = context.Background()
}