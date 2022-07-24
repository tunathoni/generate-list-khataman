package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/nleeper/goment"
)

type listKhataman []string
type dauroh string

func readFromFile(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error read file : ", err)
		os.Exit(1)
	}

	return string(content)
}

func readList(filename string) listKhataman {
	listData := readFromFile(filename)

	list := strings.Split(string(listData), ",")

	return listKhataman(list)
}

func getDauroh() dauroh {
	data := readFromFile("dauroh")

	newDauroh, err := strconv.Atoi(data)

	if err != nil {
		fmt.Println("Error read file : ", err)
		os.Exit(1)
	}

	newDauroh = newDauroh + 1

	return dauroh(strconv.Itoa(newDauroh))
}

func (list listKhataman) setNewListOrder() listKhataman {
	maxData := len(list) - 1
	firstData := list[maxData]
	remainlist := list[:maxData]

	newList := append(listKhataman{firstData}, remainlist...)

	return newList
}

func (list listKhataman) print() {
	startDate, _ := goment.New()
	endDate, _ := goment.New()
	dateFormat := "DD MMMM YYYY"
	start := startDate.Format(dateFormat)
	end := endDate.Add(6, "days").Format(dateFormat)

	startDay := getDayName(startDate.Day())
	endDay := getDayName(endDate.Add(7, "days").Day())

	fmt.Println("*Bismillahirrahmanirrahim*")
	fmt.Println("LIST Khotmil Qur'an *Bani Yunus Ubaid*")

	dauroh := "DAUROH : *" + getDauroh() + "*"

	fmt.Println(dauroh)
	fmt.Println(start, "-", end, "\n")

	for index, v := range list {
		index += 1
		listNumber := ""
		if index < 10 {
			listNumber += "📖0" + strconv.Itoa(index) + "."
		} else {
			listNumber += "📖" + strconv.Itoa(index) + "."
		}

		fmt.Println(listNumber, v)
	}

	fmt.Println()
	fmt.Println("NOTE : \nTanda Kholas = ✅ \n\n*START TILAWAH* hari " + startDay + " " + start + " pukul *18:00* (Bakdal Maghrib) setelah dibacakan FATICHAH \n\nTetap SEMANGAT n ISTIQOMAH TILAWAH sebelum tidur \n\n*Untuk yang kosong, jika ingin mbadali, dipersilahkan dan diinfokan terlebih dahulu ke grup. \n\n*Do'a Khotmil Quran dibaca bersama-sama di tempat masing2* pada " + endDay + " malam " + startDay + " " + end + " pukul *17:00* (Qoblal Maghrib) \n\nSemoga Istiqomah, menambah berkah  hidup kita dan keluarga. \n\nآمِيْنْ اَللَّهُمَّ آمِيْنْ.. \n💚📖📖📖📖💚")
}

func (list listKhataman) toString() string {
	return strings.Join(list, ",")
}

func getDayName(day int) string {
	dayData := []string{"AHAD", "SENIN", "SELASA", "RABU", "KAMIS", "JUM'AH", "SABTU"}

	return dayData[day]
}

func (list listKhataman) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(list.toString()), 0666)
}

func (d dauroh) saveDaurohToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d), 0666)
}
