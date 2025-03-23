## keenetic antifilter routing

Автоматическая генерация правил роутинга для различных онлайн сервисов.
Используется для роутеров **keenetic**. Поддерживаемы сервисы:

- [youtube.com](routes/youtube-ipv4.bat)

Правила обновляются автоматически раз в сутки (при наличии изменений).

## Как добавлять правила

Откройте раздел "Network rules" -> "Routing".
В разделе [Пользовательские статические маршруты](https://help.keenetic.com/hc/ru/articles/360000925780-Статические-маршруты) 
нажмите "Удалить все", затем загрузите нужный вам список из файла, выбрав необходимый сетевой интерфейс,
который будет использован для тунелирования целевого трафика. 

![static routes](docs/static-routes-index.png)
