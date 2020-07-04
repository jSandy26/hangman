package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"os"
	"encoding/csv"
)

var hangmanPics = []string{`    
    +---+
        |
        |
        |
       ===`, `
    +---+
    O   |
        |
        |
       ===`, `
    +---+
    O   |
    |   |
        |
       ===`, `
    +---+
    O   |
   /|   |
        |
       ===`, `
    +---+
    O   |
   /|\  |
        |
       ===`, `
    +---+
    O   |
   /|\  |
   /    |
       ===`, `
    +---+
    O   |
   /|\  |
   / \  |
       ===`}

var words = "ant baboon badger bat bear beaver camel cat clam cobra cougar coyote crow deer dog donkey duck eagle ferret fox frog goat goose hawk lion lizard llama mole monkey moose mouse mule newt otter owl panda parrot pigeon python rabbit ram rat raven rhino salmon seal shark sheep skunk sloth snake spider stork swan tiger toad trout turkey turtle weasel whale wolf wombat zebra"


func getRandomWord(words []string) string{
	seed := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(seed)
	return words[r1.Intn(len(words))]
}

func gameDisplay(missedLetters, correctLetters, guessWord string) {
	fmt.Println(hangmanPics[len(missedLetters)])
	fmt.Print("Missed letters: ")
	for _, letter := range missedLetters {
		fmt.Print(string(letter), " ")
	}
	fmt.Print("\n")
	// blanks := []string(len(guessWord), '_')
	blanks := make([]string, len(guessWord), len(guessWord))
	for i := range guessWord {
		blanks[i] = "_"
	}

	for index, value := range guessWord {
		if strings.Contains(correctLetters, string(value)){
			blanks[index] = string(value)
		}
	}

	for _, letter := range blanks {
		fmt.Print(letter, " ")
	}
	fmt.Print("\n")
}

func getGuess(alreadyGuessed string) string {
	for {
		fmt.Println("Guess a letter.")
		var guess string
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)
		if len(guess) != 1 {
			fmt.Println("Please enter a single letter.")
		} else if strings.Contains(alreadyGuessed, guess) {
			fmt.Println("You have already guessed that letter. Choose again.")
		} else if guess < "a" && guess > "z" {
			fmt.Println("Please enter a LETTER.")
		} else {
			return guess
		}
		
	}
}

func playAgain() bool {
	fmt.Println("Do you want to play again? (yes or no)")
	var answer string
	fmt.Scanln(&answer)
	answer = strings.ToLower(answer)
	if answer == "yes" || answer == "y" {
		return true
	} else {
		return false
	}
}

func parseLines(lines [][]string) []string {
	ret := make([]string, len(lines))
	for i, line := range lines {
		ret[i] = line[0]
		ret[i] = strings.ToLower(ret[i])
	}
	return ret
}

func main() {
	fmt.Println("\nH A N G M A N")
	// words := strings.Split(words, " ")
	var missedLetters string
	var correctLetters string
	gameIsDone := false

	file, err := os.Open("words.csv")

	if err != nil {
		fmt.Println("Failed to open the CSV file: words.csv")
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the provided CSV file.")
		os.Exit(1)
	}
	words := parseLines(lines)
	guessWord := getRandomWord(words)

	for {
		gameDisplay(missedLetters, correctLetters, guessWord)

		guess := getGuess(missedLetters + correctLetters)

		if strings.Contains(guessWord, guess){
			correctLetters = correctLetters + guess

			foundAllLetters := true
			for _, letter := range guessWord {
				if !(strings.Contains(correctLetters, string(letter))) {
					foundAllLetters = false
					break
				}
			}
			if foundAllLetters {
				fmt.Println("You Won!")
				gameIsDone = true
			}
		} else {
			missedLetters = missedLetters + guess
			
			if len(missedLetters) == len(hangmanPics) - 1 {
				gameDisplay(missedLetters, correctLetters, guessWord)
				fmt.Println("You lost! The secret word is:", guessWord)
				gameIsDone = true
			}
		}
		
		if gameIsDone {
			if playAgain() {
				gameIsDone = false
				missedLetters = ""
				correctLetters = ""
				guessWord = getRandomWord(words)
			} else {
				break
			}
		}
	}

}