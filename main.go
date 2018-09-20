package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Invoke as: %v config.yml\n", os.Args[0])
		os.Exit(1)
	}

	viper.SetConfigFile(os.Args[1])
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	for filename, values := range configuration.Values {
		log.Printf("Processing %+v, with values %+v", filename, values)
		svalues := strings.Split(values, " ")

		// Clean slate
		_ = os.Remove(configuration.Outdir + "/" + filename)

		// Concatenate stuff without mercy
		for _, chunk := range svalues {
			infile, err := ioutil.ReadFile(configuration.Chunks + "/" + chunk)
			if err != nil {
				panic(err)
			}

			outfile, err := os.OpenFile(configuration.Outdir+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}

			defer outfile.Close()

			if _, err = outfile.WriteString(string(infile)); err != nil {
				panic(err)
			}
		}
	}
}
