package main

import (
	"fmt"
	"sort"
	"strings"
)

const NMAX = 100

var capital, remainingCapital, debt, profit int

type tabItem struct {
	idItem    int
	sellPrice int
	buyPrice  int
	name      string
	category  string
	stock     int
}

type tabTransaction struct {
	idTransaction int
	idItemSold    int
	sellPrice     int
	quantitySold  int
	nameSold      string
	categorySold  string
}

type items [NMAX]tabItem
type transactions [NMAX]tabTransaction

func main() {
	var I items
	var T transactions
	var m, n int
	var mainMenu int
	var running bool = true

	for capital <= 0 {
		fmt.Println("Enter capital:")
		fmt.Scan(&capital)
		if capital <= 0 {
			fmt.Println("Capital should be more than 0")
		}
	}

	clearScreen()
	remainingCapital = capital
	fmt.Print("Welcome to Toko Sembako\n")
	for running {
		fmt.Println("Main Menu")
		fmt.Println("1. Item Data")
		fmt.Println("2. Transaction Data")
		fmt.Println("3. Sort Data")
		fmt.Println("4. Search Data")
		fmt.Println("5. View Data")
		fmt.Println("6. Exit")
		fmt.Print("Please choose (1/2/3/4/5/6): ")
		fmt.Scan(&mainMenu)

		switch mainMenu {
		case 1:
			itemData(&I, &n)
		case 2:
			transactionData(&I, &T, &n, &m)
		case 3:
			sortData(I, T, n, m)
		case 4:
			searchData(I, T, n, m)
		case 5:
			viewData(I, T, n, m)
		case 6:
			clearScreen()
			fmt.Println("Thanks!")
			running = false
		default:
			clearScreen()
			fmt.Println("Invalid input")
		}
	}
}

func itemData(I *items, n *int) {
	var choice int
	fmt.Println("Item Data")
	fmt.Println("1. View Item Data")
	fmt.Println("2. Add Item Data")
	fmt.Println("3. Edit Item Data")
	fmt.Println("4. Delete Item Data")
	fmt.Println("5. Back")
	fmt.Print("Please choose (1/2/3/4/5): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		clearScreen()
		viewItemData(*I, *n)
		itemData(I, n)
	case 2:
		clearScreen()
		addItemData(I, n)
		itemData(I, n)
	case 3:
		clearScreen()
		editItemData(I, *n)
		itemData(I, n)
	case 4:
		clearScreen()
		deleteItemData(I, n)
		itemData(I, n)
	case 5:
		clearScreen()
		return
	default:
		clearScreen()
		fmt.Println("Invalid input")
		itemData(I, n)
	}
}

func addItemData(I *items, n *int) {
	if *n >= NMAX {
		fmt.Println("Item data is full")
		return
	}
	fmt.Println("Add Item Data")
	fmt.Println("Remaining capital:", remainingCapital)
	fmt.Println("Add new item data")
	I[*n].idItem = *n + 1
	fmt.Print("Item Name: ")
	fmt.Scan(&I[*n].name)
	fmt.Print("Item Category: ")
	fmt.Scan(&I[*n].category)
	for I[*n].name == I[*n].category {
		fmt.Println("Name and category cannot be the same")
		fmt.Println("Please re-enter the data")
		fmt.Print("Item Name: ")
		fmt.Scan(&I[*n].name)
		fmt.Print("Item Category: ")
		fmt.Scan(&I[*n].category)
	}
	fmt.Print("Purchase Price: ")
	fmt.Scan(&I[*n].buyPrice)
	fmt.Print("Selling Price: ")
	fmt.Scan(&I[*n].sellPrice)
	for I[*n].buyPrice > I[*n].sellPrice {
		fmt.Println("Purchase price must be less than the selling price")
		fmt.Println("Please re-enter the price")
		fmt.Print("Purchase Price: ")
		fmt.Scan(&I[*n].buyPrice)
		fmt.Print("Selling Price: ")
		fmt.Scan(&I[*n].sellPrice)
	}
	fmt.Print("Stock: ")
	fmt.Scan(&I[*n].stock)
	for I[*n].stock <= 0 {
		fmt.Println("Stock must be more than 0")
		fmt.Println("Please re-enter the stock")
		fmt.Print("Stock: ")
		fmt.Scan(&I[*n].stock)
	}
	clearScreen()
	remainingCapital -= I[*n].buyPrice * I[*n].stock
	if remainingCapital < 0 {
		fmt.Println("Not enough capital")
		remainingCapital += I[*n].buyPrice * I[*n].stock
		return
	}
	*n++
	fmt.Println("Item data successfully added")
	fmt.Println("Remaining Capital:", remainingCapital)
	itemData(I, n)
}

