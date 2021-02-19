package utilities

//LoginURL is the URL for logging into churchofjesuschrist.org
//const LoginURL = "https://signin.churchofjesuschrist.org/login.html"
const LoginURL = "https://www.churchofjesuschrist.org/?lang=eng&signmein"

//const LoginURL = "https://login.churchofjesuschrist.org/api/authenticate/credentials"

// MemberListURL is the URL for getting the member list from LCR
const MemberListURL = "https://lcr.churchofjesuschrist.org/services/umlu/report/member-list"

// UserFile is the json file where users are stored
const UserFile = "users.json"

// Households is the json file where all household members are stored
const Households = "data/household-list.json"

// Roads contain all of the roads in the Foxwood Ward
var Roads = [11]string{"levi ln", "lexi loop", "kaylee ct", "addison ave", "grace ct", "jakes ln", "campbell cir", "foxwood dr", "cottage ln", "hutch ln", "wilson way"}

// PriesthoodOffices is the list of offices in the Priesthood
var PriesthoodOffices = [3]string{"ELDER", "PRIEST", "TEACHER"}

// GoogleMapsBeginURL is the url to hit for calculating the distance
var GoogleMapsBeginURL = "https://maps.googleapis.com/maps/api/distancematrix/json?origins="

// GoogleMapsEndURL is the url to hit for calculating the distance
var GoogleMapsEndURL = "&destinations=place_id:ChIJdaozlXd_TYcRclj5L5TGp2w&key=AIzaSyCBicX9DgiOaSGcs0uf7nJlr0-MvnKnyRo"

// APIkey is the api key for the google distance matrix
var APIkey = "AIzaSyCBicX9DgiOaSGcs0uf7nJlr0-MvnKnyRo"
