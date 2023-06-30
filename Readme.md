Пример идеально выполненого тестового. Переписанный мною с ютуба.</br>
К исходному красивому тестовому я добавил докер контейнер..</br>
Буду использовать как шаблон для дальнейшего построения веб приложений.
```bash
 git clone https://github.com/Sib-Coder/TestForJunior
 docker build -t umbrela_app .
 docker run -p 8090:8090 --name umbrela_app umbrela_app
```

```
Server Runnig

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8090
2023/06/30 13:30:15 red button user detected
```