func editItemData(I *items, n int) {
	var editID, menuEdit int
	var editName string
	var check, newBuyPrice, newStock int
	fmt.Println("Edit Item Data")
	if n == 0 {
		fmt.Println("Item data is empty")
		return
	}
	viewItemData(*I, n)
	fmt.Println("Edit Item Menu")
	fmt.Println("1. Edit item by ID")
	fmt.Println("2. Edit item by Item Name")
	fmt.Println("3. Back")
	fmt.Print("Please choose (1/2/3): ")
	fmt.Scan(&menuEdit)
	switch menuEdit {
	case 1:
		fmt.Print("Enter Item ID: ")
		fmt.Scan(&editID)
		check = findItemByID(*I, n, editID)
		if check != -1 {
			clearScreen()
			fmt.Println("Item found")
			fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Item ID", "Item Name", "Item Category", "Purchase Price", "Selling Price", "Stock")
			fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", I[check].idItem, I[check].name, I[check].category, I[check].buyPrice, I[check].sellPrice, I[check].stock)
			fmt.Println("Edit new item data")
			fmt.Print("Item Name: ")
			fmt.Scan(&I[check].name)
			fmt.Print("Item Category: ")
			fmt.Scan(&I[check].category)
			for I[check].name == I[check].category {
				fmt.Println("Name and category cannot be the same")
				fmt.Println("Please re-enter the data")
				fmt.Print("Item Name: ")
				fmt.Scan(&I[check].name)
				fmt.Print("Item Category: ")
				fmt.Scan(&I[check].category)
			}
			fmt.Print("Purchase Price: ")
			fmt.Scan(&I[check].buyPrice)
			fmt.Print("Selling Price: ")
			fmt.Scan(&I[check].sellPrice)
			for I[check].buyPrice > I[check].sellPrice {
				fmt.Println("Purchase price must be less than the selling price")
				fmt.Println("Please re-enter the price")
				fmt.Print("Purchase Price: ")
				fmt.Scan(&I[check].buyPrice)
				fmt.Print("Selling Price: ")
				fmt.Scan(&I[check].sellPrice)
			}
			fmt.Print("Stock: ")
			fmt.Scan(&newStock)
			for newStock <= 0 {
				fmt.Println("Stock must be more than 0")
				fmt.Println("Please re-enter the stock")
				fmt.Print("Stock: ")
				fmt.Scan(&newStock)
			}
			if newBuyPrice*newStock > I[check].buyPrice*I[check].stock {
				remainingCapital -= (newBuyPrice*newStock - I[check].buyPrice*I[check].stock)
			} else if newBuyPrice*newStock < I[check].buyPrice*I[check].stock {
				remainingCapital += (I[check].buyPrice*I[check].stock - newBuyPrice*newStock)
			}
			I[check].buyPrice = newBuyPrice
			I[check].stock = newStock
			if remainingCapital < 0 {
				debt = (remainingCapital * -1) + debt
				remainingCapital = 0
			}
			clearScreen()
			fmt.Println("Item data successfully edited")
			viewItemData(*I, n)
		} else {
			clearScreen()
			fmt.Println("Item not found")
			editItemData(I, n)
		}
	case 2:
		fmt.Print("Enter Item Name: ")
		fmt.Scan(&editName)
		check = findItemByName(*I, n, editName)
		if check != -1 {
			clearScreen()
			fmt.Println("Item found")
			fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Item ID", "Item Name", "Item Category", "Purchase Price", "Selling Price", "Stock")
			fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", I[check].idItem, I[check].name, I[check].category, I[check].buyPrice, I[check].sellPrice, I[check].stock)
			fmt.Println("Edit new item data")
			fmt.Print("Item Name: ")
			fmt.Scan(&I[check].name)
			fmt.Print("Item Category: ")
			fmt.Scan(&I[check].category)
			for I[check].name == I[check].category {
				fmt.Println("Name and category cannot be the same")
				fmt.Println("Please re-enter the data")
				fmt.Print("Item Name: ")
				fmt.Scan(&I[check].name)
				fmt.Print("Item Category: ")
				fmt.Scan(&I[check].category)
			}
			fmt.Print("Purchase Price: ")
			fmt.Scan(&I[check].buyPrice)
			fmt.Print("Selling Price: ")
			fmt.Scan(&I[check].sellPrice)
			for newBuyPrice > I[check].sellPrice {
				fmt.Println("Purchase price must be less than the selling price")
				fmt.Println("Please re-enter the price")
				fmt.Print("Purchase Price: ")
				fmt.Scan(&newBuyPrice)
				fmt.Print("Selling Price: ")
				fmt.Scan(&I[check].sellPrice)
			}
			fmt.Print("Stock: ")
			fmt.Scan(&newStock)
			for newStock <= 0 {
				fmt.Println("Stock must be more than 0")
				fmt.Println("Please re-enter the stock")
				fmt.Print("Stock: ")
				fmt.Scan(&newStock)
			}
			if newBuyPrice*newStock > I[check].buyPrice*I[check].stock {
				remainingCapital -= (newBuyPrice*newStock - I[check].buyPrice*I[check].stock)
			} else if newBuyPrice*newStock < I[check].buyPrice*I[check].stock {
				remainingCapital += (I[check].buyPrice*I[check].stock - newBuyPrice*newStock)
			}
			I[check].buyPrice = newBuyPrice
			I[check].stock = newStock
			if remainingCapital < 0 {
				debt = (remainingCapital * -1) + debt
				remainingCapital = 0
			}
			clearScreen()
			fmt.Println("Item data successfully edited")
			viewItemData(*I, n)
		} else {
			clearScreen()
			fmt.Println("Item not found")
			editItemData(I, n)
		}
	case 3:
		clearScreen()
		return
	default:
		clearScreen()
		fmt.Println("Invalid input")
		editItemData(I, n)
	}
}

