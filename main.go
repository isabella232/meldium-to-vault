package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/extemporalgenome/slug"
	"github.com/fatih/structs"
	"github.com/hashicorp/vault/api"
)

type Secret struct {
	Key         string
	DisplayName string
	UserName    string
	Password    string
	Notes       string
	Url         string
}

func main() {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("err: %s", err)
	}

	if token := client.Token(); token != "" {
		log.Printf("using token client %s", token)
	} else {
		log.Fatal("no VAULT_TOKEN supplied!")
	}

	secrets := readMeldiumCSV()
	for _, secret := range secrets {
		path := fmt.Sprintf("meldium/%s", secret.Key)
		_, err = client.Logical().Write(path, structs.Map(secret))
		if err != nil {
			log.Fatal("error writing %s: %s", path, err)
		} else {
			log.Printf("wrote %s successfully", path)
		}
	}
}

func readMeldiumCSV() []Secret {
	f, _ := os.Open("./meldium.csv")

	r := csv.NewReader(bufio.NewReader(f))
	i := 0

	var secrets []Secret
	for {
		i++

		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if i == 1 {
			continue
		}

		secret := Secret{}
		for _i, value := range record {
			switch _i {
			case 0:
				secret.Key = slug.Slug(value)
				secret.DisplayName = value
			case 1:
				secret.UserName = value
			case 2:
				secret.Password = value
			case 3:
				secret.Notes = value
			case 4:
				secret.Url = value
			}
		}

		secrets = append(secrets, secret)
	}

	return secrets
}
