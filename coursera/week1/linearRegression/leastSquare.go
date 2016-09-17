package main

import "math"

/*
* 最小二乗法
 */

func leastSquare(points []point, theta0, theta1, alpha float64, significantFigure int) (newTheta0, newTheta1 float64) {
	x, y, m := pointsToXY(points)
	for end := false; !end; {
		newTheta0, newTheta1, end = steepestDescentOnce(x, y, theta0, theta1, alpha, m, significantFigure)
		theta0, theta1 = newTheta0, newTheta1
	}
	return newTheta0, newTheta1
}

func pointsToXY(points []point) (x, y []float64, m int) {
	m = len(points)
	for i := 0; i < m; i++ {
		x = append(x, points[i].x)
		y = append(y, points[i].y)
	}
	return x, y, m
}

func steepestDescentOnce(x, y []float64, theta0, theta1, alpha float64, m, significantFigure int) (newTheta0, newTheta1 float64, end bool) {
	newTheta0 = theta0 - alpha*derivativeTheta0(x, y, theta0, theta1, m)
	newTheta1 = theta1 - alpha*derivativeTheta1(x, y, theta0, theta1, m)
	if round(newTheta0, significantFigure) == round(theta0, significantFigure) && round(newTheta1, significantFigure) == round(theta1, significantFigure) {
		return newTheta0, newTheta1, true
	}
	return newTheta0, newTheta1, false
}

func h(xi, theta0, theta1 float64) float64 {
	return theta0 + theta1*xi
}

func derivativeTheta0(x, y []float64, theta0, theta1 float64, m int) float64 {
	sum := 0.0
	for i := 0; i < m; i++ {
		sum += h(x[i], theta0, theta1) - y[i]
	}
	return sum
}

func derivativeTheta1(x, y []float64, theta0, theta1 float64, m int) float64 {
	sum := 0.0
	for i := 0; i < m; i++ {
		sum += (h(x[i], theta0, theta1) - y[i]) * x[i]
	}
	return sum
}

func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

/*
* 距離の二乗の和を求める
 */

func differentSquare(xi, yi, theta0, theta1 float64) float64 {
	return math.Exp2(yi - h(xi, theta0, theta1))
}

func sumDifferentSquare(points []point, theta0, theta1 float64) float64 {
	x, y, m := pointsToXY(points)
	var sum float64
	for i := 0; i < m; i++ {
		sum += differentSquare(x[i], y[i], theta0, theta1)
	}
	return sum
}
