package main

import(
	"time"
	"strings"
	"net/http"
	"strconv"
	"errors"
	"sort"
	
	"github.com/gin-gonic/gin"
)

var BookingMap = make(map[string]Booking)
var IdRuning int

type Booking struct {
	Id    string `bson:"id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Room  string `bson:"room" json:"room"`
	Start time.Time   `bson:"start" json:"start"`
	End   time.Time   `bson:"end" json:"end"`
}

type BookingStr struct {
	Name  string `bson:"name" json:"name"`
	Room  string `bson:"room" json:"room"`
	Start string   `bson:"start" json:"start"`
	End   string   `bson:"end" json:"end"`
}

func genarateRunning() int{
	IdRuning ++
		return IdRuning
}

func convertTime(dateStr string) time.Time{
	dateStr = strings.Replace(dateStr, "Z", "+07:00", 1)
	t, _ := time.Parse(time.RFC3339, dateStr)
	return t
}

func getBookingList() []Booking {
	bookingList := []Booking{}
	for _, booking := range BookingMap {
		bookingList = append(bookingList,booking)
	}
	return bookingList
}

func getSortBookingList() []Booking {
	bookingList := getBookingList()
	sort.SliceStable(bookingList, func(i, j int) bool {
		return bookingList[i].Start.Before(bookingList[j].Start)
	})
	return bookingList
}

func main(){
	r := gin.Default()
	r.POST("/bookings", func(c *gin.Context) {
		booking := Booking{
		}
		bookingStr := BookingStr{}
		if err := c.Bind(&bookingStr); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		booking.Id = strconv.Itoa(genarateRunning())
		booking.Name = bookingStr.Name
		booking.Room = bookingStr.Room
		booking.Start = convertTime(bookingStr.Start)
		booking.End = convertTime(bookingStr.End)
		
		BookingMap[booking.Id] = booking
		c.JSON(http.StatusOK,BookingMap)

	})

	r.GET("/bookings", func(c *gin.Context) {
		bookingList  :=  getSortBookingList()
		c.JSON(http.StatusOK,bookingList)
	})

	r.GET("/bookings/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, ok := BookingMap[id]
		if(!ok){
			c.AbortWithError(http.StatusNotFound,errors.New("id not found") )
			return
		}
		booking := BookingMap[id]
		c.JSON(http.StatusOK,booking)
	})

	r.DELETE("/bookings/:id", func(c *gin.Context) {
		id := c.Param("id")
	
		delete(BookingMap, id)
		c.JSON(http.StatusOK,``)
	})

	r.Run(":8000")

}


