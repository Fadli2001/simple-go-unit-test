package repository

import (
	"go-unit-test/model"
	"testing"
)

func TestNewCubeRepository(t *testing.T) {
	t.Run("It should return cube repository", func(t *testing.T) {
		cubeModelMockup := model.Cube{
			Side: 2,
		}
		newRepoCube := NewCubeRepository(cubeModelMockup)
		if newRepoCube == nil {
			t.Error("Cannot be nul")
		}
	})
}

func TestCubeRepository_Volume_Success(t *testing.T) {
	t.Run("It should return volume result", func(t *testing.T) {
		cubeModelMockup := model.Cube{
			Side: 4,
		}
		expected := float64(64)
		cubeRepo := NewCubeRepository(cubeModelMockup)
		actual := cubeRepo.Volume()

		if actual != expected {
			t.Error("Volume result is failed")
		}
	})
}

func TestCubeRepository_Area_Success(t *testing.T) {
	t.Run("It should return Area result", func(t *testing.T) {
		cubeModelMockup := model.Cube{
			Side: 4,
		}
		expected := float64(96)
		cubeRepo := NewCubeRepository(cubeModelMockup)
		actual := cubeRepo.Area()

		if actual != expected {
			t.Error("area result is failed")
		}
	})
}

func TestCubeRepository_Perimeter_Success(t *testing.T) {
	t.Run("It should return area result", func(t *testing.T) {
		cubeModelMockup := model.Cube{
			Side: 4,
		}
		expected := float64(48)
		cubeRepo := NewCubeRepository(cubeModelMockup)
		actual := cubeRepo.Perimeter()
		if actual != expected {
			t.Error("Perimeter result is failed")
		}
	})
}
