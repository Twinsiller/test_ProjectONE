# test_ProjectONE

--Удаление данных из таблицы и обнуление значение serial:
    TRUNCATE TABLE db_name RESTART IDENTITY;

--Библиотеки:
    _ "github.com/lib/pq"
    password "github.com/vzglad-smerti/password_hash"

firstAPP:
    для запуска:
        1. go mod init gin-notes-api **
        2. go get -u github.com/gin-gonic/gin
        3. go get -u github.com/lib/pq
        4. go get -u github.com/vzglad-smerti/password_hash