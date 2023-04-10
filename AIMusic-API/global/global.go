package global

import (
	"AIMusic-API/config"
	"AIMusic-API/proto"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

var (
	Config = &config.AIMusicConfig{}
	MQCh   *amqp.Channel
	RDB    *redis.Client
	mqConn *amqp.Connection

	AdministratorInstances []ServiceInstance
	CommunityInstances     []ServiceInstance
	CreationInstances      []ServiceInstance
	UserInstances          []ServiceInstance

	AdministratorClients []proto.AdministratorClient
	CommunityClients     []proto.CommunityClient
	CreationClients      []proto.CreationClient
	UserClients          []proto.UserClient
)

type ServiceInstance struct {
	Address string
	Port    int
}

func Init() {
	logInit()
	configureInit()
	mqInit()
	redisInit()
}

func logInit() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}

func configureInit() {
	v := viper.New()
	v.SetConfigFile("config.yml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(Config); err != nil {
		panic(err)
	}
}

func mqInit() {
	var err error
	//连接Rabbitmq服务 协议://用户名:密码@主机IP:端口
	mqConn, err = amqp.Dial("amqp://root:AIMusic@localhost:5672")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打开通道
	MQCh, err = mqConn.Channel()

	// 创建一个队列，存在则不创建
	_, err = MQCh.QueueDeclare(
		"scratch", // 定义队列的名字
		false,     // 队列是否持久化保存到硬盘。队列里面的数据是否保存得取决于发布者发布消息时的设置
		false,     // 没有消费者使用时是否删除
		false,     // 是否排他性
		false,     // 是否无需等待
		nil,       // 其他参数
	)

	// 创建一个队列，存在则不创建
	_, err = MQCh.QueueDeclare(
		"accompany", // 定义队列的名字
		false,       // 队列是否持久化保存到硬盘。队列里面的数据是否保存得取决于发布者发布消息时的设置
		false,       // 没有消费者使用时是否删除
		false,       // 是否排他性
		false,       // 是否无需等待
		nil,         // 其他参数
	)
}

func redisInit() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port),
		Password: Config.Redis.Password, // 没有密码，默认值
		DB:       0,                     // 默认DB 0
	})
}

func UpdateServiceInstances() {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d",
		Config.Consul.Host, Config.Consul.Port)

	client, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("Error creating Consul client:", err)
		return
	}

	administrator, _, err := client.Catalog().Service("AIMusic-Administrator", "", nil)
	if err != nil {
		fmt.Println("Error getting service instances:", err)
		return
	}
	administratorInstances := make([]ServiceInstance, len(administrator))
	for i, service := range administrator {
		administratorInstances[i] = ServiceInstance{
			Address: service.ServiceAddress,
			Port:    service.ServicePort,
		}
	}
	community, _, err := client.Catalog().Service("AIMusic-Community", "", nil)
	if err != nil {
		fmt.Println("Error getting service instances:", err)
		return
	}
	communityInstances := make([]ServiceInstance, len(community))
	for i, service := range community {
		communityInstances[i] = ServiceInstance{
			Address: service.ServiceAddress,
			Port:    service.ServicePort,
		}
	}
	creation, _, err := client.Catalog().Service("AIMusic-Creation", "", nil)
	if err != nil {
		fmt.Println("Error getting service instances:", err)
		return
	}
	creationInstances := make([]ServiceInstance, len(creation))
	for i, service := range creation {
		creationInstances[i] = ServiceInstance{
			Address: service.ServiceAddress,
			Port:    service.ServicePort,
		}
	}
	user, _, err := client.Catalog().Service("AIMusic-User", "", nil)
	if err != nil {
		fmt.Println("Error getting service instances:", err)
		return
	}
	userInstances := make([]ServiceInstance, len(user))
	for i, service := range user {
		userInstances[i] = ServiceInstance{
			Address: service.ServiceAddress,
			Port:    service.ServicePort,
		}
	}

	AdministratorInstances = administratorInstances
	CommunityInstances = communityInstances
	CreationInstances = creationInstances
	UserInstances = userInstances
	fmt.Println(AdministratorInstances)
	fmt.Println(CommunityInstances)
	fmt.Println(CreationInstances)
	fmt.Println(UserInstances)
}

func UpdateGRPCConn() {
	AdministratorClients = make([]proto.AdministratorClient, 8)
	CommunityClients = make([]proto.CommunityClient, 8)
	CreationClients = make([]proto.CreationClient, 8)
	UserClients = make([]proto.UserClient, 8)

	for i := 0; i < len(AdministratorInstances); i++ {
		administratorConn, _ := grpc.Dial(fmt.Sprintf("%s:%d", AdministratorInstances[i].Address, AdministratorInstances[i].Port), grpc.WithInsecure())
		AdministratorClients[i] = proto.NewAdministratorClient(administratorConn)
	}

	for i := 0; i < len(CommunityInstances); i++ {
		communityConn, _ := grpc.Dial(fmt.Sprintf("%s:%d", CommunityInstances[i].Address, CommunityInstances[i].Port), grpc.WithInsecure())
		CommunityClients[i] = proto.NewCommunityClient(communityConn)
	}

	for i := 0; i < len(CreationInstances); i++ {
		creationConn, _ := grpc.Dial(fmt.Sprintf("%s:%d", CreationInstances[i].Address, CreationInstances[i].Port), grpc.WithInsecure())
		CreationClients[i] = proto.NewCreationClient(creationConn)
	}

	for i := 0; i < len(UserInstances); i++ {
		userConn, _ := grpc.Dial(fmt.Sprintf("%s:%d", UserInstances[i].Address, UserInstances[i].Port), grpc.WithInsecure())
		UserClients[i] = proto.NewUserClient(userConn)
	}
}

func LoadBalancing(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
