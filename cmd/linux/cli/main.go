package main

import (
    "fmt"
    "os"
    "github.com/kornev-aa/lab5/internal/pkg/app/cli"
    "github.com/kornev-aa/lab5/pkg/config"
    "github.com/kornev-aa/lab5/pkg/logger"
    "github.com/kornev-aa/lab5/pkg/storage"
)

func main() {
    cfg, err := config.Load("./config.json")
    if err != nil {
        fmt.Printf("Ошибка загрузки конфига: %s\n", err.Error())
        os.Exit(1)
    }

    log := logger.New()

    var store storage.LocationStorage
    switch cfg.StorageType {
    case "file":
        store = storage.NewFileStorage(cfg.FilePath)
        log.Info("Используется файловое хранилище")
    default:
        log.Error("Неизвестный тип хранилища", nil)
        os.Exit(1)
    }

    app := cli.New(log, store)

    if len(os.Args) > 2 && os.Args[1] == "save" {
        var lat, lon float64
        fmt.Sscanf(os.Args[2], "%f", &lat)
        fmt.Sscanf(os.Args[3], "%f", &lon)
        if err := app.SaveLocation(lat, lon); err != nil {
            log.Error("Ошибка сохранения", err)
            os.Exit(1)
        }
        log.Info("Координаты сохранены")
        return
    }

    if err := app.Run(); err != nil {
        log.Error("Приложение завершилось с ошибкой", err)
        os.Exit(1)
    }

    log.Info("Приложение завершилось успешно")
}
