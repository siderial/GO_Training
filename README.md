# GO_Training
ALERT!! ไม่ได้ต่อ Database นะครับ

_______________________________________________________________________________________________________________________
Final Test
สร้าง API ที่มี endpoint ตามต่อไปนี้ เพื่อสำหรับ Booking ห้องประชุม

--------------------------------------------------------
Method	Endpoint	Request Body	Response Body
POST	/bookings/	
--------------------------------------------------------
input 
{"name": "Weerasak Chongnguluam": "room": "ชมพนา", "start": "2019-03-22T07:00:00Z", "end": "2019-03-22T08:00:00Z"}	
output
{"id": 1, "name": "Weerasak Chongnguluam": "room": "ชมพนา", "start": "2019-03-22T07:00:00Z", "end": "2019-03-22T08:00:00Z"}

--------------------------------------------------------
Method	Endpoint	Request Body	Response Body
GET	/bookings/	
--------------------------------------------------------
input 
  - 
output
[{"id": 1, "name": "Weerasak Chongnguluam": "room": "ชมพนา", "start": "2019-03-22T07:00:00Z", "end": "2019-03-22T08:00:00Z"}]

--------------------------------------------------------
Method	Endpoint	Request Body	Response Body
GET	/bookings/1	
--------------------------------------------------------
input 
  - 
output
{"id": 1, "name": "Weerasak Chongnguluam": "room": "ชมพนา", "start": "2019-03-22T07:00:00Z", "end": "2019-03-22T08:00:00Z"}

--------------------------------------------------------
Method	Endpoint	Request Body	Response Body
DELETE	/bookings/1	
--------------------------------------------------------
input 
  - 
output
``
