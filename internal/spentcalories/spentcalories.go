package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	splitData := strings.Split(data, ",")
	if len(splitData) != 3 {
		return 0, "", 0, fmt.Errorf("некорректный формат данных")
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка при преобразовании количества шагов: %v", err)
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	activity := splitData[1]

	duration, err := time.ParseDuration(splitData[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка при преобразовании строки в duration: %v", err)
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("длительность должна быть положительной")
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	stepDistance := (float64(steps) * stepLength) / float64(mInKm)
	return stepDistance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}
	dist := distance(steps, height)
	hours := duration.Hours()

	return dist / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64
	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), calories)
	return info, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные параметры")
	}
	speed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	runningSpent := (weight * speed * durationMin) / minInH
	return runningSpent, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные параметры")
	}
	speed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	walkingSpend := (weight * speed * durationMin) / minInH
	walkingSpend = walkingSpend * walkingCaloriesCoefficient
	return walkingSpend, nil
}
