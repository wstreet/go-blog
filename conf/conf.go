package conf

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbPort     string
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
)

func Init() {
	// TODO: read conf.init file
	HttpPort = ":8080"
}
