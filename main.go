package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/zahraayafni/sample-discord-bot/discord"
)

func main() {
	// ===== INTEGER INPUT =====
	// num1 := 5
	// num2 := 3

	// res := calculator.BasicCalculator(calculator.SUM_OPERATOR, num1, num2)
	// fmt.Println(res)

	// ===== INPUT FROM STRING =====
	// input := "1 * 100"
	// inputs := strings.Split(input, " ")

	// var err error
	// op, num1, num2 := "", 0, 0
	// if len(inputs) > 2 {
	// 	num1, err = strconv.Atoi(inputs[0])
	// 	if err != nil {
	// 		return
	// 	}
	// 	op = inputs[1]

	// 	num2, err = strconv.Atoi(inputs[2])
	// 	if err != nil {
	// 		return
	// 	}
	// }
	// res := calculator.BasicCalculator(op, num1, num2)
	// fmt.Println(res)

	// INIT DISCORD CLIENT
	sess, err := discord.InitDiscord()
	if err != nil {
		log.Fatal("failed init discordgo", err)
	}

	// INIT DISCORD HANDLER
	discord.RunDiscordHandler(sess)

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()

	fmt.Println("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
