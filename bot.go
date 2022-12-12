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

func getRandomPokemon(response Response) string {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)
    pokemon := response.Pokemon[r.Intn(len(response.Pokemon))].Species.Name
    ret := "a"
    if firstChar := pokemon[0:1]; isVowel(firstChar) {
        ret += "n"
    }
    return "I'm " + ret + " *" + pokemon + "* today\\!"
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
    
        article := tgbotapi.NewInlineQueryResultArticleMarkdownV2("1", "What Pokemon are you today?", getRandomPokemon(pokecache))
        article.Description = "Choose your Pokemon!"
        article.ThumbURL = "https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/Pokebola-pokeball-png-0.png/800px-Pokebola-pokeball-png-0.png"
    
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
    response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
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
    telegramBot()
}