[Документация](https://help.keenetic.com/hc/en-us/articles/360000925780-Static-routing) 
keenetic говорит нам о том, что в одном `*.bat` файле допустимо использовать не более 1024
строк.

> Important! There is a limit to the number of lines in a bat file; they should not be more than 1024.

Файлы вида `all-ipv4-*.bat` содержать конкатенацию правил роутинга всех доступных сервисов,
разбитые по 1024 строки. 

Таким образом, вы можете вгрузить точечно нужные вам сервисы,
либо все правила сразу, загрузив каждый файл `all-ipv4-*.bat` по
отдельности.
