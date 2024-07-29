package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
		"Hi, %v! You new around here?",
		"Howdy, %v! It's a rootin' tootin' good time with you around",
		"Greetings, %v! I trust you're having a marvelous day.",
		"Ahoy there, %v! What brings you to our digital shores?",
		"Salutations, %v! You've arrived just in time for an exciting adventure.",
		"Hello and welcome, %v! I hope your journey here was seamless.",
		"Salaam, %v! It's a pleasure to make your acquaintance.",
		"Well met, %v! I'm delighted to see you.",
		"Hola, %v! Bienvenido a nuestro lugar.",
		"Bonjour, %v! Je suis ravi de vous voir ici.",
		"Guten Tag, %v! Es freut mich dich zu sehen.",
		"Ciao, %v! Benvenuto nel tuo nuovo viaggio.",
		"Ni Hao, %v! Xin wèn nǐ hǎo le.",
		"Kon'nichiwa, %v! Dozo yoroshiku onegaishimasu.",
		"Aloha, %v! Mahalo for joining us.",
		"Hey there, %v! Glad to see a friendly face.",
		"Welcome, %v! It's great to have you on board.",
		"Hello and welcome, %v! Your presence is appreciated.",
		"Good day to you, %v! A warm welcome to our servers.",
		"%v! It's good to see you again!",
	}

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}