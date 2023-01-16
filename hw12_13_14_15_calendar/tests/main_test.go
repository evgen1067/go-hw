package scripts

import (
	"github.com/cucumber/godog"
	"os"
	"testing"
	"time"
)

const delay = 5 * time.Second

func TestMain(m *testing.M) {
	//log.Printf("wait %s for service availability...", delay)
	//time.Sleep(delay)

	status := godog.TestSuite{
		Name:                "integration",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:    "pretty", // Замените на "pretty" для лучшего вывода
			Paths:     []string{"features"},
			Randomize: 0, // Последовательный порядок исполнения
		},
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
