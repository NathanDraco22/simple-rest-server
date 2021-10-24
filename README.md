# simple-rest-server
A simple rest server to testing or practices your front-end app

## Endpoints

(localhost:8000) "http://127.0.0.1:8000"

[GET,POST] / 
 -> Simple test GET and POST

[GET] /pathParameter/{data}
 -> Get the value from "data" path Parameter

[GET] /queryParameter
 -> it response with the value of "miQuery" and "miQuery2", both query parameters.

[POST] /postBodyJson
 -> practice the POST method, send a json.
 
 [POST] /postFile 
  -> send a file less than 2mb, the file will put in the sever directory (Multipart Form Request), the file field must be "miFile"
 