func deleteItemData(I *items, n *int) {
	var deleteID, menuDelete int
	var deleteName, confirmDelete string
	var check int
	fmt.Println("Delete Item Data")
	if *n == 0 {
		fmt.Println("Item data is empty")
		return
	}
	viewItemData(*I, *n)
	fmt.Println("Delete Item Menu")
	fmt.Println("1. Delete item by ID")
	fmt.Println("2. Delete item by Item Name")
	fmt.Println("3. Back")
	fmt.Print("Please choose (1/2/3): ")
	fmt.Scan(&menuDelete)
	switch menuDelete {
	case 1:
		fmt.Print("Enter Item ID: ")
		fmt.Scan(&deleteID)
		check = findItemByID(*I, *n, deleteID)
		if check != -1 {
			fmt.Print("Are you sure (Y/N): ")
			fmt.Scan(&confirmDelete)
			if confirmDelete == "Y" || confirmDelete == "y" {
				remainingCapital += I[check].buyPrice * I[check].stock
				for i := check; i < *n-1; i++ {
					I[i] = I[i+1]
					I[i].idItem = i + 1
				}
				*n--
				clearScreen()
				fmt.Println("Item data successfully deleted!")
				viewItemData(*I, *n)
			} else {
				clearScreen()
				fmt.Println("Item data not deleted")
				viewItemData(*I, *n)
			}
		} else {
			clearScreen()
			fmt.Println("Item not found")
			deleteItemData(I, n)
		}
	case 2:
		fmt.Print("Enter Item Name: ")
		fmt.Scan(&deleteName)
		check = findItemByName(*I, *n, deleteName)
		if check != -1 {
			fmt.Print("Are you sure (Y/N): ")
			fmt.Scan(&confirmDelete)
			if confirmDelete == "Y" || confirmDelete == "y" {
				remainingCapital += I[check].buyPrice * I[check].stock
				for i := check; i < *n-1; i++ {
					I[i] = I[i+1]
					I[i].idItem = i + 1
				}
				*n--
				clearScreen()
				fmt.Println("Item data successfully deleted!")
				viewItemData(*I, *n)
			} else {
				clearScreen()
				fmt.Println("Item data not deleted")
				viewItemData(*I, *n)
			}
		} else {
			clearScreen()
			fmt.Println("Item not found")
			deleteItemData(I, n)
		}
	case 3:
		clearScreen()
		return
	default:
		clearScreen()
		fmt.Println("Invalid input")
		deleteItemData(I, n)
	}
}

