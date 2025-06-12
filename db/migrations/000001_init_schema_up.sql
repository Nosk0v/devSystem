-- +goose Up
CREATE TABLE IF NOT EXISTS "goose_db_version" (
                                                  "id" SERIAL PRIMARY KEY,
                                                  "version_id" BIGINT NOT NULL,
                                                  "is_applied" BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "Organization" (
                                              "organization_id" SERIAL PRIMARY KEY,
                                              "name" VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "Department" (
                                            "department_id" SERIAL PRIMARY KEY,
                                            "name" VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS "MaterialType" (
                                              "type_id" SERIAL PRIMARY KEY,
                                              "type" VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "Competency" (
                                            "competency_id" SERIAL PRIMARY KEY,
                                            "name" VARCHAR(255) NOT NULL UNIQUE,
                                            "description" TEXT NOT NULL,
                                            "parent_id" INTEGER REFERENCES "Competency"("competency_id") ON DELETE SET NULL,
                                            "create_date" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "Material" (
                                          "material_id" SERIAL PRIMARY KEY,
                                          "title" VARCHAR(255) NOT NULL,
                                          "description" TEXT NOT NULL,
                                          "type" INTEGER REFERENCES "MaterialType"("type_id") ON DELETE SET NULL,
                                          "content" TEXT NOT NULL,
                                          "create_date" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "MaterialCompetency" (
                                                    "material_id" INTEGER NOT NULL REFERENCES "Material"("material_id") ON DELETE CASCADE,
                                                    "competency_id" INTEGER NOT NULL REFERENCES "Competency"("competency_id") ON DELETE CASCADE,
                                                    PRIMARY KEY ("material_id", "competency_id")
);



CREATE TABLE IF NOT EXISTS "Role" (
                                      "id" SERIAL PRIMARY KEY,
                                      "name" VARCHAR
);

CREATE TABLE IF NOT EXISTS "Account" (
                                         "email" VARCHAR PRIMARY KEY,
                                         "password" VARCHAR,
                                         "name" VARCHAR,
                                         "role" INTEGER NOT NULL REFERENCES "Role"("id"),
                                         "organization_id" INTEGER REFERENCES "Organization"("organization_id") ON DELETE CASCADE,
                                         "department_id" INTEGER REFERENCES "Department"("department_id") ON DELETE SET NULL
);



CREATE TABLE IF NOT EXISTS "Course" (
                                        "course_id" SERIAL PRIMARY KEY,
                                        "title" VARCHAR(255) NOT NULL UNIQUE,
                                        "description" TEXT NOT NULL,
                                        "create_date" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        "created_by" VARCHAR REFERENCES "Account"("email") ON DELETE SET NULL,
                                        "organization_id" INTEGER REFERENCES "Organization"("organization_id") ON DELETE CASCADE,
                                        "department_id" INTEGER REFERENCES "Department"("department_id") ON DELETE SET NULL
);


CREATE TABLE IF NOT EXISTS "CourseMaterial" (
                                                "course_id" INTEGER NOT NULL REFERENCES "Course"("course_id") ON DELETE CASCADE,
                                                "material_id" INTEGER NOT NULL REFERENCES "Material"("material_id") ON DELETE CASCADE,
                                                PRIMARY KEY ("course_id", "material_id")
);

CREATE TABLE IF NOT EXISTS "CourseCompetency" (
                                                  "course_id" INTEGER NOT NULL REFERENCES "Course"("course_id") ON DELETE CASCADE,
                                                  "competency_id" INTEGER NOT NULL REFERENCES "Competency"("competency_id") ON DELETE CASCADE,
                                                  PRIMARY KEY ("course_id", "competency_id")
);


CREATE TABLE IF NOT EXISTS "MaterialProgress" (
                                                  "user_email" VARCHAR NOT NULL REFERENCES "Account"("email") ON DELETE CASCADE,
                                                  "course_id" INTEGER NOT NULL REFERENCES "Course"("course_id") ON DELETE CASCADE,
                                                  "material_id" INTEGER NOT NULL REFERENCES "Material"("material_id") ON DELETE CASCADE,
                                                  "is_viewed" BOOLEAN DEFAULT TRUE,
                                                  "viewed_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                  PRIMARY KEY ("user_email", "course_id", "material_id")
);

CREATE TABLE IF NOT EXISTS "CourseProgress" (
                                                "progress_id" SERIAL PRIMARY KEY,
                                                "user_email" VARCHAR NOT NULL REFERENCES "Account"("email") ON DELETE CASCADE,
                                                "course_id" INTEGER NOT NULL REFERENCES "Course"("course_id") ON DELETE CASCADE,
                                                "is_completed" BOOLEAN DEFAULT FALSE,
                                                "completed_at" TIMESTAMP,
                                                UNIQUE("user_email", "course_id")
);

CREATE TABLE "RegistrationPrefix" (
                                      prefix TEXT PRIMARY KEY,                     -- напр. "MTI", "SBER"
                                      organization_id INT REFERENCES "Organization"(organization_id),
                                      created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE "InviteCode" (
                              code TEXT PRIMARY KEY,                        -- напр. "MTI-XYZ123"
                              prefix TEXT REFERENCES "RegistrationPrefix"(prefix),
                              role INT NOT NULL DEFAULT 1,                  -- 1=user, 2=admin
                              used BOOLEAN NOT NULL DEFAULT FALSE,
                              expires_at TIMESTAMP,                         -- NULL = бессрочный
                              created_at TIMESTAMP DEFAULT now(),
                              "department_id" INTEGER REFERENCES "Department"("department_id") ON DELETE SET NULL
);


INSERT INTO "Organization" ("name") VALUES
                                        ('ООО Альфа'),
                                        ('ЗАО Омега');



INSERT INTO "Department" ("name") VALUES
                                      ('Информационные технологии'),
                                      ('Менеджмент'),
                                      ('Дизайн'),
                                      ('Маркетинг'),
                                      ('Финансы'),
                                      ('Юриспруденция'),
                                      ('Продажи'),
                                      ('Образование и развитие'),
                                      ('Клиентский сервис'),
                                      ('Производство'),
                                      ('Аналитика'),
                                      ('Инженерия'),
                                      ('Логистика и снабжение'),
                                      ('Безопасность и комплаенс'),
                                      ('Человеческие ресурсы');

INSERT INTO "MaterialType" ("type")
VALUES
    ('Текст'),
    ('Статья'),
    ('Ссылка на видео');

INSERT INTO "RegistrationPrefix" ("prefix", "organization_id")
VALUES
    ('ALFA', 1),  -- ООО Альфа
    ('OMEGA', 2); -- ЗАО Омега

INSERT INTO "Material" ("title", "description", "type", "content") VALUES
('Обучение PostgreSQL', 'Базовые концепции СУБД', 3, 'https://rutube.ru/video/11111111111111111111111111111111/'),
('Docker: мастер-класс', 'Запуск контейнеров шаг за шагом', 3, 'https://rutube.ru/video/22222222222222222222222222222222/'),
('CI/CD в действии', 'Автоматизация развертывания', 3, 'https://rutube.ru/video/33333333333333333333333333333333/'),
('UX-интерфейсы', 'Создание удобных интерфейсов', 3, 'https://rutube.ru/video/44444444444444444444444444444444/'),
('JWT аутентификация', 'Безопасность API', 3, 'https://rutube.ru/video/55555555555555555555555555555555/'),
('Node.js производительность', 'Оптимизация кода', 3, 'https://rutube.ru/video/66666666666666666666666666666666/'),
('React для новичков', 'Первый проект на React', 3, 'https://www.youtube.com/watch?v=ABCDEFG1234'),
('Анализ производительности', 'Профилирование и тюнинг', 3, 'https://www.youtube.com/watch?v=HIJKLMN5678'),

('Управление доступом в облаках', 'RBAC и IAM', 1,
 'Управление доступом в облачных системах строится на модели RBAC (Role-Based Access Control), где роли назначаются субъектам (пользователям или сервисам) и определяют, какие ресурсы он может видеть и к каким операциям имеет доступ. Эта схема позволяет гибко управлять правами: создаются роли с набором прав — например, «viewer», «editor», «admin» — затем эти роли прикрепляются к аккаунтам. Комбинация ролей может применяться к ресурсам на разных уровнях: проект, каталог, организация. Например, роль admin на уровне проекта автоматически даёт права и на вложенные ресурсы — папки и виртуальные машины.\n\nВажно соблюдать принцип наименьших прав: субъект должен иметь только те привилегии, которые необходимы для выполнения его задач. Если сотрудник отвечает только за чтение данных — ему не нужна роль, позволяющая изменять. Каждая роль документируется, и регулярно проводится аудит: проверяется, кто и когда получал доступ, и зачем. Рекомендуется использовать сервисный аккаунт для автоматизированных задач, чтобы не привязывать системные операции к личным учеткам. Такой подход повышает безопасность и упростит отслеживание действий. Также следует учитывать наследование ролей: удаление роли на уровне родительского объекта автоматически влияет на дочерние ресурсы, что позволяет быстро менять архитектуру доступа.\n\nДля экстренных ситуаций рекомендуется временно выдавать повышенные права с автоматическим отзывом. Например, назначение роли editor с автоматическим снятием через 24 часа. Это сокращает риск ошибок и упрощает комплаенс. В Yandex Cloud роли предопределены: они охватывают управление VMs, сетями, базами данных и другими сервисами. При необходимости можно создавать кастомные роли, объединяя права из разных типов. Однако их количество лучше ограничивать, чтобы не запутаться в управлении. Каждый новый ресурс или сервис оценивается с точки зрения доступа: кто и зачем будет его использовать. Такой подход позволяет выстраивать эффективную и безопасную IAM-систему.\n\nСубъекты роли включают:\n- Личные аккаунты сотрудников\n- Сервисные аккаунты\n- Федеративные и групповые аккаунты\n\nКаждая группа и сервис может получать сразу несколько ролей. В Yandex Cloud связь между организацией и IAM сервисом позволяет централизованно управлять пользователями и группами. При необходимости можно подключить федерацию через внешнего провайдера (например, Active Directory). Это удобно для предприятий, где используются свои учетные системы. Также можно использовать автоматизированные сценарии через Terraform, создавая IAM-ресурсы через код, обеспечивая стабильность и повторяемость конфигураций.'),

('Мониторинг и логирование', 'Метрики, дашборды, алерты', 1,
 'Мониторинг и логирование — краеугольные камни надёжного окружения и быстрого обнаружения проблем. Мониторинг включает сбор метрик: CPU, память, диск, задержки по сетям и ошибочные запросы. Эти метрики отображаются на дашбордах (Grafana, Kibana), что позволяет отслеживать тренды и определить аномалии. Например, внезапный рост времени ответа HTTP может сигнализировать о проблеме в базе.\n\nЛогирование обеспечивает запись событий: запросов, ошибок, исключений. Логи направляются в централизованную систему (Elastic Stack, Loki, Splunk), где удобно искать по ключевым словам или паттернам. Важные поля - timestamp, уровень (INFO, ERROR), идентификатор сессии, IP-клиент, пользователь — позволяют быстро находить первопричину.\n\nАлертинг: системы (Prometheus Alertmanager, PagerDuty) отправляют оповещения в Slack или email, когда метрика превышает порог, или в логах появляются ошибки. Например, если процент ошибок HTTP > 5% за последние 5 минут — срабатывает алерт. Такие пороги настраиваются в зависимости от критичности сервиса.\n\nМониторинг должен охватывать инфраструктуру и бизнес-метрики. Например, сколько заказов в минуту обрабатывается магазином — это показатель здоровья бизнеса. Визуализация ожидаемых и фактических показателей помогает быстрее выявлять отклонения. Также важно обеспечить аудит изменений: кто и когда менял конфигурацию метрик или алертов. При этом тестовые окружения могут обходиться без сложных алертов, но для продакшена необходим полный стек и быстрые уведомления.\n\nВажно регулярно проводить «fire drills» — тестовые срабатывания. Это позволяет убедиться, что система работает и команда знает, как реагировать. После инцидента анализируются причины — Post Mortem — и вносятся улучшения — это цикл непрерывного улучшения.\n\nТаким образом,\n1) Сбор метрик и логов\n2) Визуализация на дашбордах\n3) Настройка алертов\n4) Тестирование реакций и улучшение процесса\nявляются необходимыми шагами к стабильной и контролируемой инфраструктуре.'),

('Контейнеризация и оркестрация', 'Docker, Kubernetes', 1,
 'Контейнеризация — упаковка приложения вместе со всеми зависимостями в контейнер с помощью Docker. Это решает проблему «работало на моей машине». Контейнер содержит runtime, библиотеки и настройки — запускается одинаково на любом узле. Dockerfile описывает шаги сборки: базовый образ, копирование кода, установка зависимостей, запуск приложения.\n\nKubernetes — система оркестрации контейнеров: запускает контейнеры (Pods) на разных узлах кластера, автоматически восстанавливает упавшие, масштабирует по нагрузке, балансирует трафик.\n\nОсновные объекты:\n- Deployment: описывает, какой образ запустить и сколько реплик держать\n- Service: обеспечивает доступ к приложению внутри кластера (ClusterIP, NodePort, LoadBalancer)\n- ConfigMap/Secret: хранит конфигурацию и секреты\n- Ingress: управляет внешним доступом через HTTP\n\nKubernetes в продакшне часто разворачивают через Helm — шаблонный менеджер пакетов, позволяющий описать зависимости и собрать готовый чарт. CI/CD автоматизирует сборку образов и инициализацию Helm-чартов.\n\nПреимущества контейнеров:\n- Изоляция приложений\n- Легкость масштабирования\n- Упрощённый деплой на разные среды\n\nНедостатки:\n- Требуют настройки сети внутри кластера\n- Могут усложнить отладку из-за распределённости\n\nСоветы:\n- Минимизируйте образы, удаляйте dev-зависимости перед пушем\n- Используйте readiness и liveness пробы для Kubernetes\n- Применяйте автоматическое масштабирование (HPA) на основе CPU или queue length\n- Разграничивайте пространства имен (namespaces) для изоляции окружений\n\nКонтейнеризация — это не панацея, но если приложение требует высокой доступности и гибкости, она даёт значительные преимущества. Kubernetes идеален для микросервисов, однако для простых случаев можно обойтись Docker Compose.'),

('Инфраструктура как код', 'Terraform и описательные языки', 1,
 'Практика Infrastructure as Code (IaC) преобразует инфраструктуру (виртуальные машины, сети, базы данных) в код, описываемый декларативно. Самый популярный инструмент — Terraform: HCL позволяет описать ресурсы облаков, их связи, конфигурации. Например:\n\nresource "yandex_compute_instance" "vm" {...}\nresource "yandex_vpc_network" "net" {...}\nresource "yandex_vpc_subnet" "subnet" {...}\n\nTerraform хранит текущее состояние ресурса, сравнивает его с желаемым и предлагает изменения: apply, plan, destroy. Это гарантирует, что инфраструктура воспроизводима и повторяема.\n\nIaC обеспечивает:\n- Версионирование: инфраструктура хранится в git\n- Peer review: изменения проверяются через pull-request\n- Predictability: terraform plan точно показывает, что изменится\n- Авторазвертывание: CI/CD применяет инфраструктуру автоматически, снижая ручные ошибки\n\nТакже Terraform работает с модулями: повторно используемыми кусками конфигурации, например, машина с мониторинг-агентом или VPC с ограничением доступа.\n\nИнтеграция с CI/CD (GitLab, GitHub Actions) позволяет запускать `terraform plan` на каждом PR и деплоить только после проверки. Важен подход к secret management: пароли, токены хранятся в Vault/Secrets Manager, не в коде или state-файле.\n\nIaC становится базовым стандартом для надёжной, безопасной и масштабируемой инфраструктуры. Даже небольшие команды выигрывают от автоматизации, а большие — отповторяемости и auditability.\n\nРекомендуемые практики:\n- Хранить state файл в удалённом backend (S3, ЯндексОблако Object Storage)\n- Разделять окружения (prod/dev) в разные воркспейсы\n- Использовать линтеры (terraform fmt, validate, tflint)\n- Обновлять провайдеры с учётом обратной совместимости\n\nТаким образом, IaC — не только про удобство, но и про безопасность, прозрачность и контроль.'),


('API дизайн: лучшие практики', 'REST vs GraphQL', 2, 'https://restfulapi.net/'),
('TypeScript миграция', 'Почему стоит перейти', 2, 'https://www.typescriptlang.org/docs/handbook/migrating-from-javascript.html'),
('Безопасность OWASP', 'Top-10 уязвимостей', 2, 'https://owasp.org/www-project-top-ten/'),
('Git cheat sheet', 'Основные команды и паттерны', 2, 'https://git-scm.com/docs'),
('MDN Frontend guide', 'Учебник по HTML/CSS/JS', 2, 'https://developer.mozilla.org/ru/docs/Learn'),
('Проектирование БД', 'Нормализация, индексы', 2, 'https://www.digitalocean.com/community/tutorials/sql-normalization'),
('PostgreSQL оптимизация', 'EXPLAIN и настройки', 2, 'https://www.postgresql.org/docs/current/using-explain.html'),
('DevOps introduction', 'Continuous Integration/Delivery', 2, 'https://www.atlassian.com/continuous-delivery')
;

INSERT INTO "Competency" ("name", "description", "parent_id") VALUES
                                                                  ('Software Development', 'Основы проектирования и реализации программных решений', NULL),
                                                                  ('Backend Engineering', 'Разработка серверной логики и API-интерфейсов', 1),
                                                                  ('Frontend Engineering', 'Создание пользовательских интерфейсов веб-приложений', 1),
                                                                  ('Database Design', 'Проектирование и нормализация структур баз данных', 1),
                                                                  ('PostgreSQL', 'Работа с реляционной СУБД PostgreSQL', 4),
                                                                  ('React.js', 'Разработка SPA с использованием React', 3),
                                                                  ('REST API Design', 'Создание RESTful API с учетом лучших практик', 2),
                                                                  ('Security Basics', 'Основы безопасности приложений и защита данных', 2),
                                                                  ('Version Control', 'Использование Git и систем контроля версий', 1),
                                                                  ('DevOps Fundamentals', 'Настройка CI/CD и автоматизация процессов', 2),
                                                                  ('Unit Testing', 'Написание и запуск модульных тестов', 2),
                                                                  ('UX/UI Principles', 'Базовые принципы удобства и дизайна интерфейсов', 3),
                                                                  ('Responsive Design', 'Адаптивная вёрстка под разные устройства', 3),
                                                                  ('Typescript', 'Использование строгой типизации в разработке', 3),
                                                                  ('Node.js', 'Серверная разработка на Node.js', 2),
                                                                  ('Docker Basics', 'Контейнеризация приложений с помощью Docker', 10),
                                                                  ('Agile Practices', 'Гибкие методологии разработки, Scrum, Kanban', NULL),
                                                                  ('API Integration', 'Интеграция сторонних API в приложения', 7),
                                                                  ('Performance Optimization', 'Повышение производительности веб-приложений', 2),
                                                                  ('Authentication', 'Реализация регистрации и входа с токенами', 8);

INSERT INTO "MaterialCompetency" ("material_id", "competency_id") VALUES
                                                                      (1, 1), (1, 9), (1, 17),
                                                                      (2, 2), (2, 7), (2, 8),
                                                                      (3, 3), (3, 6), (3, 13),
                                                                      (4, 4), (4, 1), (4, 17),
                                                                      (5, 5), (5, 4), (5, 19),
                                                                      (6, 6), (6, 14), (6, 11),
                                                                      (7, 7), (7, 2), (7, 18),
                                                                      (8, 8), (8, 20), (8, 7),
                                                                      (9, 9), (9, 1), (9, 17),
                                                                      (10, 10), (10, 16), (10, 5),
                                                                      (11, 11), (11, 2), (11, 20),
                                                                      (12, 12), (12, 3), (12, 13),
                                                                      (13, 13), (13, 6), (13, 3),
                                                                      (14, 14), (14, 3), (14, 11),
                                                                      (15, 15), (15, 2), (15, 7),
                                                                      (16, 16), (16, 10), (16, 5),
                                                                      (17, 17), (17, 1), (17, 18),
                                                                      (18, 18), (18, 7), (18, 2),
                                                                      (19, 19), (19, 3), (19, 14),
                                                                      (20, 20), (20, 8), (20, 18);
-- Роли
INSERT INTO "Role" ("id", "name")
VALUES
    (1, 'Пользователь'),
    (0, 'Администратор'),
    (2, 'Супер админ');

-- Аккаунты
INSERT INTO "Account" ("email", "password", "name", "role", "organization_id", "department_id") VALUES
                                                                                                    ('1234@mail.ru', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Иванов Иван Иванович', 0, 1, NULL),
                                                                                                    ('5678@mail.ru', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Петров Петр Петрович', 1, 2, 1),
                                                                                                    ('root@system.dev', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Суперадмин', 2, NULL, NULL);

INSERT INTO "Course" ("title", "description", "created_by", "organization_id", "department_id") VALUES
                                                                                                    ('Курс по Backend', 'Освоение серверной разработки на Go и PostgreSQL', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('Frontend для начинающих', 'Изучение React и вёрстки с нуля', '5678@mail.ru', 2, 1), -- Информационные технологии
                                                                                                    ('Основы Golang', 'Курс по основам языка программирования Go', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('Микросервисы на Go', 'Создание микросервисной архитектуры', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('GRPC и Protocol Buffers', 'Работа с gRPC и Protobuf в бэкенде', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('JWT и безопасность API', 'Реализация авторизации на JWT', '1234@mail.ru', 1, 14), -- Безопасность и комплаенс
                                                                                                    ('Docker для разработчиков', 'Контейнеризация сервисов и Dockerfile', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('CI/CD с GitHub Actions', 'Автоматизация сборок и деплоя', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('SQL и оптимизация запросов', 'Повышение производительности SQL-запросов', '1234@mail.ru', 1, 11), -- Аналитика
                                                                                                    ('PostgreSQL в бою', 'Использование расширений и индексов', '1234@mail.ru', 1, 11), -- Аналитика
                                                                                                    ('Работа с REST API', 'Создание и тестирование REST-интерфейсов', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('Тестирование бэкенда', 'Unit, интеграционные и e2e тесты', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('Проектирование БД', 'Связи, нормализация, схема', '1234@mail.ru', 1, 11), -- Аналитика
                                                                                                    ('Введение в DevOps', 'Как связаны разработка и эксплуатация', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('Мониторинг и алертинг', 'Grafana, Prometheus, alertmanager', '1234@mail.ru', 1, 1), -- Информационные технологии
                                                                                                    ('Методологии Scrum и Agile', 'Управление командной разработкой', '1234@mail.ru', 1, 2), -- Менеджмент
                                                                                                    ('Секреты производительности Go', 'Профилирование и ускорение кода', '1234@mail.ru', 1, 1); -- Информационные технологии

-- Связи курс-материал
INSERT INTO "CourseMaterial" ("course_id", "material_id") VALUES
                                                              (1, 2),
                                                              (1, 5),
                                                              (2, 3),
                                                              (2, 6),
                                                              (3, 1), (3, 2),
                                                              (4, 5), (4, 7),
                                                              (5, 16), (5, 10),
                                                              (6, 10), (6, 11),
                                                              (7, 4), (7, 19),
                                                              (8, 5), (8, 4),
                                                              (9, 7), (9, 18),
                                                              (10, 11), (10, 14),
                                                              (11, 4), (11, 1),
                                                              (12, 10), (12, 16),
                                                              (13, 20), (13, 5),
                                                              (14, 17), (14, 1),
                                                              (15, 2), (15, 19),
                                                              (16, 1), (16, 15),
                                                              (17, 2), (17, 14);

-- Связи курс-компетенция
INSERT INTO "CourseCompetency" ("course_id", "competency_id") VALUES
                                                                  (1, 2),
                                                                  (1, 5),
                                                                  (2, 3),
                                                                  (2, 6),
                                                                  (3, 2), (3, 14),
                                                                  (4, 20), (4, 8),
                                                                  (5, 16), (5, 10),
                                                                  (6, 10), (6, 2),
                                                                  (7, 4), (7, 19),
                                                                  (8, 5), (8, 4),
                                                                  (9, 7), (9, 18),
                                                                  (10, 11), (10, 20),
                                                                  (11, 4), (11, 1),
                                                                  (12, 10), (12, 5),
                                                                  (13, 20), (13, 8),
                                                                  (14, 17), (14, 1),
                                                                  (15, 2), (15, 7),
                                                                  (16, 1), (16, 19),
                                                                  (17, 14), (17, 2);

INSERT INTO "Competency" ("name", "description", "parent_id") VALUES
  ('Коммуникативная грамотность', 'Совокупность знаний, умений и навыков человека, позволяющих ему эффективно общаться в стандартных коммуникативных ситуациях в письменной и устной форме', NULL),
  ('Партнерство/сотрудничество', 'Корректное взаимодействие с другими людьми, выстраивание отношений сотрудничества, выявление и учёт потребностей и интересов других, предложение взаимовыгодных решений и работу над совместным развитием идей или проектов для достижения общей цели', NULL),
  ('Эмоциональный интеллект', 'Способность человека распознавать свои и чужие эмоции, понимать намерения собеседника, его мотивацию и желания', NULL),
  ('Лидерство', 'Способность определённого человека эффективно влиять на действия, поведение и мотивацию других людей с целью достижения определённого результата', NULL),
  ('Стрессоустойчивость', 'Совокупность качеств, позволяющих организму спокойно переносить действие стрессоров без вредных всплесков эмоций, влияющих на деятельность и на окружающих, а также способных вызывать психические расстройства', NULL);


-- +goose Down
DROP TABLE IF EXISTS "MaterialCompetency";
DROP TABLE IF EXISTS "CourseMaterial";
DROP TABLE IF EXISTS "CourseCompetency";
DROP TABLE IF EXISTS "MaterialProgress";
DROP TABLE IF EXISTS "CourseProgress";
DROP TABLE IF EXISTS "Course";
DROP TABLE IF EXISTS "Material";
DROP TABLE IF EXISTS "Competency";
DROP TABLE IF EXISTS "MaterialType";
DROP TABLE IF EXISTS "InviteCode";
DROP TABLE IF EXISTS "RegistrationPrefix";
DROP TABLE IF EXISTS "Account";
DROP TABLE IF EXISTS "Role";
DROP TABLE IF EXISTS "Department";
DROP TABLE IF EXISTS "Organization";
DROP TABLE IF EXISTS "goose_db_version";