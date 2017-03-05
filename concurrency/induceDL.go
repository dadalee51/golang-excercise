package main

import (
	"fmt"
	"time"
)

// reader will read
type Book struct {
	Title  string
	Body   string
	Author string
}

type Library struct {
	Shelf   []Book
	Outlet  chan Book
	Returns chan Book
}

// the person interface can either borrow a book or return a book.
type Person interface {
	BorrowBook(outlet chan Book)
	ReturnBook(retlet chan Book)
}

type Librarian struct {
	ReturnBasket []Book
	Name         string
}

type BookWorm struct {
	Bag  []Book
	Name string
}

// librarian put all her books in return basket to library shelf
func (libn Librarian) ReturnBook(retlet chan Book) {
	for _, book := range libn.ReturnBasket {
		retlet <- book
	}
	fmt.Println("Librarian finished returning books to shelf.")
}

// bookworm borrows 3 books from library each time.
func (bkwm BookWorm) BorrowBook(outlet chan Book) {
	i := 0
	for book := range outlet {
		if i > 3 {
			fmt.Println("BookWorm finished borrowing.")
			return
		}
		bkwm.Bag = append(bkwm.Bag, book)
		i++
	}
}

// library books are hooked to returns and outlets.
func (library Library)BookAction(){
	for _, book:= range library.Shelf{
		select{
		case book := <-library.Returns :
			library.Shelf = append(library.Shelf, book)
		case library.Outlet<- book:
			//do nothing?
		}
	}
}

// we implement a small library system where users can borrow books and return them to the library.
// the main program load some books into the Library shelf, and create some Characters
func main() {
	lib := Library{
		Shelf : []Book{},
		Outlet : make(chan Book),
		Returns: make(chan Book),
	}
	for i:=0;i<100;i++{
		lib.Shelf = append(lib.Shelf, Book{"title", "body", "auth"})
	}
	
	bw := BookWorm{
		[]Book{},
		"Adam",
	}
	
	lb := Librarian{
		[]Book{},
		"Sherl",
	}
	go lib.BookAction()
	go bw.BorrowBook(lib.Outlet)
	go lb.ReturnBook(lib.Returns)
	
	time.Sleep(time.Second * 20)
}
