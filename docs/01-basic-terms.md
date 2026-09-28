# Handler
Piece of Go code that receives an HTTP request and produces an HTTP response.  

# Router/ServeMux
Decides which handler should run.  
For example: `/about` -> `aboutHandler`

# Web Server
Listens for incoming HTTP requests, and passes them into the application (eventually, into the 
routers and handlers)  
