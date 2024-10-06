package git

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// MakeRandomCommits делает случайное количество коммитов
func MakeRandomCommits() {

	rand.Seed(time.Now().UnixNano())
	commitsCount := rand.Intn(5)
	fmt.Printf("Сегодня будет сделано %d коммитов\n", commitsCount)

	for i := 0; i < commitsCount; i++ {
		// Создаем изменения в файле
		file, err := os.Create("random_file.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Пишем случайные данные
		_, err = file.WriteString(fmt.Sprintf("Random data %d\n", rand.Int()))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Добавляем файл в индекс
		execCommand("git", "add", "random_file.txt")

		// Делаем коммит
		execCommand("git", "commit", "-m", fmt.Sprintf("Random commit %d", i+1))

		// Пушим изменения на GitHub
		execCommand("git", "push")
	}

	fmt.Println("Все коммиты выполнены.")
}

// execCommand выполняет команду git
func execCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Ошибка выполнения команды %s: %v\n", name, err)
	}
}
