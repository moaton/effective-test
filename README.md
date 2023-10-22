# Effective test

### Задача

> Реализовать сервис, который будет получать по апи ФИО, из открытых апи обогащать ответ наиболее вероятными возрастом, полом и национальностью и сохранять данные в БД. По запросу выдавать инфу о найденных людях. Необходимо реализовать следующее

<ol>
  <li>Выставить rest методы</li>
  <ol>
    <li>Для получения данных с различными фильтрами и пагинацией</li>
    <li>Для удаления по идентификатору</li>
    <li>Для изменения сущности</li>
    <li>Для добавления новых людей в формате</li>
  </ol>
  <code>
    ```
    {
      "name": "Dmitriy",
      "surname": "Ushakov",
      "patronymic": "Vasilevich" // необязательно
    }
    ```
  </code>
  <li>Корректное сообщение обогатить</li>
  <ol>
    <li>Возрастом - https://api.agify.io/?name=Dmitriy</li>
    <li>Полом - https://api.genderize.io/?name=Dmitriy</li>
    <li>Национальностью - https://api.nationalize.io/?name=Dmitriy</li>
  </ol>
  <li>Обогащенное сообщение положить в БД postgres (структура БД должна быть создана путем миграций)</li>
  <li>Покрыть код debug- и info-логами</li>
  <li>Вынести конфигурационные данные в .env</li>
</ol>
 

##  Запуск
```
make run
```