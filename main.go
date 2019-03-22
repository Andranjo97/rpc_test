package main

import "fmt"

type Item struct {
    title string
    body string
}

var database []Item

func getByName(title string) Item {
    var getItem Item

    for _, val := range database {
        if val.title == title {
            getItem = val
        }
    }

    return getItem
}

func addItem(item Item) Item {
    database = append(database, item)
    return item
}

func editItem(title string, edit Item) Item {
    var changed Item

    for idx, val := range database {
        if val.title == title {
            database[idx] = edit
            changed = edit
        }
    }

    return changed
}

func deleteItem(item Item) Item {
    var del Item

    for idx, val := range database {
        if val.title == item.title && val.body == item.body {
            database = append(database[:idx], database[idx+1:]...)
            del = item
            break
        }
    }

    return del
}

func main() {
    fmt.Println("initial database", database)
    a := Item{"first", "a test item"}
    b := Item{"second", "a second item"}
    c := Item{"third", "a third item"}

    addItem(a)
    addItem(b)
    addItem(c)

    fmt.Println("second database", database)

    deleteItem(b)
    fmt.Println("third database", database)

    editItem("third", Item{"fourth", "a new Item"})
    fmt.Println("fourth database", database)

    x := getByName("fourth")
    y := getByName("first")

    fmt.Println(x, y)
}