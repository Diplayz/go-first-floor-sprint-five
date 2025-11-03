package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	splitData := strings.Split(data, " ")
	if len(splitData) != 2 {
		return 0, 0, fmt.Errorf("некорректный формат данных")
	}

	stepsStr := splitData[0]
	durationStr := splitData[1]

	if strings.Contains(stepsStr, " ") || strings.Contains(durationStr, " ") {
		return 0, 0, fmt.Errorf("некорректный формат данных")
	}

	intData, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("некорректное количество шагов")
	}

	if intData <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("некорректная длительность")
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("длительность должна быть положительной")
	}

	return intData, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	if duration <= 0 {
		return ""
	}

	distanceKm := (stepLength * float64(steps)) / float64(mInKm)
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKm, spentCalories)
}
