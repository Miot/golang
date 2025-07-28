//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import "fmt"
import "time"

type Title string
type Name string
type LendAudit struct {
	checkOut time.Time
	checkIn time.Time
}
type Member struct {
	name Name
	books map[Title]LendAudit
}
type BookEntry struct {
	total int
	lended int
}
type Library struct {
	members map[Name]Member
	books map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "->", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLaibarayBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book", title, "not found")
		return false
	}
	if book.lended == book.total {
		fmt.Println("Book", title, "is not available")
		return false
	}
	book.lended += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book", title, "not found")
		return false
	}
	if book.lended == 0 {
		fmt.Println("Book", title, "is not lended")
		return false
	}
	book.lended -= 1
	library.books[title] = book
	member.books[title] = LendAudit{checkIn: time.Now()}
	return true
}

func main() {
	library := Library{
		books: make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books[Title("The Great Gatsby")] = BookEntry{total: 4, lended: 0}
	library.books[Title("The Catcher in the Rye")] = BookEntry{total: 3, lended: 0}
	library.books[Title("The Hobbit")] = BookEntry{total: 2, lended: 0}
	library.books[Title("The Lord of the Rings")] = BookEntry{total: 1, lended: 0}

	library.members[Name("John")] = Member{"John", make(map[Title]LendAudit)}
	library.members[Name("Billy")] = Member{"Billy", make(map[Title]LendAudit)}
	library.members[Name("Susan")] = Member{"Susan", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLaibarayBooks(&library)
	printMemberAudits(&library)

	member := library.members[Name("John")]
	checkedOut := checkoutBook(&library, Title("The Great Gatsby"), &member)
	fmt.Println("\nAfter checkout:")
	if checkedOut {
		printLaibarayBooks(&library)
		printMemberAudits(&library)
	}
	returned := returnBook(&library, Title("The Great Gatsby"), &member)
	fmt.Println("\nAfter return:")
	if returned {
		printLaibarayBooks(&library)
		printMemberAudits(&library)
	}
	
}

