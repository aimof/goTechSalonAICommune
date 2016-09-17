package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/*
* コンフィグファイルを読み込み、pointsを取得する
 */

type config struct {
	inputFileName     string  `json:"inputFileName"`
	alpha             float64 `json:"alpha"`
	significantFigure int     `json:"significantFigure"`
	firstTheta0       float64 `json:"firstTheta0"`
	firstTheta1       float64 `json:"firstTheta1"`
}

type point struct {
	x float64 `json:"x"`
	y float64 `json:"y"`
}

func getPoints(inputFileName string) (points []point) {
	file, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		log.Fatal("サンプルファイルが開けませんでした")
	}
	err = json.Unmarshal(file, &points)
	if err != nil {
		log.Fatal("サンプルファイルの形式に誤りがあります。")
	}
	return points
}

func readConfig() (conf config) {
	file, err := ioutil.ReadFile("$GOPATH/src/github.com/go-fmt/goTechSalonAICommune/coursera/week1/linearRegression/config.json")
	if err != nil {
		log.Fatal("コンフィグファイルが開けませんでした", err)
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatal("コンフィグファイル形式が合致しません")
	}
	return conf
}
