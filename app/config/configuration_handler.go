package config

import (
	"github.com/go-pg/pg"
	"github.com/gorilla/schema"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"ic-service/app/config/locales/local_config"
	"ic-service/kafka_producers"

	"log"
	"os"
	"strings"
	"time"
)

var (
	config                 *Config
	requestParamsValidator *validator.Validate
	requestParamsDecoder   *schema.Decoder
	session                *pg.DB
	syncProducer           kafka_producers.SyncProducer
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Kafka    KafkaConfiguration
}

func GetSyncProducer() kafka_producers.SyncProducer {
	return syncProducer
}

type ServerConfig struct {
	Port string
}

type KafkaConfiguration struct {
	Brokers             []string
	IceCreamCreateTopic string
}

func GetDBConnection() *pg.DB {
	return session
}

type PostgresConfig struct {
	Host           string
	Username       string
	Password       string
	Database       string
	MigrationsPath string
	PoolSize       int
	DialTimeout    int
	ReadTimeout    int
	WriteTimeout   int
}

type SectorConfig struct {
	Id        string
	Companies []CompanyConfig
}

type CompanyConfig struct {
	Name     string
	Response string
}

func GetConfig() *Config {
	return config
}

func SetDummyConfig() {
	config = &Config{}
}

func InitializeKafkaSyncProducer() {
	kConfig := kafka_producers.KafkaConfig{}
	kConfig.Brokers = GetConfig().Kafka.Brokers
	producer, err := kafka_producers.NewSyncProducer(kConfig)
	if err != nil {
		log.Fatal("SERVER_STARTUP : Error during initializing syncProducer: %s", err)
	}
	log.Print("Initialized kafka connection")
	syncProducer = producer
}

func InitializeDb() {
	log.Print("Initializing db connection")
	session = pg.Connect(&pg.Options{
		User:         config.Postgres.Username,
		Password:     config.Postgres.Password,
		Database:     config.Postgres.Database,
		Addr:         config.Postgres.Host,
		PoolSize:     config.Postgres.PoolSize,
		DialTimeout:  time.Duration(config.Postgres.DialTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(config.Postgres.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.Postgres.WriteTimeout) * time.Millisecond,
	})
	var num int
	_, err := session.Query(pg.Scan(&num), "SELECT ?", 42)
	log.Print("Initializing postgres connection")
	if err != nil {
		log.Fatal(nil, "Error with Postgres connections")
	}
	log.Print("Upgrading db schema")
	_ = "postgres://" + GetConfig().Postgres.Username + ":" + GetConfig().Postgres.Password +
		"@" + GetConfig().Postgres.Host + "/" + GetConfig().Postgres.Database + "?sslmode=disable"

	log.Print("Initialized db connection successfully")
}

func Initialize() error {
	viper.SetConfigFile(getConfigFile())
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, os.Getenv((strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}"))))
		}
	}
	return viper.Unmarshal(&config)
}

func InitializeTestConfigurationHandler() {
	viper.SetConfigFile("../config/dev.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("SERVER_STARTUP : Error reading config file, %s", err)
	}

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, os.Getenv((strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}"))))
		}
	}

	jsonMarshalErr := viper.Unmarshal(&config)
	if jsonMarshalErr != nil {
		log.Fatalf("SERVER_STARTUP : Error unmarshalling config file, %s", jsonMarshalErr)
	}

}

func InitializeTestConfig() {

	InitializeTestConfigurationHandler()
	InitializeDecoderAndValidator()
	InitializeDb()
	InitializeKafkaSyncProducer()

}

func InitializeConfig() {

	err := Initialize()
	local_config.InitializeLocalizer()
	InitializeDecoderAndValidator()
	InitializeDb()
	InitializeKafkaSyncProducer()

	if err != nil {
		log.Fatal(nil, "error initializing Config :", err)
		return
	}

}

func GetReqParamsValidator() *validator.Validate {
	return requestParamsValidator
}

func GetReqParamsDecoder() *schema.Decoder {
	return requestParamsDecoder
}

func getConfigFile() string {
	return "app/config/" + os.Getenv("ENV") + ".yml"
}

func InitializeDecoderAndValidator() {
	log.Print("Initializing params decoder & validator")
	requestParamsDecoder = schema.NewDecoder()
	requestParamsValidator = validator.New()
	requestParamsDecoder.IgnoreUnknownKeys(true)
}
