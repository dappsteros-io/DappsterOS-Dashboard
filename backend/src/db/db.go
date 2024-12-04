package db

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Conn struct {
	/* dependencies */
	URI string `json:"uri" bson:"uri"`
}

func (c *Conn) Init() error {
	// Setup the mgm default config
	mongoURI := (map[bool]string{true: "mongodb://localhost:27017/dappsteros", false: "mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/dappsteros"})[os.Getenv("MONGO_HOST") == ""]

	if mongoURI == "" {
		log.Fatal("DB_ADDR environment variable not specified")
	}
	log.Println(mongoURI)

	err := mgm.SetDefaultConfig(nil, "dappsteros", options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
