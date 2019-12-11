package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestMonitoringStation(t *testing.T) {
	testCases := []struct {
		In       string
		Detected int
	}{
		{".#..#\n.....\n#####\n....#\n...##", 8},
	}

	for i, c := range testCases {
		_, _, detected := BestMonitoringStation(Load(c.In))
		assert.Equal(t, c.Detected, detected, "(case=%d)", i)
	}
}

func TestAngleToY(t *testing.T) {
	testCases := []struct {
		InX, InY int
		Angle    float64
	}{
		{1, 5, 11.3},
		{5, 5, 45},
		{1, -5, 168.69},
		{0, 5, 0},
		{0, -5, 180},
		{-1, 5, 348.69},
		{-1, -5, 191.31},
	}

	for i, c := range testCases {
		assert.InDelta(t, c.Angle, AngleToY(c.InX, c.InY), 0.01, "(case=%d)", i)
	}
}

func TestShootAsteroids(t *testing.T) {
	testCases := []struct {
		InX, InY int
		InField  string
	}{
		{2, 3, "#####\n#####\n.....\n..#..\n....."},
		{2, 3, ".....\n.....\n.....\n..#..\n#####"},
	}

	for _, c := range testCases {
		ShootAsteroid(c.InX, c.InY, Load(c.InField))
		//assert.InDelta(t, c.Angle, AngleToY(c.InX, c.InY), 0.01, "(case=%d)", i)
	}
}
