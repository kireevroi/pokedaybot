package main

import (
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "os"
	"github.com/joho/godotenv"
    "net/http"
	"log"
    "io/ioutil"
    "encoding/json"
    "math/rand"
    "time"
    "fmt"
)


type Response struct {
    Name    string    `json:"name"`
    Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
    EntryNo int            `json:"entry_number"`
    Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
    Name string `json:"name"`
}

func getRandomPokemon(response Response) Pokemon {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)
    pokemon := response.Pokemon[r.Intn(len(response.Pokemon))]
    return pokemon
}
func a_an(str string) string {
    ret := "a"
    if firstChar := str[0:1]; isVowel(firstChar) {
        ret += "n"
    }
    return ret
}
func pokeToString(pokemon Pokemon) string {
    return "I'm " + a_an(pokemon.Species.Name) + " *" + pokemon.Species.Name + "* today\\!"
}
func getPokemonUrl(pokemon Pokemon) string {
    num := pokemon.EntryNo
    s := fmt.Sprintf("%03d", num)
    ret := "https://kireevroman.com/pokemon/" + s + ".jpeg"
    return ret
}
func isVowel(char string) bool {
    switch char {
    case
        "a",
        "o",
        "u",
        "i",
        "e":
        return true
    }
    return false
}
func telegramBot() {
		err := godotenv.Load("token.env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
    if err != nil {
        panic(err)
    }
    pokecache:= fetchPokemon()
    //Устанавливаем время обновления
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    //Получаем обновления от бота
    updates := bot.GetUpdatesChan(u)
	if err == nil {
		log.Println("No error");
	}
    if err != nil {
        log.Fatal(err)
    }

    for update := range updates {
    
        pokemon := getRandomPokemon(pokecache)
        ThumbURL := getPokemonUrl(pokemon)
        article := tgbotapi.NewInlineQueryResultPhoto("1", getPokemonUrl(pokemon))
        article.Caption = pokeToString(pokemon)
        article.Description = "Choose your Pokemon!"
        article.ThumbURL = ThumbURL
        article.Width = 256
        article.Height = 256
        article.ParseMode="MarkdownV2"
        inlineConf := tgbotapi.InlineConfig{
            InlineQueryID: update.InlineQuery.ID,
            IsPersonal:    true,
            CacheTime:     1,
            Results:       []interface{}{article},
        }
        if _, err := bot.Request(inlineConf); err != nil {
            log.Println(err)
        }
    }
}

func fetchPokemon() Response {
    //var res string
    response, err := http.Get("http://pokeapi.co/api/v2/pokedex/national/")
    if err != nil {
        log.Println(err)
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Println(err)
    }
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
    return responseObject
}


func main() {
    //test()
    telegramBot()
}