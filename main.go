package main

import (
	   "net/http"
	   "SipToolKit/urls"
	   )

func main()  {

	urls.LoadURLS()
	http.ListenAndServe(":80", nil)
	
}