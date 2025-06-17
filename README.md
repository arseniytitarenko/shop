# shop

### Архитектура
Архитектура программы и запросы соответствует предложенной схеме, а именно: \
`order-service` и `postgres` для него \
`payment-service` и `postgres` для него \
`nginx` в качестве `API Gateway` \
`rabbitmq` в качестве брокера сообщений \
Сервисы соответствуют чистой архитектуре.

Применены паттерны Transactional Inbox и Outbox в обоих сервисах.

### Спецификация
Примеры всех запросов по сервисам есть в postman коллекции: \
https://www.postman.com/olympguide/workspace/files/collection/40644038-86bd0aba-0c55-4957-bdce-a083ef1cb1f2?action=share&creator=40644038

### Запуск
Достаточно склонировать репозиторий и запустить из корня командой: \
`docker-compose up --build`