package main

import "fmt"

func DumpCategoryData(c *Category) {
	fmt.Printf("\n\n\n")

	fmt.Println("--Category", "id", c.ID, "name", c.Name)
	for _, product := range c.Products {
		fmt.Println(" |")
		fmt.Println(" +--Product", "id", product.ID, "name", product.Name)
		for _, item := range product.Items {
			fmt.Println(" |  |")
			fmt.Println(" |  +--Item", "id", item.ID, "name", item.Name)
		}

		for i, factory := range product.Factories {
			if i == 0 {
				fmt.Println(" |  |")
			} else {
				fmt.Println(" |  |")
			}

			fmt.Println(" |  +--Factory", "id", factory.ID, "name", factory.Name)

			for _, workshop := range factory.Workshops {
				if i < (len(product.Factories) - 1) {
					fmt.Println(" |  |  |")
					fmt.Println(" |  |  +--Workshop", "id", workshop.ID, "name", workshop.Name)
				} else {
					fmt.Println(" |     |")
					fmt.Println(" |     +--Workshop", "id", workshop.ID, "name", workshop.Name)
				}

			}
		}
	}

	fmt.Printf("\n\n\n")
}
