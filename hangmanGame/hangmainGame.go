package hangmangame

import (
	"fmt"
	"strings"
)

func Hangman() {
	word := "elephant"

	//loookup for entries made by user
	entries := map[string]bool{}

	//list of "_" corrosponding to number of letters in word. [_ _ _ _]
	placeholders := []string{}

	//get length of the word
	//initialize the slize with each element as "_"
	for i := 0; i < len(word); i++ {
		placeholders = append(placeholders, "_")
	}
	fmt.Println("entries  are -->", entries)
	fmt.Println("placeholders  is -->", placeholders)
	chances := 8
	for {

		//If user gueses a wrong input then they loss a chance.
		userInput := strings.Join(placeholders, "")
		if chances == 0 && userInput != word {
			fmt.Println("Game over!! Please try again.")
			break
		}

		//evaluate a win
		if userInput == word {
			fmt.Print("You win")
			break
		}

	}

}
