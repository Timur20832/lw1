package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Header []headerdata
	menu   []menudata
	bigp   []bigpdata
	smallp []smallpdata
	//footer []footerdata
}

// start header
type headerdata struct {
	Escape string
	Nav    []navdata
}

type navdata struct {
	First  string
	Second string
	Third  string
	Fourth string
}

// end header

// start menu
type menudata struct {
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
}

// end menu

// start bigpdata
type bigpdata struct {
	TitlePost    string
	InfoPostOneB []dataInfoPostOneB
	InfoPostTwoB []dataInfoPostTwoB
}

type dataInfoPostOneB struct {
	TitlePostOne string
	TextPostOne  string
	ImagePerson  string
	PersonName   string
	PersonDate   string
}

type dataInfoPostTwoB struct {
	TitlePostTwo string
	TextPostTwo  string
	ImagePerson  string
	PersonName   string
	PersonDate   string
}

// end bigpdata

// start smallpost
type smallpdata struct {
	TitlePost      string
	InfoPostOneS   []dataInfoPostOneS
	InfoPostTwoS   []dataInfoPostTwoS
	InfoPostThreeS []dataInfoPostThreeS
	InfoPostFourS  []dataInfoPostFourS
	InfoPostFiveS  []dataInfoPostFiveS
	InfoPostSixS   []dataInfoPostSixS
}

type dataInfoPostOneS struct {
	ImgPostOneS   string
	TitlePostOneS string
	TextPostOne   string
	ImagePerson   string
	PersonName    string
	PersonDate    string
}

type dataInfoPostTwoS struct {
	ImgPostTwoS   string
	TitlePostTwoS string
	TextPostTwo   string
	ImagePerson   string
	PersonName    string
	PersonDate    string
}

type dataInfoPostThreeS struct {
	ImgPostThreeS   string
	TitlePostThreeS string
	TextPostThreeS  string
	ImagePerson     string
	PersonName      string
	PersonDate      string
}

type dataInfoPostFourS struct {
	ImgPostFourS   string
	TitlePostFourS string
	TextPostFourS  string
	ImagePerson    string
	PersonName     string
	PersonDate     string
}

type dataInfoPostFiveS struct {
	ImgPostFiveS   string
	TitlePostFiveS string
	TextPostFiveS  string
	ImagePerson    string
	PersonName     string
	PersonDate     string
}

type dataInfoPostSixS struct {
	ImgPostSixS   string
	TitlePostSixS string
	TextPostSixS  string
	ImagePerson   string
	PersonName    string
	PersonDate    string
}

//end small posts

// start footer
/*type footerdata struct {
	FormFound    []dataFormFound
	NavbarFooter []dataNavbarFooter
}

type dataFormFound struct {
}

type dataNavbarFooter struct {
}*/

// end footer
func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Header: header(),
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Mega Error", 500)
		log.Println(err.Error())
		return
	}
}

func header() []headerdata {
	return []headerdata{
		{
			Escape: "static/image/Escape..svg",
			Nav:    nav(),
		},
	}
}

func nav() []navdata {
	return []navdata{
		{
			First:  "HOME",
			Second: "CATEGORIES",
			Third:  "ABOUT",
			Fourth: "CONTACT",
		},
	}
}