func transactionData(I *items, T *transactions, n, m *int) {
	var transactionChoice int
	fmt.Println("Transaction Data")
	fmt.Println("1. View Transactions")
	fmt.Println("2. Add Transaction")
	fmt.Println("3. Back")
	fmt.Print("Please choose (1/2/3): ")
	fmt.Scan(&transactionChoice)

	switch transactionChoice {
	case 1:
		clearScreen()
		viewTransactionData(*T, *m)
	case 2:
		clearScreen()
		addTransactionData(I, T, n, m)
	case 3:
		clearScreen()
		return
	default:
		clearScreen()
		fmt.Println("Invalid input")
		transactionData(I, T, n, m)
	}
}

func addTransactionData(I *items, T *transactions, n, m *int) {
	if *n == 0 {
		fmt.Println("No items available for transaction")
		return
	}

	var itemID, quantitySold int
	var transactionConfirmed string
	viewItemData(*I, *n)
	fmt.Print("Enter Item ID to sell: ")
	fmt.Scan(&itemID)
	itemIndex := findItemByID(*I, *n, itemID)

	if itemIndex == -1 {
		fmt.Println("Item not found")
		return
	}

	fmt.Print("Enter quantity to sell: ")
	fmt.Scan(&quantitySold)

	if quantitySold <= 0 || quantitySold > I[itemIndex].stock {
		fmt.Println("Invalid quantity")
		return
	}

	fmt.Print("Confirm transaction (Y/N): ")
	fmt.Scan(&transactionConfirmed)

	if transactionConfirmed != "Y" && transactionConfirmed != "y" {
		fmt.Println("Transaction cancelled")
		return
	}

	T[*m].idTransaction = *m + 1
	T[*m].idItemSold = I[itemIndex].idItem
	T[*m].sellPrice = I[itemIndex].sellPrice
	T[*m].quantitySold = quantitySold
	T[*m].nameSold = I[itemIndex].name
	T[*m].categorySold = I[itemIndex].category

	I[itemIndex].stock -= quantitySold
	profit += (I[itemIndex].sellPrice - I[itemIndex].buyPrice) * quantitySold
	*m++

	clearScreen()
	fmt.Println("Transaction successfully added!")
	viewTransactionData(*T, *m)
}

func viewTransactionData(T transactions, m int) {
	if m == 0 {
		fmt.Println("No transactions available")
		return
	}

	fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Transaction ID", "Item ID", "Item Name", "Category", "Sell Price", "Quantity Sold")
	for i := 0; i < m; i++ {
		fmt.Printf("%-20d %-20d %-20s %-20s %-20d %-20d\n", T[i].idTransaction, T[i].idItemSold, T[i].nameSold, T[i].categorySold, T[i].sellPrice, T[i].quantitySold)
	}
}

func sortData(I items, T transactions, n, m int) {
	var sortChoice, sortOrder int
	fmt.Println("Sort Data")
	fmt.Println("1. Sort Items")
	fmt.Println("2. Sort Transactions")
	fmt.Println("3. Back")
	fmt.Print("Please choose (1/2/3): ")
	fmt.Scan(&sortChoice)

	switch sortChoice {
	case 1:
		fmt.Println("Sort Items by")
		fmt.Println("1. Name")
		fmt.Println("2. Category")
		fmt.Println("3. Buy Price")
		fmt.Println("4. Sell Price")
		fmt.Println("5. Stock")
		fmt.Print("Please choose (1/2/3/4/5): ")
		fmt.Scan(&sortOrder)
		sortItems(&I, n, sortOrder)
	case 2:
		fmt.Println("Sort Transactions by")
		fmt.Println("1. Item Name")
		fmt.Println("2. Category")
		fmt.Println("3. Sell Price")
		fmt.Println("4. Quantity Sold")
		fmt.Print("Please choose (1/2/3/4): ")
		fmt.Scan(&sortOrder)
		sortTransactions(&T, m, sortOrder)
	case 3:
		clearScreen()
		return
	default:
		clearScreen()
		fmt.Println("Invalid input")
		sortData(I, T, n, m)
	}
}

