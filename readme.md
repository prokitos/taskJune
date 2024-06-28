При решении задачи была использована СУБД PostgreSQL. 
Был использован фрейморк FIBER для запуска сервера.
Для более простой миграции и обращению к бд использовалась библиотека GORM.

Для примера, сервер запущен на порте :8888
Посмотреть глобальные переменные можно по пути internal/config/local.yaml

При запуске сервера автоматически создастся база данных vortex с нужными таблицами и связями.

Доступны 4 роута.
GET	http://localhost:8888/GetOrderBook	
GET	http://localhost:8888/GetOrderHistory
POST	http://localhost:8888/SaveOrderBook
POST	http://localhost:8888/SaveOrder

примеры использований роутов можно посмотреть в сваггере.
http://localhost:8888/swagger/index.html	

запуск тестов
go test -v -cover ./...