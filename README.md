# Практическая работа №5
## Бакланова ЕС
## ЭФМО-01-25

Структура проекта:

<img width="471" height="422" alt="image" src="https://github.com/user-attachments/assets/88ac8c1b-4d4c-427f-98c1-f292e6f74ab5" />

#### Суть

В данной практике использовалась технология ORM и GORM для работы с Postgre БД.  

ORM позволяет работать с <L через объекты программирования вместо прямых SQL-запросов. Это ускоряет разработку и делает код более чистым.

GORM позволил:

1. Автоматически создавать SQL таблицы из Go-структур

2. Упростить работу со связями между таблицами многие-к-одному и многие-ко-многим

3. Заменить написания шаблонного SQL-кода 

#### Запросы

1. GET/health

   <img width="974" height="446" alt="image" src="https://github.com/user-attachments/assets/a898b2be-df86-4d1a-bee4-9c5b65090a60" />

2. POST /users

   Создает пользователя (имя, почта)
   <img width="974" height="479" alt="image" src="https://github.com/user-attachments/assets/53efe8c5-ec0b-45e3-a2bc-e24e1ea80582" />

4. POST /notes

   Создает заметку, привязанную к пользователю по ID и имеющую несколько тэгов
   <img width="974" height="587" alt="image" src="https://github.com/user-attachments/assets/4ec1606a-3504-4936-bf82-bcb477f1b7d9" />
   <img width="908" height="1189" alt="image" src="https://github.com/user-attachments/assets/1cf8907b-4f5f-4689-8b45-7c14a6a92fdd" />

5. GET /notes/2

   Получает заметку по ID Ответ отрадает связь многие-ко-многим и многие-к-одному (tags, User)
   <img width="919" height="1191" alt="image" src="https://github.com/user-attachments/assets/8705bce6-497c-4b3b-9ac9-af074c1670a5" />

   

