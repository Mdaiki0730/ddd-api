package config

import (
  "log"
  "os"

  "github.com/joho/godotenv"
)

type Env struct {
  // port config
  GRPCPort string
  RESTPort string
}

var Global Env = Env{}

func Load() {
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("Error loading .env file: %v \n", err)
  }

  // port config
  Global.GRPCPort = os.Getenv("GRPC_PORT")
  Global.RESTPort = os.Getenv("REST_PORT")
}
