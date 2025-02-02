# Разработка StaticSiteWithAuth

Мы с Адалиндой пишем текстовую игрушку и документируем всё это дело в Docusaurus. Штука отличная: на вход принимает Markdown, на выходе дает полноценный статический сайт. Его можно легко захостить, к примеру, на GitHub Pages и он станет виден... для всех. Вот как раз всем показывать документацию мы не хотим. Поэтому необходимо какое-то решение, которое позволит читать документацию только своим.

## Решения

Напишу то, что я примерно перебрал, но так ни разу не попробовал:

* GitHub Enterprise
* Скорее всего, это решается через nginx с какой-нибудь волшебной конфигурацией, но на познание этого монстра уходит куча времени и сил. Ну его нафиг
* Твое решение с рабочим названием StaticSiteWithAuth

## Как это будет выглядеть

* Разворачиваем StaticSiteWithAuth на хостинг либо с помощью инструкции, либо средствами docker'а. Предпочитаю всегда первое, потому что дешевле и понятнее 👻
* Задаем `config.yaml` (для примера взял YAML, но можно любой другой формат):

  ```yaml
  users:
    - adalapki:
        password: обычный-пароль
    - agentlapki:
        password: ну-очень-сложный-пароль
    - bogdan:
        password: стоп-ты-что-подглядываешь?!
  ```

* Закидываем документацию в папку `site` (структура файлов и папок расписана [ниже](#структура-файлов-и-папок))
* Закидываем страницу с авторизацией `index.html` в папку `auth` (API расписал [ниже](#примерное-api))
* Запускаем эту махину

В процессе работы сервера можно спокойно менять содержимое папки `site`, т.е. обновлять статику.

## Структура файлов и папок

```ts
root
├─auth
│ ├─public
│ │ ├─...
│ │ └─logo.png
│ ├─...
│ └─index.html
├─site
│ ├─assets
│ ├─blog
│ ├─docs
│ ├─...
│ ├─404.html
│ └─index.html
├─... // файлы и папки сервера
└─config.yml
```

## Примерное API

* `GET /auth` — (публичный API) отдает `auth/index.html`
* `GET /auth/static/*` — (публичный API) отдает файлы из `auth/public` (или лучше просто в `public`?)
* `POST /auth`

  Запрос от клиента должен идти с полем `Authorization: Basic $(echo '$username:$password' | base64)`.

  Сервер должен возвращать:

  * `200 OK`

    С полем `Set-Cookie: sessionId=?; HttpOnly` в качестве успеха.

    Какое значение должно иметь `sessionId` — фиг его знает. Можно пропустить пару (`$username`, `$password`) сквозь какую-нибудь функцию с заданным в конфиге `seed`'ом для защиты от чтения, а можно влепить голый пароль, если лень делать, лол
  * `403 Forbidden`

* `GET *` (приватный API)

  Все остальные роуты обращаются к папке `site`, где и будет расположен, к примеру, Docusaurus. В качестве примера можно взять скомпилированный Docusaurus в [ветке gh-pages funnysockorg.github.io](https://github.com/funnysockorg/funnysockorg.github.io/tree/gh-pages).

  При доступе к этому роуту должны проверяться куки с `sessionId`. Если они не совпадают, то пользователя перенаправляет на `/auth`.
