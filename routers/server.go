package routers

// Init : start server at 8080 port
func Init() {
	r := NewRouter()
	r.Run(":8000")
}
