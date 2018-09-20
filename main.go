package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

	log.Printf("Loaded configuration: %+v", configuration)

	for _, chunks := range configuration.Chunks {
		log.Printf("Processing %+v, with values %+v", chunks.Filename, chunks)

		// Clean slate
		_ = os.Remove(configuration.Outdir + "/" + chunks.Filename)

		// Concatenate stuff without mercy
		for _, chunk := range chunks.Parts {
			infile, err := ioutil.ReadFile(configuration.Indir + "/" + chunk)
			if err != nil {
				panic(err)
			}

			outfile, err := os.OpenFile(configuration.Outdir+"/"+chunks.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
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
