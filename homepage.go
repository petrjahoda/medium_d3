package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"net/http"
	"time"
)

type CalendarInput struct {
	StartDate string
	EndDate   string
}

type CalendarOutput struct {
	Result string
	Data   []CalendarData
}

type CalendarData struct {
	Date            string
	ProductionValue int
}

func homepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logInfo("MAIN", "Serving homepage")
	http.ServeFile(writer, request, "./html/homepage.html")
}

func getCalendarData(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logInfo("MAIN", "Calendar function called")
	var data CalendarInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		logError("MAIN", "Error parsing data: "+err.Error())
		var responseData CalendarOutput
		responseData.Result = "nok: " + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		logInfo("MAIN", "Parsing data ended")
		return
	}
	logInfo("MAIN", "Serving data from "+data.StartDate+" to "+data.EndDate)
	var calendarDummyData []CalendarData
	lastYearStart := time.Now().YearDay() + 365
	for i := 0; i < lastYearStart; i++ {
		var oneCalendarData CalendarData
		oneCalendarData.Date = time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		randomProduction := rand.Intn(100-0) + 0
		oneCalendarData.ProductionValue = randomProduction
		calendarDummyData = append(calendarDummyData, oneCalendarData)
	}
	var responseData CalendarOutput
	responseData.Result = "ok"
	responseData.Data = calendarDummyData
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	logInfo("MAIN", "Parsing data ended")
}
