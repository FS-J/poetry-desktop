package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"strings"
	"math/rand"
	"time"
	"strconv"
	"os"
	"net/http"
	"io"
	"runtime"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fetch_poetry(url string) {

	file_name_tokens := strings.Split(url, "/")
	file_name := "json/" + file_name_tokens[len(file_name_tokens)-1]

	// Create json dir if it doesn't exist
	if _, err := os.Stat("json"); os.IsNotExist(err) {
		os.Mkdir("json", 0755)
	}
	
	// Check file existence first with io.IsExist
	if _, err := os.Stat(file_name); os.IsNotExist(err) {

		fmt.Println("下载中", url, "保存至", file_name)
		
		output, err := os.Create(file_name)
		if err != nil {
			fmt.Println("创建文件时出现错误", file_name, "-", err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("下载文件时出现错误", url, "-", err)

			// Delete created json file when download fails
			os.Remove(file_name)
			
			return
		}
		defer response.Body.Close()

		n, err := io.Copy(output, response.Body)
		if err != nil {
			fmt.Println("下载文件时出现错误", url, "-", err)
			os.Remove(file_name)
			return
		}
		fmt.Println("下载完成：", (n / 1000), " KB")
	} else {
		fmt.Println("\n已在本地发现诗词json文件，路径：", file_name)
	}
}

func main() {

	type Poetry struct {
		Author string
		Paragraphs []string
		Rhythmic string
		Title string
	}

	poet_song_base_url := "https://raw.githubusercontent.com/chinese-poetry/chinese-poetry/master/json/poet.song."
	songci_base_url := "https://raw.githubusercontent.com/chinese-poetry/chinese-poetry/master/ci/ci.song."
	poet_tang_base_url := "https://raw.githubusercontent.com/chinese-poetry/chinese-poetry/master/json/poet.tang."
	
	var str_bldr strings.Builder

	// Generate random number for choosing from poet or Songci
	rand_seed := rand.NewSource(time.Now().UnixNano())
	rand_0_1 := rand.New(rand_seed).Intn(2)

	// Choosing from poet or Songci
	if rand_0_1 == 0 {
	
		// Generate random number for choosing from poet song or poet tang
		rand_seed = rand.NewSource(time.Now().UnixNano())
		rand_0_1 = rand.New(rand_seed).Intn(2)
		if rand_0_1 == 0 {
			str_bldr.WriteString(poet_song_base_url)

			// Get a random number in {0, 1000, 2000, ..., 254000} for choosing a random Song poet
			rand_0_254000 := rand.New(rand_seed).Intn(255) * 1000

			// Choose a random Song poet
			str_bldr.WriteString(strconv.Itoa(rand_0_254000))
			str_bldr.WriteString(".json")
		} else {
			str_bldr.WriteString(poet_tang_base_url)

			// Get a random number in {0, 1000, 2000, ..., 57000} for choosing a random Tang poet
			rand_0_57000 := rand.New(rand_seed).Intn(58) * 1000

			// Choose a random Tang poet
			str_bldr.WriteString(strconv.Itoa(rand_0_57000))
			str_bldr.WriteString(".json")
		}
	} else {
		str_bldr.WriteString(songci_base_url)

		// Get a random number in {0, 1000, 2000, ..., 21000} for choosing a random Song ci
		rand_0_21000 := rand.New(rand_seed).Intn(22) * 1000

		// Choose a random Song ci
		str_bldr.WriteString(strconv.Itoa(rand_0_21000))
		str_bldr.WriteString(".json")
	}

	json_url := str_bldr.String()
	fetch_poetry(json_url)
	
	// Reset string builder
	str_bldr.Reset()

	json_url_splited := strings.Split(json_url, "/")
	json_loc := "json/" + json_url_splited[len(json_url_splited) - 1]
	json_data, err := ioutil.ReadFile(json_loc)
	check(err)

	byte_json_data := []byte(string(json_data))
	var json_slice []Poetry
	//var json_slice_ci []Songci

	err = json.Unmarshal(byte_json_data, &json_slice)
	check(err)

	// Generate random number for poetry and songci picking
	rand_0_999 := rand.New(rand_seed).Intn(1000)

	// Re_rand it when there are not enough elements in a last json slice
	for rand_0_999 >= len(json_slice) {
		rand_0_999 = rand.New(rand_seed).Intn(1000)
	}
	rand_poetry := json_slice[rand_0_999]

	// Construct poetry string
	str_bldr.WriteString("\n")

	// Read from Songci-Rhythmic or poetry-Title
	str_bldr.WriteString(rand_poetry.Title)
	str_bldr.WriteString(rand_poetry.Rhythmic)
	str_bldr.WriteString(" - ")
	str_bldr.WriteString(rand_poetry.Author)
	str_bldr.WriteString("\n")
	for ord := range rand_poetry.Paragraphs {
		str_bldr.WriteString("\n")
		str_bldr.WriteString(rand_poetry.Paragraphs[ord])
		str_bldr.WriteString("\n")
	}
	poetry_string := str_bldr.String()

	// Reset string builder
	str_bldr.Reset()

	// Print poetry to terminal
	fmt.Println(poetry_string)

	if runtime.GOOS == "windows" {
		
		convCmdStr := "images/poetry_bg_0.png -font fonts/NotoSansCJK-Regular.ttc -gravity center -fill rgba(0,0,0,0.8) -pointsize 25 -annotate +0+0 images/out_bg_0.jpg"
		convCmdSlice := strings.Split(convCmdStr, " ")

		// Insert poetry string to convCmdSlice before the last element
		convCmdSlice = append(convCmdSlice, "")
		copy(convCmdSlice[(len(convCmdSlice) - 1):], convCmdSlice[(len(convCmdSlice) - 2):])
		convCmdSlice[len(convCmdSlice) - 2] = poetry_string

		// Magic, miracle at "convCmdSlice..." !
		convCmd := exec.Command("bin/convert.exe", convCmdSlice...)
		_, err := convCmd.Output()
		check(err)

		// Set wallpaper use a binary
		setWCmd := exec.Command("bin/wallpaper.exe", "images/out_bg_0.jpg")
		_, err = setWCmd.Output()
		check(err)

		
	} else if runtime.GOOS == "linux" {
		
		convCmdStr := "images/poetry_bg_0.png -font fonts/NotoSansCJK-Regular.ttc -gravity center -pointsize 25 -annotate +0+0 images/out_bg_0.jpg"
		convCmdSlice := strings.Split(convCmdStr, " ")

		// Insert poetry string to convCmdSlice before the last element
		convCmdSlice = append(convCmdSlice, "")
		copy(convCmdSlice[(len(convCmdSlice) - 1):], convCmdSlice[(len(convCmdSlice) - 2):])
		convCmdSlice[len(convCmdSlice) - 2] = poetry_string

		// Magic, miracle at "convCmdSlice..."
		convCmd := exec.Command("convert", convCmdSlice...)
		_, err := convCmd.Output()
		check(err)
	}
}
