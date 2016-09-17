package main

import (
	"encoding/json"
	"log"
	"os"
)

/*
* コンフィグファイルを読み込み、pointsを取得する
 */

type config struct {
	InputFileName     string
	Alpha             float64
	SignificantFigure int
	FirstTheta0       float64
	FirstTheta1       float64
}

type point struct {
	x float64
	y float64
}

func getPoints(inputFileName string) (points []point) {
	file, err := os.Open(inputFileName)
	defer file.Close()
	if err != nil {
		log.Fatal("サンプルファイルが開けませんでした", err)
	}
	dec := json.NewDecoder(file)
	err = dec.Decode(&points)
	if err != nil {
		log.Fatal(err)
	}
	return points
}

func readConfig() (conf config) {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatal("コンフィグファイルが開けませんでした", err)
	}
	dec := json.NewDecoder(file)
	err = dec.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}
