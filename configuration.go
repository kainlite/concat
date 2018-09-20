package main

type Chunks struct {
	Filename string
	Parts    []string
}

type Configuration struct {
	Outdir string
	Indir  string
	Chunks map[string]Chunks
}
