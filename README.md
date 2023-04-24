# link-shortener
Тестовое задание для стажера-разработчика OZON. Сервис для создания сокращенных ссылок.

# Содержание
1. [Задача](#Задача)
2. [Решения](#Решения)
3. [Эндпоинты](#Эндпоинты)
4. [Сборка](#Сборка)
5. [Архитектура](#Архитектура)
6. [Тестирование](#Тестирование)

## Задача

Разработать сервис для создания сокращенных ссылок и предоставить взаимодействие через REST API и gRPC.
Реализовать in-memory и PostgreSQL хранилища данных. Покрыть unit-тестами. Реализовать развёртывание через Dockerfile.
Подробное условие задачи в [TASK.md](https://github.com/ilyadubrovsky/link-shortener/blob/main/TASK.md).

## Решения
+ **REST API** для работы сервиса по HTTP с использованием фреймворка [Gin](https://github.com/gin-gonic/gin);
+ Реализована работа сервиса через **gRPC**;
+ Используются принципы **Clean Architecture** в основе проектирования приложения;
+ СУБД **PostgreSQL** для хранения данных (на SQL запросах): [pgx](https://github.com/jackc/pgx);
+ **in-memory** решение для хранения данных реализовано с помощью **map** и **sync.RWMutex**;
+ **Docker** и надстройка **docker-compose** для развёртывания;
+ **Unit-тестирование** для уровня transport и service с помощью [gomock](https://github.com/golang/mock) и [testify](https://github.com/stretchr/testify);
+ Конфигурация приложения: [cleanenv](https://github.com/ilyakaznacheev/cleanenv);
+ Многоуровневое логирование: [logrus](https://github.com/sirupsen/logrus);

## Эндпоинты

- **GET:** /api/:token - получение оригинального URL
   - Параметры запроса:
      - token - уникальный токен, полученный при сокращении URL
   - Тело ответа:
      - raw_url - оригинальный URL
- **POST:** /api - сокращение URL
   - Тело запроса:
      - raw_url - оригинальный URL для сокращения
   - Тело ответа:
      - token - уникальный токен после сокращения

## Сборка

1. Создайте `.env` файл для всех сервисов и укажите переменные среды, перечисленные ниже;
2. Выполните сборку проекта: `make build`;
3. Выполните запуск проекта: `make run-inmemory` для in-memory хранилища данных и `make run-postgresql` для PostgreSQL;
4. При первом запуске с PostgreSQL примените миграции к базе данных: `make migrate-up`.

   | Переменная        | Описание                                             |
   |------------------------------------------------------|-----------------------------------------------------|
   | `SERVER_PORT`     | Порт, по умолчанию указывать `8080`                  |
   | `STORAGE_CONNECT` | Тип хранилища данных: `postgresql` или `inmemory`    |
   | `TRANSPORT_TYPE`  | Тип сервера: `grpc` или `http`                       |
   | `PG_USERNAME`     | Имя пользователя, по умолчанию указывать `postgres`  |
   | `PG_PASSWORD`     | Пароль, по умолчанию указывать `12345678`            |
   | `PG_HOST`         | Хост, по умолчанию указывать `db`                    |
   | `PG_PORT`         | Пароль, по умолчанию указывать `12345678`            |
   | `PG_DATABASE`     | База данных, по умолчанию указывать `link-shortener` |

## Архитектура

Проект делится на уровни: transport, service, storage. 
1. Transport - транспортный уровень, отвечает за доставку данных, взаимодействует со слоем service.
2. Service - уровень бизнес логики, отвечает за всю бизнес логику проекта, взаимодействует со слем storage.
3. Storage - уровень данных, отвечает за хранение данных.

## Тестирование

1. Выполните генерацию моков: `make generate`;
2. Запустите выполнение тестов: `make test`;