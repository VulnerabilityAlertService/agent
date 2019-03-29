package config

import (
  "log"
  "github.com/spf13/viper"
  "fmt"
  "os"
  "os/user"
  "bufio"
  "path/filepath"
)

const (
  SAVED_DIR = ".vas/"
  FILE_NAME = "config"
)

func Input() map[string]string {
  var token string
  for token == "" {
    fmt.Print("Enter your token key: ")
    stdin := bufio.NewScanner(os.Stdin)
    stdin.Scan()
    token = stdin.Text()
  }
  return map[string]string{"token": token}
}

func Read() *viper.Viper {
  conf := viper.New()
  conf.SetConfigName(FILE_NAME)
  conf.SetConfigType("yaml")
  conf.AddConfigPath(savedDir())
  err := conf.ReadInConfig()
  if err != nil {
    initialize()
    if err := conf.ReadInConfig(); err != nil {
      log.Fatalln("Read Err: ", err)
    }
  }
  return conf
}

func savedDir() string {
  usr, _ := user.Current()
  return usr.HomeDir + "/" + SAVED_DIR
}

func initialize() {
  filename := filepath.Join(savedDir(), fmt.Sprintf("%s.yml", FILE_NAME))
  if err := os.Mkdir(savedDir(), 0600); err != nil {
    log.Fatalln("Mkdir Err: ", err)
  }
  _, err := os.Create(filename)
  if err != nil {
    log.Fatalln("Create file Err: ", err)
  }
}

func Save(conf *viper.Viper, data map[string]string) {
  conf.Set("token", data["token"])
  conf.WriteConfig()
}

func Token() string {
  token := Read().Get("token")
  if stoken, ok := token.(string); ok {
    return stoken
  }
  return ""
}
