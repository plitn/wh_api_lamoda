Для того чтобы запустить сервис необходимо в терминале, находясь в директории проекта, запустить команду ```make up``` в директории проекта

Чтобы оставновить работу сервсиа необходимо запустить команду ```make clean``` в терминале, находясь в директории проекта.

Как работает логика резервации / отмены резервации товара: Товары с одинаковым product_id могут находится на разных складах, поэтому при резервации товара нужно указывать склад (wh_id), на котором находится товар 
и массив связок product_id + qty , где qty -- количество, которое нужно зарезервировать / снять с резерва

Инфу об айдишниках и тд можно посмотреть в init.sql или же залеть в БД после поднятия инфраструктуры
