# poseidon
poseidon - opensource multi-system registry for docker images

# In development
Only v2 manifest schema(no manifest list, oci)

on development - comments and todo write into Russian lang


# Запуск
Для запуска пропишите в hosts связь домена(можете выбрать любой) с 
127.0.0.1:port (port - порт по умолчанию можно посмотреть в main или выставить в env)

Не запускайте в докере, чтобы избежать танцев с бубнами, тк в демоне докера хардкод по пулу ток 
в оригинальный registry localhost, если знаете как обойти ограничение - создайте issue на fix
