INSERT INTO items(name, size, quantity)
VALUES ('книга', 2, 100),
       ('стол', 7, 30),
       ('цветок', 4, 17),
       ('платье', 1, 4);

INSERT INTO stores(name, availability)
VALUES ('первый склад', FALSE),
       ('второй склад', TRUE),
       ('третий склад', TRUE);

INSERT INTO available(store_id, item_id, item_quantity, reserved_item)
VALUES (1, 3, 7, 0),
       (1, 1, 33, 8),
       (2, 1, 51, 23),
       (3, 1, 16, 16),
       (1, 2, 15, 0),
       (2, 2, 10, 5),
       (3, 2, 5, 1),
       (3, 3, 10, 2),
       (3, 4, 4, 0);