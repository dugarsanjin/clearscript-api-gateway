# project clearscript-api-gateway

### Branching strategy: 
Trunk-based development
(предполагается внедрение управления фича флагами в будущем)

### Configuration:
задать переменную окружения `CONFIG_PATH` с путем к файлу конфигурации

### Endpoints:

- GET /api/v1/users/{id} - возвращает информацию о пользователе
- GET /api/v1/lessons/{id} - возвращает контент урока с письменными символами
