package main

import (
	server "CVBackend"
	"CVBackend/pkg/handler"
	"CVBackend/pkg/repository"
	"CVBackend/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/siruspen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error config %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error env var %s", err)
	}
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error init db %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	//err = repo.Default.InDefaultContent()
	//if err != nil {
	//	logrus.Fatalf("error init default %s", err)
	//}

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouter()); err != nil {
		logrus.Fatalf("error run server %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