func sortItems(I *items, n int, order int) {
	switch order {
	case 1:
		sort.Slice(I[:n], func(i, j int) bool { return I[i].name < I[j].name })
	case 2:
		sort.Slice(I[:n], func(i, j int) bool { return I[i].category < I[j].category })
	case 3:
		sort.Slice(I[:n], func(i, j int) bool { return I[i].buyPrice < I[j].buyPrice })
	case 4:
		sort.Slice(I[:n], func(i, j int) bool { return I[i].sellPrice < I[j].sellPrice })
	case 5:
		sort.Slice(I[:n], func(i, j int) bool { return I[i].stock < I[j].stock })
	default:
		fmt.Println("Invalid sort order")
		return
	}
	clearScreen()
	fmt.Println("Items successfully sorted")
	viewItemData(*I, n)
}

func sortTransactions(T *transactions, m int, order int) {
	switch order {
	case 1:
		sort.Slice(T[:m], func(i, j int) bool { return T[i].nameSold < T[j].nameSold })
	case 2:
		sort.Slice(T[:m], func(i, j int) bool { return T[i].categorySold < T[j].categorySold })
	case 3:
		sort.Slice(T[:m], func(i, j int) bool { return T[i].sellPrice < T[j].sellPrice })
	case 4:
		sort.Slice(T[:m], func(i, j int) bool { return T[i].quantitySold < T[j].quantitySold })
	default:
		fmt.Println("Invalid sort order")
		return
	}
	clearScreen()
	fmt.Println("Transactions successfully sorted")
	viewTransactionData(*T, m)
}

func searchData(I items, T transactions, n, m int) {
	var searchChoice int
	fmt.Println("Search Data")
	fmt.Println("1. Search Items")
	fmt.Println("2. Search Transactions")
	fmt.Println("3. Back")
	fmt.Print("Please choose (1/2/3): ")
	fmt.Scan(&searchChoice)

	switch searchChoice {
	case 1:
		searchItems(I, n)
	case 2:
		searchTransactions(T, m)
	case 3:
		clearScreen()
		return
	default:
		clearScreen()
		fmt.Println("Invalid input")
		searchData(I, T, n, m)
	}
}

func searchItems(I items, n int) {
	var searchQuery string
	fmt.Print("Enter item name to search: ")
	fmt.Scan(&searchQuery)
	clearScreen()
	fmt.Println("Search Results")
	fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Item ID", "Item Name", "Item Category", "Purchase Price", "Selling Price", "Stock")
	for i := 0; i < n; i++ {
		if strings.Contains(strings.ToLower(I[i].name), strings.ToLower(searchQuery)) {
			fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", I[i].idItem, I[i].name, I[i].category, I[i].buyPrice, I[i].sellPrice, I[i].stock)
		}
	}
}

func searchTransactions(T transactions, m int) {
	var searchQuery string
	fmt.Print("Enter item name to search: ")
	fmt.Scan(&searchQuery)
	clearScreen()
	fmt.Println("Search Results")
	fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Transaction ID", "Item ID", "Item Name", "Category", "Sell Price", "Quantity Sold")
	for i := 0; i < m; i++ {
		if strings.Contains(strings.ToLower(T[i].nameSold), strings.ToLower(searchQuery)) {
			fmt.Printf("%-20d %-20d %-20s %-20s %-20d %-20d\n", T[i].idTransaction, T[i].idItemSold, T[i].nameSold, T[i].categorySold, T[i].sellPrice, T[i].quantitySold)
		}
	}
}

func viewData(I items, T transactions, n, m int) {
	var viewChoice int
	fmt.Println("View Data")
	fmt.Println("1. View Items")
	fmt.Println("2. View Transactions")
	fmt.Println("3. Back")
	fmt.Print("Please choose (1/2/3): ")
	fmt.Scan(&viewChoice)

	switch viewChoice {
	case 1:
		clearScreen()
		viewItemData(I, n)
	case 2:
		clearScreen()
		viewTransactionData(T, m)
	case 3:
		clearScreen()
		return
	default:
		clearScreen()
		fmt.Println("Invalid input")
		viewData(I, T, n, m)
	}
}

func viewItemData(I items, n int) {
	fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Item ID", "Item Name", "Item Category", "Purchase Price", "Selling Price", "Stock")
	for i := 0; i < n; i++ {
		fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", I[i].idItem, I[i].name, I[i].category, I[i].buyPrice, I[i].sellPrice, I[i].stock)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func findItemByID(I items, n, id int) int {
	for i := 0; i < n; i++ {
		if I[i].idItem == id {
			return i
		}
	}
	return -1
}

func findItemByName(I items, n int, name string) int {
	for i := 0; i < n; i++ {
		if strings.EqualFold(I[i].name, name) {
			return i
		}
	}
	return -1
}
