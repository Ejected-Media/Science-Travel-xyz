# Science-Travel-xyz ~

package main



import (
    
    
)


func app_welcome_center_page() {
    
    
}




// . appHandler, ~ for Data Pages °
func appHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/app" {
    	http.NotFound(w, r)
    	return
    }
    
// ,
  pageTitle := "~ www.Science-Travel.xyz - // - Website App"
  pagePath := r.URL.Path
  
  pageType := ".."
  
// ,  ° . +
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  if pagePath == "/settings" {
      pageTitle = "Settings Page"
      pageList = pageList
  }
  
  // ,  ° . +
  if pagePath == "/inventory" {
      pageTitle = "Inventory Page"
      pageList = pageList
  }
  // ,  ° . +
  if pagePath == "/itinerary" {
      pageTitle = "Itinerary Page"
      pageList = pageList
  }
  // ,  ° . +
  if pagePath == "/quest_list" {
      pageTitle = "Quest List Page"
      pageList = pageList
  }
  // ,  ° . +
  if pagePath == "/travel_log" {
      pageTitle = "Travel Log Page"
      pageList = pageList
  }
  
  
// ,  ° . + 
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  appHandler




// . pageHandler, ~ for User Pages °
func pageHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +

  pageTitle := "~ www.Science-Travel.xyz - // - Website App"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  if pagePath == "/account" {
      pageTitle = "Account Page"
      pageList = pageList
  }
  
  if pagePath == "/profile" {
      pageTitle = "Profile Page"
      pageList = pageList
  }
  
// ,  ° . +
  if pagePath == "/portfolio" {
      pageTitle = "Portfolio Page"
      pageList = pageList
  }
  
  if pagePath == "/resume" {
      pageTitle = "Resume Page"
      pageList = pageList
  }


// ,  ° . +
pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  pageHandler



// . indexHandler,  ~ for Public Pages °
func indexHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "~ www.Science-Travel.xyz - // - Website App"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  if pagePath == "/" {
      pageTitle = "Index Page"
      pageList = pageList
  }
  
    if pagePath == "/front" {
      pageTitle = "Front Page"
      pageList = pageList
  }
  
    if pagePath == "/main" {
      pageTitle = "Main Page"
      pageList = pageList
  }
  
    if pagePath == "/home" {
      pageTitle = "Home Page"
      pageList = pageList
  }
  
    if pagePath == "/start" {
      pageTitle = "Start Page"
      pageList = pageList
  }

// ,  ° . +
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler



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
  
    // ,  ° . +
  http.HandleFunc("/inventory", pageHandler)
    // ,  ° . +
  http.HandleFunc("/itinerary", pageHandler)
    // ,  ° . +
  http.HandleFunc("/quest_list", pageHandler)
    // ,  ° . +
  http.HandleFunc("/travel_log", pageHandler)
    // ,  ° . +
  http.HandleFunc("/quest_needs", pageHandler)
    // ,  ° . +
  http.HandleFunc("/quest_people", pageHandler)
  




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