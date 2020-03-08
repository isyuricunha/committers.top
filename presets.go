package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
	"austria":            PresetLocations{"austria", "österreich", "vienna", "wien", "linz", "salzburg", "graz", "innsbruck", "klagenfurt", "wels", "dornbirn"},
	"finland":            PresetLocations{"finland", "suomi", "helsinki", "tampere", "oulu", "espoo", "vantaa", "turku", "rovaniemi", "jyväskylä", "lahti", "kuopio", "pori", "lappeenranta", "vaasa"},
	"sweden":             PresetLocations{"sweden", "sverige", "stockholm", "malmö", "uppsala", "göteborg", "gothenburg"},
	"norway":             PresetLocations{"norway", "norge", "oslo", "bergen"},
	"germany":            PresetLocations{"germany", "deutschland", "berlin", "frankfurt", "munich", "münchen", "hamburg", "cologne", "köln"},
	"netherlands":        PresetLocations{"netherlands", "nederland", "amsterdam", "rotterdam", "hague", "utrecht", "holland", "delft"},
	"ukraine":            PresetLocations{"ukraine", "kiev", "kyiv", "kharkiv", "dnipro", "odesa", "donetsk", "zaporizhia"},
	"japan":              PresetLocations{"japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
	"russia":             PresetLocations{"russia", "moscow", "saint+petersburg", "novosibirsk", "yekaterinburg", "nizhny+novgorod", "samara", "omsk", "kazan", "chelyabinsk", "rostov-on-don", "ufa", "volgograd"},
	"estonia":            PresetLocations{"estonia", "eesti", "tallinn", "tartu", "narva", "pärnu"},
	"denmark":            PresetLocations{"denmark", "danmark", "copenhagen", "aarhus", "odense", "aalborg"},
	"portugal":           PresetLocations{"portugal", "lisbon", "lisboa", "braga", "porto", "aveiro", "coimbra", "funchal", "madeira"},
	"france":             PresetLocations{"france", "paris", "marseille", "lyon", "toulouse", "nice", "nantes", "strasbourg", "montpellier", "bordeaux", "lille", "rennes", "reims", "rouen", "toulon", "le+havre", "grenoble", "dijon", "le+mans", "brest,france", "tours"},
	"spain":              PresetLocations{"spain", "españa", "madrid", "barcelona", "valencia", "seville", "sevilla", "zaragoza", "malaga", "murcia", "palma", "bilbao", "alicante", "cordoba"},
	"italy":              PresetLocations{"italy", "italia", "rome", "roma", "milan", "naples", "napoli", "turin", "torino", "palermo", "genoa", "genova", "bologna", "florence", "firenze", "bari", "catania", "venice", "verona"},
	"uk":                 PresetLocations{"uk", "england", "scotland", "wales", "northern+ireland", "london", "birmingham", "leeds", "glasgow", "sheffield", "bradford", "manchester", "edinburgh", "liverpool", "bristol", "cardiff", "belfast", "leicester", "wakefield", "coventry", "nottingham", "newcastle"},
	"croatia":            PresetLocations{"croatia", "hrvatska", "zagreb", "split", "rijeka", "osijek", "zadar", "pula"},
	"worldwide":          PresetLocations{},
	"china":              PresetLocations{"china", "中国", "guangzhou", "shanghai", "beijing", "hangzhou", "hong+kong"},
	"india":              PresetLocations{"india", "mumbai", "delhi", "bangalore", "hyderabad", "ahmedabad", "chennai", "kolkata", "jaipur"},
	"israel":             PresetLocations{"israel", "tel+aviv", "jerusalem", "beer+sheva", "beersheva", "netanya", "ramat+gan", "haifa", "herzliya", "rishon"},
	"indonesia":          PresetLocations{"indonesia", "jakarta", "surabaya", "bandung", "medan", "bekasi", "semarang", "tangerang", "depok", "makassar", "palembang"},
	"pakistan":           PresetLocations{"pakistan", "karachi", "lahore", "faisalabad", "rawalpindi", "peshawar", "islamabad"},
	"brazil":             PresetLocations{"brazil", "brasil", "são+paulo", "brasília", "salvador", "fortaleza", "belo+horizonte", "manaus", "curitiba", "recife", "porto+alegre", "florianópolis"},
	"nigeria":            PresetLocations{"nigeria", "lagos", "kano", "ibadan", "benin+city", "port+harcourt", "jos", "ilorin"},
	"bangladesh":         PresetLocations{"bangladesh", "dhaka", "chittagong", "khulna", "rajshahi", "barisal", "sylhet", "rangpur", "comilla", "gazipur"},
	"mexico":             PresetLocations{"mexico", "mexico+city", "guadalajara", "puebla", "tijuana", "mexicali", "monterrey", "hermosillo", "zapopan", "ciudad+juarez", "chihuahua", "aguascalientes", "mx"},
	"philippines":        PresetLocations{"philippines", "pilipinas", "quezon", "manila", "davao", "caloocan", "cebu", "zamboanga", "bohol", "pasig", "bacolod", "makati", "baguio", "cavite"},
	"luxembourg":         PresetLocations{"luxembourg", "esch-sur-alzette", "differdange", "dudelange", "ettelbruck", "diekirch", "wiltz", "echternach", "rumelange", "grevenmacher", "bertrange", "mamer", "capellen", "strassen", "diekirch"},
	"egypt":              PresetLocations{"egypt", "cairo", "alexandria", "giza", "port+said", "suez", "luxor", "el+mahalla", "asyut", "al+mansurah", "tanda"},
	"ethiopia":           PresetLocations{"ethiopia", "addis+ababa", "gondar", "adama", "hawassa", "bahir+dar"},
	"vietnam":            PresetLocations{"vietnam", "viet+nam", "ho+chi+minh", "hanoi", "ha+noi", "hai+phong", "da+nang", "can+tho", "bien+hoa", "nha+trang", "vinh"},
	"iran":               PresetLocations{"iran", "tehran", "mashhad", "isfahan", "esfahan", "karaj", "shiraz", "tabriz", "qom", "ahvaz", "ahwaz", "kermanshah", "urmia", "rasht", "kerman"},
	"congo":              PresetLocations{"congo", "drc", "kinshasa", "lubumbashi", "bukavu", "kananga", "goma"},
	"turkey":             PresetLocations{"turkey", "turkiye", "istanbul", "ankara", "izmir", "bursa", "adana", "gaziantep", "konya", "antalya", "kayseri", "mersin", "eskisehir", "samsun", "denizli", "malatya"},
	"thailand":           PresetLocations{"thailand", "bangkok", "nonthaburi", "nakhon", "phuket", "pattaya", "chiang+mai"},
	"south africa":       PresetLocations{"south+africa", "south+africa", "johannesburg", "cape+town", "rsa", "durban", "port+elizabeth", "pretoria", "nelspruit"},
	"myanmar":            PresetLocations{"myanmar", "burma", "yangon", "rangoon", "mandalay", "nay+pyi+taw", "taunggyi", "bago", "mawlamyine"},
	"tanzania":           PresetLocations{"tanzania", "dar+es+salaam", "mwanza", "arusha", "dodoma", "mbeya", "morogoro", "tanga", "kilimanjaro"},
	"south korea":        PresetLocations{"south+korea", "ROK", "korea", "seoul", "busan", "incheon", "daegu", "daejeon", "gwangju", "대한민국", "서울", "서울시"},
	"colombia":           PresetLocations{"colombia", "bogota", "medellin", "cali", "barranquilla", "cartagena", "cucuta", "bucaramanga", "ibague", "soledad", "pereira", "santa+marta"},
	"kenya":              PresetLocations{"kenya", "nairobi", "mombasa", "kisumu", "nakuru", "eldoret", "kisii", "nyeri"},
	"argentina":          PresetLocations{"argentina", "buenos+aires", "cordoba", "rosario", "mendoza", "la+plata", "tucuman", "mar+del+plata", "salta", "resistencia"},
	"algeria":            PresetLocations{"algeria", "algiers", "oran", "constantine", "annaba", "blida", "batna", "djelfa", "setif", "sidi+bel+abbes", "biskra", "tiaret", "relizane", "mostaganem", "tlemcen", "chlef", "jijel"},
	"sudan":              PresetLocations{"sudan", "khartoum", "omdurman"},
	"poland":             PresetLocations{"poland", "polska", "warsaw", "krakow", "lodz", "wroclaw", "poznan", "gdansk", "szczecin", "bydgoszcz", "lublin", "katowice", "bialystok"},
	"canada":             PresetLocations{"canada", "ottawa", "edmonton", "winnipeg", "vancouver", "toronto", "quebec", "montreal", "mississauga", "calgary"},
	"australia":          PresetLocations{"australia", "sydney", "melbourne", "brisbane", "perth", "adelaide", "canberra", "hobart"},
	"new zealand":        PresetLocations{"new+zealand", "auckland", "wellington", "christchurch", "hamilton", "tauranga", "napier-hastings", "dunedin", "palmerston+north", "nelson", "rotorua", "whangarei", "new+plymouth", "invercargill", "whanganui", "gisborne"},
	"belgium":            PresetLocations{"belgium", "antwerp", "ghent", "charleroi", "liege", "brussels", "belgique"},
	"greece":             PresetLocations{"greece", "Ελλάδα", "athens", "thessaloniki", "patras", "heraklion", "larissa", "volos", "rhodes", "ioannina", "chania", "crete"},
	"peru":               PresetLocations{"peru", "lima", "cusco", "cuzco", "ica", "arequipa", "trujillo", "chiclayo", "huancayo", "piura", "chimbote", "iquitos", "juliaca", "cajamarca"},
	"hungary":            PresetLocations{"hungary", "magyarország", "budapest", "szeged", "miskolc"},
	"albania":            PresetLocations{"albania", "tirana", "durres", "vlore", "elbasan", "shkoder"},
	"uganda":             PresetLocations{"uganda", "kampala", "mbarara", "mukono", "jinja", "arua", "gulu", "masaka"},
	"zambia":             PresetLocations{"zambia", "lusaka", "kitwe", "ndola"},
	"sri lanka":          PresetLocations{"sri+lanka", "balangoda", "ratnapura", "colombo", "moratuwa", "negombo", "galle", "jaffna"},
	"singapore":          PresetLocations{"singapore"},
	"latvia":             PresetLocations{"latvia", "latvija", "riga", "rīga", "kuldiga", "kuldīga", "ventspils", "liepaja", "liepāja", "daugavpils", "jelgava", "jurmala", "jūrmala"},
	"romania":            PresetLocations{"romania", "bucharest", "cluj", "iasi", "timisoara", "craiova", "brasov", "sibiu", "constanta", "oradea", "galati", "ploesti", "pitesti", "arad", "bacau"},
	"belarus":            PresetLocations{"belarus", "minsk", "brest,belarus", "grodno", "gomel", "vitebsk", "mogilev", "slutsk", "borisov", "pinsk", "baranovichi", "bobruisk", "soligorsk"},
	"malta":              PresetLocations{"malta", "birgu", "bormla", "mdina", "qormi", "senglea", "siġġiewi", "valletta", "zabbar", "zebbuġ", "zejtun"},
	"rwanda":             PresetLocations{"rwanda", "kigali", "butare", "muhanga", "ruhengeri", "gisenyi"},
	"saudi arabia":       PresetLocations{"Saudi", "KSA", "Riyadh", "Mecca"},
	"morocco":            PresetLocations{"morocco", "casablanca", "fez", "tangier", "marrakesh", "salé", "meknes", "rabat", "oujda", "kenitra", "agadir", "tetouan", "temara", "safi", "mohammedia", "khouribga", "el+jadida"},
	"uzbekistan":         PresetLocations{"uzbekistan", "tashkent", "namangan", "samarkand", "andijan", "nukus", "bukhara", "qarshi", "fergana"},
	"malaysia":           PresetLocations{"malaysia", "kuala+lumpur", "kajang", "klang", "subang", "penang", "ipoh", "selangor", "melaka", "johor", "sabah"},
	"afghanistan":        PresetLocations{"afghanistan", "kabul", "kandahar", "herat", "mazar-e-sharif", "jalalabad", "ghazni", "nangarhar"},
	"venezuela":          PresetLocations{"venezuela", "caracas", "maracaibo", "barquisimeto", "guayana", "maturín", "zulia", "bolivar"},
	"ghana":              PresetLocations{"ghana", "accra", "kumasi", "sekondi", "ashaiman", "sunyani", "tamale", "tema"},
	"angola":             PresetLocations{"angola", "luanda", "huambo", "lobito", "benguela"},
	"nepal":              PresetLocations{"nepal", "kathmandu", "pokhara", "lalitpur", "bharatpur", "birgunj", "biratnagar", "janakpur", "ghorahi"},
	"yemen":              PresetLocations{"yemen", "sana'a", "taiz", "aden", "mukalla", "ibb"},
	"mozambique":         PresetLocations{"mozambique", "maputo", "matola", "nampula", "beira", "sofala", "chimoio", "tete", "quelimane"},
	"ivory coast":        PresetLocations{"ivory", "abidjan", "bouaké", "daloa", "yamoussoukro"},
	"cameroon":           PresetLocations{"cameroon", "Douala", "Yaoundé", "Bafoussam", "Bamenda", "Garoua", "Maroua", "Ngaoundéré", "Kumba", "Nkongsamba", "Buea"},
	"taiwan":             PresetLocations{"taiwan", "Taichung", "Kaohsiung", "Taipei", "Taoyuan", "Tainan", "Hsinchu", "Keelung", "Chiayi", "Changhua"},
	"niger":              PresetLocations{"niger", "Niamey", "Maradi", "Zinder", "Tahoua", "Agadez", "Arlit", "Birni-N'Konni", "Dosso", "Gaya", "Tessaoua"},
	"burkina faso":       PresetLocations{"burkina+faso", "Ouagadougou", "Bobo-Dioulasso", "Koudougou", "Banfora", "Ouahigouya", "Pouytenga", "Kaya", "Tenkodogo", "Fada+N'gourma", "Houndé"},
	"mali":               PresetLocations{"mali", "bamako", "sikasso", "kalabancoro", "koutiala", "ségou", "kayes", "kati", "mopti", "niono"},
	"malawi":             PresetLocations{"malawi", "Lilongwe", "Blantyre", "Mzuzu", "Zomba", "Karonga", "Kasungu", "Mangochi", "Salima", "Liwonde", "Balaka"},
	"chile":              PresetLocations{"chile", "Santiago", "Valparaíso", "Concepción", "La+Serena", "Antofagasta", "Temuco", "Rancagua", "Talca", "Arica", "Chillán"},
	"kazakhstan":         PresetLocations{"kazakhstan", "Almaty", "Shymkent", "Karagandy", "Taraz", "Nur-Sultan", "Pavlodar", "Oskemen", "Semey"},
	"guatemala":          PresetLocations{"Guatemala", "mixco", "villa+nueva", "petapa", "Quetzaltenango"},
	"ecuador":            PresetLocations{"ecuador", "Guayaquil", "Quito", "Cuenca", "Machala"},
	"syria":              PresetLocations{"syria", "aleppo", "homs", "latakia", "hama", "raqqa"},
	"cambodia":           PresetLocations{"cambodia", "phnom", "battambang", "siem+reap", "kampong"},
	"senegal":            PresetLocations{"senegal", "dakar", "touba", "thies", "rufisque"},
	"chad":               PresetLocations{"chad", "tchad", "n'djamena", "moundou"},
	"somalia":            PresetLocations{"somalia", "mogadishu", "hargeisa", "bosaso", "borama"},
	"zimbabwe":           PresetLocations{"zimbabwe", "harare", "bulawayo", "mutare", "gweru", "kwekwe"},
	"guinea":             PresetLocations{"conakry"},
	"benin":              PresetLocations{"benin", "cotonou", "porto-novo", "abomey"},
	"haiti":              PresetLocations{"haiti", "port-au-prince", "cap-haitien", "carrefour", "delmas", "petion-ville"},
	"cuba":               PresetLocations{"cuba", "havana", "santiago+de+cuba", "camaguey", "holguin", "guantanamo", "bayamo"},
	"bolivia":            PresetLocations{"bolivia", "santa+cruz+de+la+sierra", "el+alto", "la+paz", "cochabamba", "oruro", "sucre"},
	"tunisia":            PresetLocations{"tunisia", "tunis", "sfax", "sousse", "kairouan", "ariana", "gabes", "bizerte"},
	"south sudan":        PresetLocations{"south sudan", "juba"},
	"burundi":            PresetLocations{"burundi", "bujumbura", "gitega"},
	"dominican republic": PresetLocations{"dominican+republic", "republica+dominicana", "santo+domingo", "la+vega", "macoris"},
	"czech republic":     PresetLocations{"czech", "ceska", "prague", "budejovice", "plzen", "karlovy", "ostrava", "brno"},
	"jordan":             PresetLocations{"jordan", "amman", "zarqa", "irbid"},
	"azerbaijan":         PresetLocations{"azerbaijan", "baku", "sumqayit", "ganja", "lankaran"},
	"uae":                PresetLocations{"uae", "emirates", "dubai", "abu+dhabi", "sharjah", "al+ain", "ajman"},
	"honduras":           PresetLocations{"honduras", "tegucigalpa", "san+pedro+sula", "choloma", "la+ceiba", "el+progreso", "choluteca", "comayagua"},
	"tajikistan":         PresetLocations{"tajikistan", "dushanbe", "khujand"},
	"papua new guinea":   PresetLocations{"papua+new+guinea", "port+moresby", "lae"},
	"serbia":             PresetLocations{"serbia", "belgrade", "novi+sad", "nis", "kragujevac", "subotica", "zrenjanin", "pancevo", "cacak", "novi+pazar", "kraljevo", "smederevo"},
	"switzerland":        PresetLocations{"switzerland", "zurich", "zürich", "geneva", "basel", "lausanne", "bern", "winterthur", "lucerne", "gallen", "lugano", "biel", "thun"},
	"togo":               PresetLocations{"togo", "lome"},
	"sierra leone":       PresetLocations{"sierra+leone", "freetown", "makeni", "koidu"},
	"ireland":            PresetLocations{"ireland", "dublin", "cork", "limerick", "galway", "waterford+ireland", "drogheda", "dundalk"},
	"hong kong":          PresetLocations{"hong+kong", "kowloon"}}

func Preset(name string) []string {
	return PRESETS[name]
}
