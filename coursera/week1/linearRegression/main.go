package main

import "fmt"

/*
* Created by go-fmt
* 最小二乗法、垂直降下法による線形回帰
* 関数名はcourseraに準拠。あとは英語から適当に推測してください
* "$GOPATH/src/github.com/go-fmt/goTechSalomAICommune"にgit cloneするように作られています
 */

func main() {
	conf := readConfig()
	points := getPoints(conf.inputFileName)
	resultTheta0, resultTheta1 := leastSquare(points, conf.firstTheta0, conf.firstTheta1, conf.alpha, conf.significantFigure)
	diff := sumDifferentSquare(points, resultTheta0, resultTheta1)
	fmt.Printf("y=%f+%f*x\n距離は%fです\n", resultTheta0, resultTheta1, diff)
}
