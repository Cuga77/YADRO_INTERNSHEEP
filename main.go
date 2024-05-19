package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

type Table struct {
    IsBusy bool
    Client string
    Income int
    Time   time.Duration
}

type Client struct {
    Name  string
    Table int
}

type Event struct {
    Time     time.Time
    ID       int
    Client   string
    Table    int
    IsIncome bool
}

func main() {
    // Загрузка данных из файла
    fileName := os.Args[1] 
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    tableCount, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println(err)
        return
    }
    tables := make([]Table, tableCount)

    scanner.Scan()
    workTime := strings.Split(scanner.Text(), " ")
    if len(workTime) < 2 {
        fmt.Println("Invalid input file")
        return
    }
    startTime, err := time.Parse("15:04", workTime[0])
    if err != nil {
        fmt.Println(err)
        return
    }
    endTime, err := time.Parse("15:04", workTime[1]) 
    if err != nil {
        fmt.Println(err)
        return
    }

    scanner.Scan()
    price, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println(err)
        return
    }

    var events []Event
    for scanner.Scan() {
        eventData := strings.Split(scanner.Text(), " ")
        if len(eventData) < 2 {
            fmt.Println("Invalid input file")
            return
        }
        eventTime, err := time.Parse("15:04", eventData[0])
        if err != nil {
            fmt.Println(err)
            return
        }
        eventID, err := strconv.Atoi(eventData[1])
        if err != nil {
            fmt.Println(err)
            return
        }
        eventClient := eventData[2]
        eventTable := 0
        if len(eventData) > 3 {
            eventTable, err = strconv.Atoi(eventData[3]) 
            if err != nil {
                fmt.Println(err)
                return
            }
        }
        events = append(events, Event{eventTime, eventID, eventClient, eventTable, false})
    }

    // Обработка событий
    var clients []Client
    for _, event := range events {
        switch event.ID {
        case 1:
            // Клиент пришел
            clients = append(clients, Client{event.Client, 0})
        case 2:
            // Клиент сел за стол
            for i, client := range clients {
                if client.Name == event.Client {
                    clients[i].Table = event.Table
                    tables[event.Table-1].IsBusy = true
                    tables[event.Table-1].Client = event.Client
                    break
                }
            }
        case 3:
            // Клиент ожидает
        case 4:
            // Клиент ушел
            for i, client := range clients {
                if client.Name == event.Client {
                    clients = append(clients[:i], clients[i+1:]...)
                    break
                }
            }
            if tables[event.Table-1].Client == event.Client {
                tables[event.Table-1].IsBusy = false
                tables[event.Table-1].Client = ""
            }
        }
    }

    // Подсчет выручки и времени занятости каждого стола
    for i := range tables {
        if tables[i].IsBusy {
            tables[i].Income += price
            tables[i].Time += endTime.Sub(startTime)
        }
    }

    // Вывод результатов
    fmt.Println(startTime.Format("15:04"))
    for _, event := range events {
        fmt.Printf("%s %d %s", event.Time.Format("15:04"), event.ID, event.Client)
        if event.Table != 0 {
            fmt.Printf(" %d", event.Table)
        }
        fmt.Println()
    }
    fmt.Println(endTime.Format("15:04"))
    for i, table := range tables {
        fmt.Printf("%d %d %s\n", i+1, table.Income, table.Time.String())
    }
}