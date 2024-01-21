### Описание
Сервис, который получает по API ФИО, из открытых API обогощает ответ 
наиболее вероятным возрастом, полом и национальностью. Сохраняет данные в БД.
## Используемые открытые API
* https://api.agify.io/ - API, из которого сервис получает предпологаемый возраст
* https://api.genderize.io/ - API, из которого сервис получает предпологаемый пол
* https://api.nationalize.io/ - API, из которого сервис получает предпологаемую национальность
## Используемые библиотеки
*  https://github.com/gofrs/uuid - используется для работы с uuid
*  https://github.com/gorilla/mux - используется для создания сервера
*  https://github.com/jackc/pgx/v5 - драйвер для работы с postgresql
*  https://github.com/joho/godotenv - используется для работы с файлом .env
* https://github.com/Masterminds/squirrel - билдер запросов к БД
* https://github.com/sirupsen/logrus - используется для логгирования
## Стек
* Goland
* Postgres
