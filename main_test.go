package main

import (
    "strconv"
    "strings"
    "testing"
    "time"
)

func TestParseTime(t *testing.T) {
    timeString := "10:00"
    parsedTime, err := time.Parse("15:04", timeString)
    if err != nil {
        t.Errorf("Не удалось разобрать время '%s': %v", timeString, err)
    }
    if parsedTime.Hour() != 10 || parsedTime.Minute() != 0 {
        t.Errorf("Время '%s' разобрано некорректно: %v", timeString, parsedTime)
    }
}

func TestEventParsing(t *testing.T) {
    eventData := strings.Split("10:00 1 Ivan 1", " ")
    eventTime, err := time.Parse("15:04", eventData[0])
    if err != nil {
        t.Errorf("Не удалось разобрать время события '%s': %v", eventData[0], err)
    }
    eventID, err := strconv.Atoi(eventData[1])
    if err != nil {
        t.Errorf("Не удалось разобрать ID события '%s': %v", eventData[1], err)
    }
    eventClient := eventData[2]
    eventTable, err := strconv.Atoi(eventData[3])
    if err != nil {
        t.Errorf("Не удалось разобрать номер стола события '%s': %v", eventData[3], err)
    }
    if eventTime.Hour() != 10 || eventTime.Minute() != 0 || eventID != 1 || eventClient != "Ivan" || eventTable != 1 {
        t.Errorf("Данные события разобраны некорректно: %v", eventData)
    }
}
