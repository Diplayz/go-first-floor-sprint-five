package daysteps

import (
	"fmt"
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
		return 0, 0, fmt.Errorf("")
	}
	intData, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("")
	}
	if intData <= 0 {
		return 0, 0, fmt.Errorf("")

	}
	duration, err := time.ParseDuration(splitData[1])
	if err != nil {
		return 0, 0, fmt.Errorf("")
	}

	return intData, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}

	distanceKm := (stepLength * float64(steps)) / float64(mInKm)
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция: %.2f Км.\nВы сожгли %.2f калл.\n", steps, distanceKm, spentCalories)
}
