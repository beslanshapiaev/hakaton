package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const token = "660c38115eea2660c38115eea6"

// тут не будет ручного ввода, сделано для теста
func main() {
	for {
		fmt.Print("Введите команду (move, collect, reset, rounds, quit): ")
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения команды:", err)
			continue
		}
		command = strings.TrimSpace(strings.ToLower(command))
		switch command {
		case "move":
			moveToNextPlanet()
		case "collect":
			collectGarbage()
		case "reset":
			sendResetRequest()
		case "rounds":
			getRoundsData()
		case "quit":
			fmt.Println("Выход из программы...")
			return
		default:
			fmt.Println("Неверная команда")
		}
		time.Sleep(1 * time.Second)
	}
}

// тут будет реализован механизм перемещения, пока тоже тест
func moveToNextPlanet() {
	universeData, err := getUniverseData()
	if err != nil {
		fmt.Println("Ошибка получения данных о вселенной:", err)
		return
	}
	fmt.Println(universeData)

	request := TravelRequest{
		Planets: []string{universeData.Universe[0][1].(string), universeData.Universe[1][1].(string), universeData.Universe[2][1].(string)},
	}

	response, err := sendTravelRequest(request)
	if err != nil {
		fmt.Println("Ошибка отправки запроса на перемещение:", err)
		return
	}
	fmt.Printf("Расход топлива: %d\n", response.FuelDiff)
}

// тут будет реализован механизм для сборки
func collectGarbage() {
	universeData, err := getUniverseData()
	if err != nil {
		fmt.Println("Ошибка получения данных о вселенной:", err)
		return
	}
	fmt.Println(universeData)
	request := CollectRequest{
		Garbage: map[string][][]int{
			"Garbage1": {{0, 0}, {0, 1}, {1, 1}},
			"Garbage2": {{2, 2}, {2, 3}, {3, 3}},
		},
	}

	response, err := sendCollectRequest(request)
	if err != nil {
		fmt.Println("Ошибка отправки запроса на сбор мусора:", err)
		return
	}
	fmt.Println(response)
}

func getUniverseData() (UniverseResponse, error) {
	url := "https://datsedenspace.datsteam.dev/player/universe"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return UniverseResponse{}, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("X-Auth-Token", token)
	resp, err := client.Do(req)
	if err != nil {
		return UniverseResponse{}, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	var data UniverseResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return UniverseResponse{}, fmt.Errorf("ошибка при декодировании JSON: %v", err)
	}
	return data, nil
}

func sendTravelRequest(request TravelRequest) (TravelResponse, error) {
	url := "https://datsedenspace.datsteam.dev/player/travel"

	requestBody, err := json.Marshal(request)
	if err != nil {
		return TravelResponse{}, fmt.Errorf("ошибка при кодировании запроса: %v", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return TravelResponse{}, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return TravelResponse{}, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	var response TravelResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return TravelResponse{}, fmt.Errorf("ошибка при декодировании JSON: %v", err)
	}

	return response, nil
}

func sendCollectRequest(request CollectRequest) (CollectResponse, error) {
	url := "https://datsedenspace.datsteam.dev/player/collect"

	requestBody, err := json.Marshal(request)
	if err != nil {
		return CollectResponse{}, fmt.Errorf("ошибка при кодировании запроса: %v", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return CollectResponse{}, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return CollectResponse{}, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	var response CollectResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return CollectResponse{}, fmt.Errorf("ошибка при декодировании JSON: %v", err)
	}

	return response, nil
}

func sendResetRequest() error {
	url := "https://datsedenspace.datsteam.dev/player/reset"

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("неправильный статус код ответа: %d", resp.StatusCode)
	}

	return nil
}

func getRoundsData() (RoundsResponse, error) {
	url := "https://datsedenspace.datsteam.dev/player/rounds"

	// Создаем HTTP клиент
	client := &http.Client{}

	// Создаем GET запрос
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RoundsResponse{}, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	// Добавляем заголовок с API ключом
	req.Header.Set("X-Auth-Token", token)

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		return RoundsResponse{}, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	// Декодируем JSON ответ
	var response RoundsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return RoundsResponse{}, fmt.Errorf("ошибка при декодировании JSON: %v", err)
	}

	return response, nil
}
