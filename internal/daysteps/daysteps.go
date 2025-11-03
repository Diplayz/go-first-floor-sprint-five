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
	splitData := strings.Split(data, ",") // ИЗМЕНИТЬ на запятую
	if len(splitData) != 2 {
		return 0, 0, fmt.Errorf("некорректный формат данных")
	}

	intData, err := strconv.Atoi(strings.TrimSpace(splitData[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("некорректное количество шагов")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(splitData[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("некорректная длительность")
	}

	return intData, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
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
