CREATE TABLE rep_db(
                    id SERIAL PRIMARY KEY,
                    title VARCHAR(50)UNIQUE ,
                    descriptions VARCHAR(255) ,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO rep_db(title, descriptions, created_at, updated_at) VALUES
                                            ('Вопрос','Сколько можно пинать пенисы йоу', '2025-01-30 15:02:30', '2025-01-30 15:03:30'),
                                            ('Ответ','Нога не кракен, жопа не отсос','2025-01-30 15:02:30','2025-01-30 15:03:30'),
                                            ('Загадка','Сосал??','2025-01-30 15:03:30','2025-01-30 15:04:30');
