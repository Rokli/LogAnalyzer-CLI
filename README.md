# Log Analyzer CLI

Консольная утилита, которая анализирует текстовые логи и показывает статистику.

## Установка

### Установка из исходного кода
```bash
git clone https://github.com/Rokli/LogAnalyzer-CLI.git
cd LogAnalyzer-CLI
make install
```
## Пример поддерживаемых логов

```
2024-01-01 10:00:00 INFO User logged in
2024-01-01 10:00:05 ERROR Database connection failed
```
## Пример команд
```
# Статистика по логам
log-analyzer -file app.log -stats

# Фильтрация по уровню
log-analyzer -file app.log -level ERROR

# Поиск подстроки
log-analyzer -file app.log -search "connection"

# Экспорт в JSON
log-analyzer -file app.log -output json

# Ограничение вывода
log-analyzer -file app.log -limit 100
```
