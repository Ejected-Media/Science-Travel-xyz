# Science-Travel-xyz ~

package main



import (
    
    
)


func app_welcome_center_page() {
    
    
}





//  .  html url routes 
//  .  as well as terminal cli logs

func main() {
// ,  ° . +
  appName := "~ www.Science-Travel.xyz - // - Website App"

// ,  ° . +
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/front", indexHandler)
  http.HandleFunc("/main", indexHandler)
  http.HandleFunc("/home", indexHandler)
  http.HandleFunc("/start", indexHandler)
  
  // ,  ° . +
  http.HandleFunc("/user", pageHandler)
  http.HandleFunc("/account", pageHandler)
  http.HandleFunc("/profile", pageHandler)
  
  http.HandleFunc("/portfolio", pageHandler)
  http.HandleFunc("/resume", pageHandler)
  
  
  // ,  ° . +
  http.HandleFunc("/settings", appHandler)
  




// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
// ,  ° . +
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }