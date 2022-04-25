package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
	"github.com/xuri/excelize/v2"
)

func excel(date string, food string) {
	// PUT THE PATH TO YOUR .xlxs FILE HERE
	var path = ""

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// close spreadsheet
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.Rows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	//get the "current row"
	row := strconv.Itoa(rows.TotalRows() + 1)
	dateAxis := "A" + row
	foodAxis := "B" + row
	//write the entry into the file
	if err := f.SetCellValue("Sheet1", dateAxis, date); err != nil {
		fmt.Println(err)
	}
	if err := f.SetCellValue("Sheet1", foodAxis, food); err != nil {
		fmt.Println(err)
	}
	//save
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
}
func main() {
	var date string
	var food string

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add an entry to the log",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "date",
						Usage:       "date of the entry, in (m{m}/d{d}/yyyy)",
						Destination: &date,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "food",
						Usage:       "the item of food to be logged",
						Destination: &food,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					excel(date, food)
					return nil
				},
			},
		},
		Name:        "FoodLog",
		Description: "keep track of what you eat and when",
		Usage:       "use the 'add' subcommand to log your food entries",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
